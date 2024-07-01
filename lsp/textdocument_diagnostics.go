package lsp

type PublishDiagnosticsNotification struct {
	Notification
	Params PublishDiagnosticsParams `json:"params"`
}

type PublishDiagnosticsParams struct {
	URI         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Range Range `json:"range"`
	// 1 - Error, 2 - Warning, 3 - Information, 4 - Hint
	Severity int             `json:"severity"`
	Source   string          `json:"source"`
	Message  string          `json:"message"`
	Data     *DiagnosticData `json:"data,omitempty"`
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
