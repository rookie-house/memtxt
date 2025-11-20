package tools

import (
	"context"
	"fmt"

	redishandler "github.com/harshduche/memtxt/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/redis/go-redis/v9"
)


type StreamWriteInput struct {
	Stream string            `json:"stream"`
	Fields map[string]string `json:"fields"`
}

type StreamWriteOutput struct {
	ID string `json:"id"`
}

func StreamWrite(ctx context.Context, req *mcp.CallToolRequest, in StreamWriteInput) (
	*mcp.CallToolResult,
	StreamWriteOutput,
	error,
) {
	id, err := redishandler.Client.XAdd(redishandler.Ctx, &redis.XAddArgs{
		Stream: in.Stream,
		Values: in.Fields,
	}).Result()

	if err != nil {
		return nil, StreamWriteOutput{}, fmt.Errorf("stream write error: %v", err)
	}

	return nil, StreamWriteOutput{ID: id}, nil
}

