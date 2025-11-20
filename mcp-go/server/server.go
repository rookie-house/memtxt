package main

import (
	"context"
	"log"

	"github.com/harshduche/memtxt/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {

	redis
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "memtxt-server",
		Version: "v1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "greet",
		Description: "Greets a person by name",
	}, tools.Greet)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}

