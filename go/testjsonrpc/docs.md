# MCP in Go

This file encapsulates communication with MCP servers.

Here is its public interface:

```go
fetchServer := mcp.NewServer("uvx", "mcp-server-fetch")
err := fetchServer.Start() // starts the subprocess, and takes care of the initialization handshake
tools, err := fetchServer.ListTools() // returns list of tool definitions
fetchTool, found := tools.ByName("fetch")
fetchTool.Name // "fetch"
fetchTool.Description // "Long description of when to use the fetch tool"
fetchTool.RawInputSchema // raw json bytes of the input schema
content, err := fetchServer.Call("fetch", map[string]any{...}) // content is []mcp.ToolResultContent
fetchServer.Close() // shut down the server

type ToolResultContent struct {
  Type string // "text" or "image"
  Text string // non-empty when type == "text"
  Data string // non-empty base64 encoded data when type == "image"
  MimeType string // non-empty mime type for type == "image"
}
```

## Tools

### Listing tools - `tools/list`

Request:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/list",
  "params": {
    "cursor": "optional-cursor-value"
  }
}
```

Response:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "tools": [
      {
        "name": "get_weather",
        "description": "Get current weather information for a location",
        "inputSchema": {
          "type": "object",
          "properties": {
            "location": {
              "type": "string",
              "description": "City name or zip code"
            }
          },
          "required": ["location"]
        }
      }
    ],
    "nextCursor": "next-page-cursor"
  }
}
```

### Calling tools - `tools/call`

Request:

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/call",
  "params": {
    "name": "get_weather",
    "arguments": {
      "location": "New York"
    }
  }
}
```

Response:

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Current weather in New York:\nTemperature: 72°F\nConditions: Partly cloudy"
      }
    ],
    "isError": false
  }
}
```

## Setup

1. Set up subprocess & create I/O pipe to send/receive JSON-RPC to the subprocess
2. Set up ReadWriteCloser as an adapter to write
3. Set up codec (transition layer between high-level RPC interface and low-level network protocol) `rpc.NewClientWithCodec` to convert Go RPC calls into JSON-RPC 2.0 format and parse it back
4. Prepare request (params, method to call, etc.)

## Data flow

```
Go Application Layer:     client.Call("tools/list", params, &reply)
                                    ↓
RPC Framework:           rpc.Request{ServiceMethod: "tools/list", Seq: 123}
                                    ↓
Custom Codec:            {"jsonrpc":"2.0","method":"tools/list","id":123}
                                    ↓
Transport Layer:         JSON bytes over stdin/stdout pipes
                                    ↓
MCP Server:              Processes JSON-RPC 2.0 message
```

Request path (Go -> Subprocess)

1. Client with codec invokes `Call()` (Go RPC layer)
2. Codec serializes the JSON-RPC 2.0 request
3. Sub-process writes the request to stdin
4. MCP server processes the request

Response path (Subprocess -> Go)

1. MCP server sends JSON-RPC response to stdout of subprocess
2. `ReadCloser()` display on console + process RPC
3. Codec deserializes the response

> We implement request & body mutexes to prevent race condition when multiple goroutines make RPC calls simultaneously

## Translation responsibilities

`WriteRequest()` converts Go structures to JSON-RPC 2.0 (invoked by `Call()`)

- Go method call `client.Call("tools/list", params, &reply)`
- **Into JSON-RPC 2.0:** `{"jsonrpc":"2.0","method":"tools/list","params":{},"id":1}

## Notifications

We need to send a notification for a request we've just successfully made as required my MCP specification
