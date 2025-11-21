package tools

import (
	"context"
	"fmt"

	redishandler "github.com/harshduche/memtxt/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)


type GitInfo struct{
	Logs string;
	History string;
}

type EventInput struct {
	Summary string `json:"summary"`
	GitInfo GitInfo `json:"gitinfo"`
	RoomId  string  `json:"roomid"`
}


type EventOutput struct {
	Status string `json:"status"`
}

func PublishEvent(ctx context.Context, res *mcp.CallToolRequest, in EventInput) (*mcp.CallToolResult, EventOutput, error)  {
	err :=  redishandler.Client.Publish(redishandler.Ctx, in.RoomId, in)
	if err != nil {
		return nil, EventOutput{}, fmt.Errorf("error in publish: %v", err)
	}

	return nil, EventOutput{Status: "success"}, nil 
}
