package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
type InitializeRequest struct {
	Request
	Params InitializeParams `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeParams
type InitializeParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeResult
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
type ServerCapabilities struct {
	// TextDocumentSyncKind: 0 - None, 1 - Full, 2 - Incremental
	TextDocumentSync   int                `json:"textDocumentSync"`
	CodeActionProvider bool               `json:"codeActionProvider"`
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionOptions
type CompletionOptions struct {
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1,
				CodeActionProvider: true,
				CompletionProvider: &CompletionOptions{},
			},
			ServerInfo: ServerInfo{
				Name:    "conventional-commit-lsp",
				Version: "0.0.1",
			},
		},
	}
}
