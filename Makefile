API_PATH 		= api/student
PROTO_API_DIR 	= api/student
PROTO_OUT_DIR 	= pkg/studentServiceApi
PROTO_API_OUT_DIR = ${PROTO_OUT_DIR}
MOCKS_DIR = internal/student

.PHONY: gen-proto-ss
gen-proto-ss:
	mkdir -p ${PROTO_OUT_DIR}
	protoc \
		-I ${API_PATH} \
		-I third_party/googleapis \
		--include_imports \
        --grpc-gateway_out=logtostderr=true:$(PROTO_OUT_DIR) \
		--descriptor_set_out=$(PROTO_API_OUT_DIR)/api.pb \
		./${PROTO_API_DIR}/*.proto

.PHONY: go/lint
go/lint:
	golangci-lint run  --config=.golangci.yml --timeout=180s ./...

.PHONY: go-mock-install
go-mock-install:
	GO111MODULE=on go get github.com/golang/mock/mockgen@v1.5.0

.PHONY: mocks-generate
mocks-generate:
	@echo "generate-mocks"
	mockgen -package=student -source=$(MOCKS_DIR)/client.go -destination=$(MOCKS_DIR)/mock_gen.go
	@echo "successfully"
