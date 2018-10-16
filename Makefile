all: gen

init:
	# https://github.com/uber/prototool#installation
	curl -L https://github.com/uber/prototool/releases/download/v1.3.0/prototool-$(shell uname -s)-$(shell uname -m) -o ./prototool
	chmod +x ./prototool

gen:
	go install -v ./vendor/github.com/golang/protobuf/protoc-gen-go \
					./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
					./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
					./vendor/github.com/go-swagger/go-swagger/cmd/swagger
	find . -name '*.pb.go' -not -path './vendor/*' -delete
	find . -name '*.pb.ge.go' -not -path './vendor/*' -delete
	./prototool all
