package main

import (
	"context"
	"fmt"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	// Create the MCP server
	s := mcp.NewServer()

	// Run the server over stdio transport
	if err := s.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "Error running server: %v\n", err)
		os.Exit(1)
	}
}
