package tools

import (
	"context"
	"fmt"

	redishandler "github.com/harshduche/memtxt/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SubInput struct {
	Channel string `json:"channel"`
}

type SubOutput struct {
	Messages []string `json:"messages"`
}

func Sub(ctx context.Context, res *mcp.CallToolRequest, in SubInput) (*mcp.CallToolResult, SubOutput, error) {

	pubsub := redishandler.Client.Subscribe(redishandler.Ctx, in.Channel)
	defer pubsub.Close()

	msg, err := pubsub.ReceiveMessage(redishandler.Ctx)

	if err != nil {
		return nil, SubOutput{}, fmt.Errorf("subscriber error: %v", err)
	}

	return nil, SubOutput{Messages: []string{msg.Payload}}, nil
}
