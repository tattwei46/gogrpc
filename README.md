# **Setting GOPATH**
1. export PATH=$PATH=/Users/davidcheah/Desktop/Tools/protoc-3.10.1-osx-x86_64/bin
2. Setup $GOPATH to working directory using . ./tools.setup.sh

# **Dependency Management**
1. Create /src/domain and cd into it
2. dep init 
3. dep ensure --add google.golang.org/grpc
4. dep ensure --add github.com/golang/protobuf/protoc-gen-go

# **Without Dependency Management**
1. go get -u google.golang.org/grpc
2. go get -u github.com/golang/protobuf/protoc-gen-go

# **Compile .proto and generate .pb.go
1. protoc -I api/     -I${GOPATH}/src     --go_out=plugins=grpc:api     api/api.proto