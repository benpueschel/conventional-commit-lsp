package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionParams
type CompletionParams struct {
	TextDocumentPositionParams
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItem
type CompletionItem struct {
	// The label of this completion item.
	// The label property is also by default the text that
	// is inserted when selecting this completion.
	Label string `json:"label"`
	// Additional details for the label
	Detail string `json:"detail"`
	// A human-readable string that represents a doc-comment.
	Documentation string `json:"documentation"`
}
