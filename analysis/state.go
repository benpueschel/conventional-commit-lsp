package analysis

import "github.com/benpueschel/conventional-commit-lsp/lsp"

type State struct {
	// Map of document URIs to text contents
	Documents        map[string]string
	DocumentAnalysis map[string]DocumentAnalysis
}

type DocumentAnalysis struct {
	Type                 CommitMessageType
	Description          string
	Body                 *string
	Footers              []CommitMessageFooter
	BreakingChangeFooter *string
}

type CommitMessageType struct {
	Type           string
	Scope          *string
	BreakingChange bool
}

type CommitMessageFooter struct {
	Type  string
	Value string
}

func LineRange(startLine int, startCharacter int, endLine int, endCharacter int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      startLine,
			Character: startCharacter,
		},
		End: lsp.Position{
			Line:      endLine,
			Character: endCharacter,
		},
	}
}

func NewState() State {
	return State{
		Documents: make(map[string]string),
	}
}

func (s *State) OpenDocument(uri string, text string) []lsp.Diagnostic {
	return s.UpdateDocument(uri, text)
}

func (s *State) UpdateDocument(uri string, text string) []lsp.Diagnostic {
	// TODO: handle incremental updates
	s.Documents[uri] = text
	return GetDiagnostics(text)
}
