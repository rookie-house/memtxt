package tools

import (
	"context"
	"fmt"

	redishandler "github.com/harshduche/memtxt/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type PubInput struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

type PubOutput struct {
	Status string `json:"status"` 
}


func Pub(ctx context.Context, res *mcp.CallToolRequest, in PubInput) (*mcp.CallToolResult, PubOutput, error)  {
	
	err := redishandler.Client.Publish(redishandler.Ctx, in.Channel, in.Message).Err()
	
	if err != nil{
		return nil, PubOutput{}, fmt.Errorf("publish failed: %v",err)
	}

	return  nil, PubOutput{Status: "Message Published"}, nil	
}
