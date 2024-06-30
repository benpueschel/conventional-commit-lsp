package lsp

type TextDocumentDidChangeNotification struct {
	Notification
	Params TextDocumentDidChangeParams `json:"params"`
}

type TextDocumentDidChangeParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TODO: handle incremental updates
type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
