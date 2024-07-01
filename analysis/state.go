package analysis

import (
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

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

func GetText(text string, range_ lsp.Range) string {
	lines := strings.Split(text, "\n")
	newText := ""
	for i, line := range lines {
		if i >= range_.Start.Line && i <= range_.End.Line {
			// if we're on the first line, start at the character offset
			if i == range_.Start.Line {
				line = line[range_.Start.Character:]
			}
			// if we're on the last line, end at the character offset
			if i == range_.End.Line {
				line = line[:range_.End.Character]
			}
			newText += line + "\n"
		}
	}
	return newText
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

