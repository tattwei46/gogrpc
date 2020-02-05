# **Setting GOPATH**
1. export PATH=$PATH=/Users/davidcheah/Desktop/Tools/protoc-3.10.1-osx-x86_64/bin
2. Setup $GOPATH to working directory using . ./tools.setup.sh

# **Dependency Management**
1. Create /src/<domain> and cd $GOPATH/src/<domain>
2. dep init 
3. dep ensure --add google.golang.org/grpc
4. dep ensure --add github.com/golang/protobuf/protoc-gen-go

# **Without Dependency Management**
1. go get -u google.golang.org/grpc
2. go get -u github.com/golang/protobuf/protoc-gen-go

# **Compile .proto and generate .pb.go**
1. go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
2. protoc -I proto/     -I${GOPATH}/src    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis    --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:proto     proto/user.proto
