update all protoc-gen-go using go get -u (destination port)google.org {protoc,protoc-gen-go}
move to the respective folder of greet/greetpb and execute protoc command
protoc greet.proto --go_out=plugins=grpc:. *.proto
will generate the pb.go file




commands :

protoc greet.proto --go_out=plugins=grpc:. *.proto

