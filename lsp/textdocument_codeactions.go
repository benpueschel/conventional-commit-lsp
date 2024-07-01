package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_codeAction
type CodeActionRequest struct {
	Request
	Params CodeActionParams `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeActionParams
type CodeActionParams struct {
	// The document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// The range for which the command was invoked.
	Range Range `json:"range"`
	// Context carrying additional information.
	Context CodeActionContext `json:"context"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeActionContext
type CodeActionContext struct {
	// An array of diagnostics known on the client side overlapping the provided range.
	Diagnostics []Diagnostic `json:"diagnostics"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_codeAction
type CodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#codeAction
type CodeAction struct {
	// A short, human-readable, title for this code action.
	Title string `json:"title"`
	// The kind of the code action.
	// Used to filter code actions.
	Kind *string `json:"kind,omitempty"`
	// Marks this as a preferred action. Preferred actions are used by the
	// "auto fix" command and can be targeted by keybindings.
	// A quick fix should be marked preferred if it properly addresses the
	// underlying error. A refactoring should be marked preferred if it is the
	// most reasonable choice of actions to take.
	IsPreferred *bool `json:"isPreferred,omitempty"`
	// The workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitempty"`
}
