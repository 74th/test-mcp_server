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
	readDatasheetTool := mcp.NewTool("read_ch32v003_datasheet",
		mcp.WithDescription("CH32V003のデータシートを取得"),
		mcp.WithString("mcu_name",
			mcp.Required(),
			mcp.Description("MCUの名前 (CH32V003)"),
			mcp.Enum("CH32V003"),
		),
	)

	s.AddTool(readDatasheetTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// mcuName := request.Params.Arguments["mcu_name"].(string)
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

	// // CH32V003 Reference Manual
	// referenceManualTool := mcp.NewTool("read_ch32v003_reference_manual",
	// 	mcp.WithDescription("CH32V003のリファレンスマニュアルを取得"),
	// 	mcp.WithString("mcu_name",
	// 		mcp.Required(),
	// 		mcp.Description("MCUの名前 (CH32V003)"),
	// 		mcp.Enum("CH32V003"),
	// 	),
	// )

	// s.AddTool(referenceManualTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 	// mcuName := request.Params.Arguments["mcu_name"].(string)
	// 	f, err := os.Open("/Users/nnyn/ghq/github.com/74th/test-mcp_server/20250406-go/datasheet/ch32v003-rm.txt")
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to open reference manual: %w", err)
	// 	}
	// 	defer f.Close()

	// 	data, err := io.ReadAll(f)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to read reference manual: %w", err)
	// 	}

	// 	return mcp.NewToolResultText(string(data)), nil
	// })

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
