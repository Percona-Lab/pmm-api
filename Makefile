all: gen

gen:
	go install -v ./vendor/github.com/golang/protobuf/protoc-gen-go
	go install -v ./vendor/github.com/Percona-Lab/wsrpc/cmd/protoc-gen-wsrpc

	rm -f agent/*.go
	protoc agent/*.proto --go_out=.
	protoc agent/*.proto --wsrpc_out=.

	rm -f gateway/*.go
	protoc gateway/*.proto --go_out=.
	protoc gateway/*.proto --wsrpc_out=.

	go install -v ./...
