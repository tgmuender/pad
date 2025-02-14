package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	oauth2Config *oauth2.Config
	state        = "randomStateString"
	tokenFile    = filepath.Join(os.Getenv("HOME"), ".padctl_token.json")
)

func init() {
	oauth2Config = &oauth2.Config{
		ClientID:     "padcli",
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://127.0.0.1:5556/auth",
			TokenURL: "http://127.0.0.1:5556/token",
		},
	}
}

// loginCommand returns a cobra command for login
func loginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate to the server and store your authentication credentials for future requests.",
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/login", handleLogin)
			http.HandleFunc("/callback", handleCallback)
			log.Println("Starting server at :8080")

			go func() {
				url := oauth2Config.AuthCodeURL(state)
				err := exec.Command("xdg-open", url).Start()
				if err != nil {
					log.Fatalf("Failed to open URL: %v", err)
				}
			}()
			log.Fatal(http.ListenAndServe(":8080", nil))

		},
	}
	return cmd
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("state") != state {
		http.Error(w, "State mismatch", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokenData, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Failed to marshal token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := ioutil.WriteFile(tokenFile, tokenData, 0600); err != nil {
		http.Error(w, "Failed to save token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Login successful, token saved.")
	fmt.Println("Login successful, token saved.")
	os.Exit(0)
}
