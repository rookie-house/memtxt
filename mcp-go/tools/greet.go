package tools

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GreetInput struct {
	Name string `json:"name" jsonschema:"name of the person to greet"`
}

type GreetOutput struct {
	Greeting string `json:"greeting" jsonschema:"the greeting"`
}

func Greet(ctx context.Context, req *mcp.CallToolRequest, in GreetInput) (
	*mcp.CallToolResult,
	GreetOutput,
	error,
) {
	return nil, GreetOutput{
		Greeting: "Hello " + in.Name,
	}, nil
}

