
### Project Layout

See 
- https://github.com/golang-standards/project-layout 
- https://github.com/golang-standards/project-layout/tree/master/internal

Code examples:

- https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
- https://protobuf.dev/getting-started/gotutorial/
- https://protobuf.dev/programming-guides/style/
- https://grpc.io/docs/languages/go/basics/#starting-the-server

### OAuth2 Client

`docker exec -it auth-hydra-1 hydra create oauth2-client --skip-consent --name petadvisor --secret testli --endpoint http://127.0.0.1:4445 --redirect-uri='http://localhost:4180/oauth2/callback' --scope email,profile,openid`

```
CLIENT ID       0fd8932c-71d7-4c48-a7d5-04f6863eb5a9
CLIENT SECRET   testli
GRANT TYPES     authorization_code
RESPONSE TYPES  code
SCOPE           offline_access offline openid
AUDIENCE
REDIRECT URIS   http://localhost:8888/callback
```

### OAuth2 Client Google

- https://support.google.com/cloud/answer/6158849?hl=en
- https://console.cloud.google.com/welcome?inv=1&invt=Abl1Ww&project=petadvisor-432206