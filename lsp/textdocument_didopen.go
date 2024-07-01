package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didOpen
type TextDocumentDidOpenNotification struct {
	Notification
	Params TextDocumentDidOpenParams `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didOpenTextDocumentParams
type TextDocumentDidOpenParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}
