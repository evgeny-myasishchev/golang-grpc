// +build tools

// See https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
// for some notes on this file

package tools

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
