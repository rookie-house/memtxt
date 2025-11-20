




func StreamRead(ctx context.Context, req *mcp.CallToolRequest, in StreamReadInput) (
	*mcp.CallToolResult,
	StreamReadOutput,
	error,
) {

	res, err := memredis.Client.XRevRangeN(memredis.Ctx, in.Stream, "+", "-", in.Count).Result()
	if err != nil {
		return nil, StreamReadOutput{}, fmt.Errorf("stream read error: %v", err)
	}

	messages := []map[string]string{}

	for _, msg := range res {
		converted := map[string]string{}
		for k, v := range msg.Values {
			converted[k] = fmt.Sprintf("%v", v)
		}
		messages = append(messages, converted)
	}

	return nil, StreamReadOutput{Messages: messages}, nil
}

