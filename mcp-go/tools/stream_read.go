package tools

import (
	"context"
	"fmt"

	redishandler "github.com/harshduche/memtxt/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type StreamReadInput struct {
	Stream string `json:"stream"`
	Count  int64  `json:"count"`
}

type StreamReadOutput struct {
	Messages []map[string]string `json:"messages"`
}

func StreamRead(ctx context.Context, req *mcp.CallToolRequest, in StreamReadInput) (
	*mcp.CallToolResult,
	StreamReadOutput,
	error,
) {
	res, err := redishandler.Client.XRevRangeN(redishandler.Ctx, in.Stream, "+", "-", in.Count).Result()
	if err != nil {
		return nil, StreamReadOutput{}, fmt.Errorf("stream read error: %v", err)
	}

	var messages []map[string]string
	for _, msg := range res {
		messages = append(messages, msg.Values)
	}

	return nil, StreamReadOutput{Messages: messages}, nil
}

