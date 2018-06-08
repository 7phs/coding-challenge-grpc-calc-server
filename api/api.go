//go:generate protoc -I . -I${GOPATH}/src --go_out=plugins=grpc:. calc.proto

package api
