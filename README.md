# store_cqrsclient

```zsh
go get -u github.com/akira-saneyoshi/store_pb@v1.0.0

go get -u google.golang.org/grpc
go get -u github.com/onsi/ginkgo/v2
go get -u github.com/onsi/gomega
go get -u go.uber.org/fx
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/go-openapi/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

go install github.com/swaggo/swag/cmd/swag@latest
```

```zsh
swag init --parseDependency -g main.go
```
