all: the-client the-server

the-client: client/main.go protobuf-src/challenge.pb.go
	go build -o the-client ./client

the-server: server/main.go protobuf-src/challenge.pb.go
	go build -o the-server ./server

protobuf-src/challenge.pb.go: protobuf-src/challenge.proto
	protoc -I protobuf-src/ --go_out=plugins=grpc:protobuf-src protobuf-src/challenge.proto
