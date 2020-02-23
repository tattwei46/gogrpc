# **Install dependencies**
1. Install protoc and place in /bin
2. go get -u github.com/golang/protobuf/protoc-gen-go

# **Compile .proto and generate .pb.go**
1. go generate server/server.go

# **Run project**
1. go run main.go