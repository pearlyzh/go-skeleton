install:
	go mod download
	go mod tidy
install-plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
gen-proto:
	rm -rf generated
	protoc --go_out=. --go-grpc_out=. --go_opt=module=github.com/pearlyzh/go-skeletons --go-grpc_opt=module=github.com/pearlyzh/go-skeletons proto/*.proto