# Lineo ID proto library

### _Go plugins_ for the protocol compiler

Install the protocol compiler plugins for Go using the following commands:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Install GRPC Gateway(optional)

```bash
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```

### Generate proto file to Go

```bash
make generate
```

with buf

```bash
buf generate
# or
go run github.com/bufbuild/buf/cmd/buf@v1.28.1 generate
```

with task

```bash
task generate
```
