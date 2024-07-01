package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didChange
type TextDocumentDidChangeNotification struct {
	Notification
	Params TextDocumentDidChangeParams `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#didChangeTextDocumentParams
type TextDocumentDidChangeParams struct {
	// The document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`
	// The actual content changes.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentContentChangeEvent
// TODO: handle incremental updates
type TextDocumentContentChangeEvent struct {
	// The new text of the whole document.
	Text string `json:"text"`
}
