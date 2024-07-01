package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_publishDiagnostics
type PublishDiagnosticsNotification struct {
	Notification
	Params PublishDiagnosticsParams `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#publishDiagnosticsParams
type PublishDiagnosticsParams struct {
	// The URI for which diagnostic information is reported.
	URI string `json:"uri"`
	// An array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnostic
type Diagnostic struct {
	Range Range `json:"range"`
	// 1 - Error, 2 - Warning, 3 - Information, 4 - Hint
	Severity int `json:"severity"`
	// A human-readable string describing the source of this
	// diagnostic ('conventional-commit-lsp' in our case)
	Source string `json:"source"`
	// A human-readable string describing the source of this
	Message string `json:"message"`
	// Extra data reserved for the server. We use this for code actions.
	Data *DiagnosticData `json:"data,omitempty"`
}

type DiagnosticData struct {
	DiagnosticType DiagnosticType `json:"diagnosticType"`
}

type DiagnosticType int

const (
	CommitMessageHeaderMissing DiagnosticType = iota
	CommitMessageTypeNotAlphabetical
	CommitMessageScopeNotAlphabetical
	CommitMessageNoNewlineAfterHeader
	CommitMessageHeaderBreakingInvalid
	CommitMessageBreakingChangeCaseInvalid
)
