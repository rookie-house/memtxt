package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	client := mcp.NewClient(&mcp.Implementation{
		Name:    "memtxt-client",
		Version: "v1.0.0",
	}, nil)

	transport := &mcp.CommandTransport{
		Command: exec.Command("../server/memtxt-server"),
	}

	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatal("connection error:", err)
	}
	defer session.Close()

	params := &mcp.CallToolParams{
		Name: "greet",
		Arguments: map[string]any{
			"name": "Harsh",
		},
	}

	res, err := session.CallTool(ctx, params)
	if err != nil {
		log.Fatal("tool call failed:", err)
	}

	for _, c := range res.Content {
		log.Println(c.(*mcp.TextContent).Text)
	}
}

