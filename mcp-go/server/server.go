package main

import (
	"context"
	"log"

	"github.com/harshduche/memtxt/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "memtxt-server",
		Version: "v1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "greet",
		Description: "Greets a person by name",
	}, tools.Greet)

	mcp.AddTool(server, &mcp.Tool{Name: "pub", Description: "Publish event"}, tools.Pub)
	mcp.AddTool(server, &mcp.Tool{Name: "sub", Description: "Subscribe room"}, tools.Sub)
	mcp.AddTool(server, &mcp.Tool{Name: "stream_write", Description: "write to streams"}, tools.StreamWrite)
	mcp.AddTool(server, &mcp.Tool{Name: "stream_read", Description: "read streams"}, tools.StreamRead)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
