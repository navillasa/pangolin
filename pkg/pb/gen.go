package pb

//go:generate protoc -I=. --go_out=plugins=grpc:. ./server.proto