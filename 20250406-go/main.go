package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"read_ch32v003_datasheet",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// CH32V003 Datasheet
	calculatorTool := mcp.NewTool("read_ch32v003_datasheet",
		mcp.WithDescription("CH32V003のデータシートを取得"),
		// mcp.WithString("operation",
		// 	mcp.Required(),
		// 	mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
		// 	mcp.Enum("add", "subtract", "multiply", "divide"),
		// ),
	)

	s.AddTool(calculatorTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		f, err := os.Open("/Users/nnyn/ghq/github.com/74th/test-mcp_server/20250406-go/datasheet/ch32v003-ds.txt")
		if err != nil {
			return nil, fmt.Errorf("failed to open datasheet: %w", err)
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("failed to read datasheet: %w", err)
		}

		return mcp.NewToolResultText(string(data)), nil
	})

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
