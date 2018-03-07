all: gen

gen:
	go install -v ./vendor/github.com/golang/protobuf/protoc-gen-go
	go install -v ./vendor/github.com/Percona-Lab/wsrpc/cmd/protoc-gen-wsrpc
	rm -f gateway/*.pb.go gateway/*.wsrpc.go
	protoc gateway/*.proto --go_out=.
	protoc gateway/*.proto --wsrpc_out=.
	go install -v ./...
