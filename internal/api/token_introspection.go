package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"strings"
	"xgmdr.com/pad/internal/logger"
)

func introspectToken(token *jwt.Token) {
	var parts []string = strings.Split(token.Raw, ".")

	for i, v := range parts {
		if i == 2 {
			continue // skip signature part
		}

		v = addPadding(v)
		decoded, error := base64.URLEncoding.DecodeString(v)
		if error != nil {
			logger.Get().Warn("Unable to decode token",
				zap.Error(error),
			)
			return
		}

		var objmap interface{}
		error = json.Unmarshal(decoded, &objmap)

		if error != nil {
			logger.Get().Warn("Unable to unmarshal token",
				zap.Error(error),
			)
			return
		}

		indentedJson, error := json.MarshalIndent(objmap, "", "\t")

		fmt.Println(string(indentedJson))
	}
}

func addPadding(v string) string {
	if remainder := len(v) % 4; remainder > 0 {
		for j := 0; j < 4-remainder; j++ {
			v = v + "="
		}
	}
	return v
}
