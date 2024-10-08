PROTO_PATH=proto
GO_OUT_PATH=gen/go
JAVA_PROTO_PATH=proto-lite
JAVA_OUT_PATH=gen/java
KOTLIN_OUT_PATH=gen/kotlin
SWAGGER_DOC_DIR=docs

install:
	go mod tidy
clean:
	go clean -cache
	go clean -i
	go clean -testcache
	go clean -modcache
generate:
	rm -rf ${GO_OUT_PATH}/*.go
	rm -f ${SWAGGER_DOC_DIR}/*.swagger.json
	protoc --proto_path=${PROTO_PATH} \
	--go_out=${GO_OUT_PATH} \
	--go_opt=paths=source_relative \
	--go-grpc_out=${GO_OUT_PATH} \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out ${GO_OUT_PATH} --grpc-gateway_opt paths=source_relative \
    --openapiv2_out=${SWAGGER_DOC_DIR} \
    --openapiv2_opt=allow_merge=true,merge_file_name=apidocs \
	${PROTO_PATH}/*.proto

generate-java:
	rm -rf ${JAVA_OUT_PATH}/*ci
	protoc --proto_path=${JAVA_PROTO_PATH} \
	--java_out=${JAVA_OUT_PATH} \
	--grpc-java_out=${JAVA_OUT_PATH} \
	${JAVA_PROTO_PATH}/*.proto

generate-kotlin:
	rm -rf ${KOTLIN_OUT_PATH}/*ci
	protoc --proto_path=${PROTO_PATH} \
	--kotlin_out=${KOTLIN_OUT_PATH} \
	--grpc-kotlin_out=${KOTLIN_OUT_PATH} \
;	--plugin=protoc-gen-kotlin=tmp/protoc-gen-kotlin.bat \
	${PROTO_PATH}/*.proto

.PHONY: install clean generate generate-java generate-kotlin
