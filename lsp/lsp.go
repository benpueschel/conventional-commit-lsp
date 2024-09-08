package lsp

const ServerName = "conventional-commit-lsp"

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#requestMessage
type Request struct {
	// The JSON RPC version (2.0).
	RPC string `json:"jsonrpc"`
	// The request id.
	ID int `json:"id"`
	// The method to be invoked.
	Method string `json:"method"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage
type Response struct {
	// The JSON RPC version (2.0).
	RPC string `json:"jsonrpc"`
	// The request id.
	ID *int `json:"id,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#notificationMessage
type Notification struct {
	// The JSON RPC version (2.0).
	RPC string `json:"jsonrpc"`
	// The method to be invoked.
	Method string `json:"method"`
}
