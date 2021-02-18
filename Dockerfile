FROM golang:1.15.7@sha256:4c9eeab8adf54d893450f6199f52cf7bb39264750ee2a11018dd41acfe6aeaba

RUN go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \ 
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway 


RUN apt-get update 

RUN apt install -y protobuf-compiler
RUN protoc --version

VOLUME ["/grpc"]
WORKDIR /grpc

ENTRYPOINT [ "go", "run", "server.go" ]