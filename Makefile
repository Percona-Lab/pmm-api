all: gen

gen:
	go install -v ./vendor/github.com/golang/protobuf/protoc-gen-go

	rm -f agent/*.go
	protoc agent/*.proto --go_out=plugins=grpc:.

	rm -f gateway/*.go
	protoc gateway/*.proto --go_out=plugins=grpc:.

	go install -v ./...
