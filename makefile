proto-gen:
	protoc --proto_path=shared/proto --go_out=shared/proto --go_grpc_out=shared/proto shared/proto/profile.proto

list:
	node project-content.js