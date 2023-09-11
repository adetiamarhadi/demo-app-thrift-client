package main

import (
	"context"
	"fmt"

	"os"

	"github.com/adetiamarhadi/demo-app-thrift-client/thrift/myservice"
	"github.com/apache/thrift/lib/go/thrift"
)

func main() {
	ctx := context.Background()
	transport, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		fmt.Println("Error creating socket:", err)
		os.Exit(1)
	}
	if err := transport.Open(); err != nil {
		fmt.Println("Error opening socket:", err)
		os.Exit(1)
	}
	defer transport.Close()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := myservice.NewMyServiceClientFactory(transport, protocolFactory)
	param := "Hello, Thrift!"
	response, err := client.SayHello(ctx, param)
	if err != nil {
		fmt.Println("Error calling SayHello:", err)
		os.Exit(1)
	}
	fmt.Printf("Response: %+v\n", response)
}
