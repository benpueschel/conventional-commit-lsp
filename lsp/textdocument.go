package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentItem
type TextDocumentItem struct {
	// The text document's URI.
	URI string `json:"uri"`
	// The text document's language identifier.
	LanguageID string `json:"languageId"`
	// The version number of this document (it will strictly increase after each change, including undo/redo).
	Version int `json:"version"`
	// The content of the opened text document.
	Text string `json:"text"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#versionedTextDocumentIdentifier
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	// The version number of this document. If a versioned text document identifier
	Version int `json:"version"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentIdentifier
type TextDocumentIdentifier struct {
	// The text document's URI.
	URI string `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentPositionParams
type TextDocumentPositionParams struct {
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// The position inside the text document.
	Position Position `json:"position"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#position
type Position struct {
	// Line position in a document (zero-based).
	Line int `json:"line"`
	// Character offset on a line in a document (zero-based).
	Character int `json:"character"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#range
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspaceEdit
type WorkspaceEdit struct {
	// Holds changes to existing resources.
	// Maps DocumentURIs to arrays of TextEdits.
	Changes map[string][]TextEdit `json:"changes"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textEdit
type TextEdit struct {
	// The range of the text document to be manipulated. To insert
	// text into a document create a range where start === end.
	Range Range `json:"range"`
	// The string to be inserted. For delete operations use an
	// empty string.
	NewText string `json:"newText"`
}
