


## Running swagger-ui

```
docker run -e SWAGGER_JSON=/spec/openapi3.yaml -v $(pwd):/spec -p 8091:8080 swaggerapi/swagger-ui
```

## Handler code generation

APIを増やす場合は、 `openapi3.yaml` を定義した後

```
go generate ./...
```

を実行します. これにより `server/server_gen.go` が生成されるので、その後 `StrictServerInterface` を満たすように `server/handler.go` を実装します
