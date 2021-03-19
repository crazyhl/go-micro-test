// +build tools

package tools

import _ "google.golang.org/protobuf/cmd/protoc-gen-go" // need go install
import _ "github.com/asim/go-micro/cmd/protoc-gen-micro/v3" // need go install
import _ "github.com/asim/go-micro/v3"
import _ "github.com/asim/go-micro/plugins/registry/etcd/v3"
