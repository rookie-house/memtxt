package tools

import (
	"context"
	"fmt"

	"github.com/harshduche/memtxt/redis"
	"github.com/modelcontextprotocol/go-sdk/mcp"
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
	id, err := redis.Client.XAdd(redis.Ctx, &redis.XAddArgs{
		Stream: in.Stream,
		Values: in.Fields,
	}).Result()

	if err != nil {
		return nil, StreamWriteOutput{}, fmt.Errorf("stream write error: %v", err)
	}

	return nil, StreamWriteOutput{ID: id}, nil
}

