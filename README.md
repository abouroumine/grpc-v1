# grpc-v1
Basic Application that shows the implementation of gRPC Communication Protocol in Go.
we use in this example a simple uniry Communication style between the Server <=> Client.

# Steps
1- We start by implementing the Service we want in the (.proto) file.
2- We install the required modules:

		$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
		$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
		
3- Update your PATH so that the protoc compiler can find the plugins:

		$ export PATH="$PATH:$(go env GOPATH)/bin"

for more Info you can check ( https://grpc.io/docs/languages/go/quickstart/ )

4- We use the commands to generate the (*.pb.go files):

	    	$ protoc --go_out=. --go_opt=paths=source_relative filename.proto
	    	$ protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative filename.proto
    
5- We implement the Methods in the Client / Server sides to utilise the Service defined in the filename.proto
