package analysis

import (
	"strings"
	"unicode"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func (s *State) GetCodeActions(request lsp.CodeActionRequest) lsp.CodeActionResponse {
	actions := []lsp.CodeAction{}

	for _, diagnostic := range request.Params.Context.Diagnostics {
		if diagnostic.Source != lsp.ServerName {
			continue
		}
		if diagnostic.Data == nil {
			continue
		}
		switch diagnostic.Data.DiagnosticType {
		case lsp.CommitMessageHeaderMissing:
			actions = append(actions, s.generateHeaderAction(request, diagnostic.Range))
		case lsp.CommitMessageTypeNotAlphabetical:
			actions = append(actions, s.removeNonAlphabeticalAction(request, diagnostic.Range))
		case lsp.CommitMessageScopeNotAlphabetical:
			actions = append(actions, s.removeNonAlphabeticalAction(request, diagnostic.Range))
		case lsp.CommitMessageNoNewlineAfterHeader:
			actions = append(actions, s.insertNewlineAction(request, diagnostic.Range))
		case lsp.CommitMessageHeaderBreakingInvalid:
			actions = append(actions, s.removeExtraCharsAction(request, diagnostic.Range))
		case lsp.CommitMessageBreakingChangeCaseInvalid:
			actions = append(actions, s.toUppercaseAction(request, diagnostic.Range))
		}
	}

	return lsp.CodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &request.Request.ID,
		},
		Result: actions,
	}
}

// Inserts a newline before the specified range
func (s *State) insertNewlineAction(request lsp.CodeActionRequest, range_ lsp.Range) lsp.CodeAction {
	return lsp.CodeAction{
		Title: "Insert newline",
		Edit: &lsp.WorkspaceEdit{
			Changes: map[string][]lsp.TextEdit{
				request.Params.TextDocument.URI: {
					{
						Range: lsp.Range{
							Start: lsp.Position{
								Line:      range_.Start.Line,
								Character: 0,
							},
							End: lsp.Position{
								Line:      range_.Start.Line,
								Character: 0,
							},
						},
						NewText: "\n",
					},
				},
			},
		},
	}
}

// Converts the text in the specified range to uppercase
func (s *State) toUppercaseAction(request lsp.CodeActionRequest, range_ lsp.Range) lsp.CodeAction {
	text := GetText(s.Documents[request.Params.TextDocument.URI], range_)
	newText := strings.ToUpper(text)
	var title string
	if strings.Contains(text, "-") {
		title = "Convert BREAKING-CHANGE to uppercase"
	} else {
		title = "Convert BREAKING CHANGE to uppercase"
	}
	return lsp.CodeAction{
		Title: title,
		Edit: &lsp.WorkspaceEdit{
			Changes: map[string][]lsp.TextEdit{
				request.Params.TextDocument.URI: {
					{
						Range:   range_,
						NewText: newText,
					},
				},
			},
		},
	}
}

// Removes all characters in the specified range
func (s *State) removeExtraCharsAction(request lsp.CodeActionRequest, range_ lsp.Range) lsp.CodeAction {
	return lsp.CodeAction{
		Title: "Remove invalid characters",
		Edit: &lsp.WorkspaceEdit{
			Changes: map[string][]lsp.TextEdit{
				request.Params.TextDocument.URI: {
					{
						Range:   range_,
						NewText: "",
					},
				},
			},
		},
	}
}

// Generates a header in the specified range
func (s *State) generateHeaderAction(request lsp.CodeActionRequest, range_ lsp.Range) lsp.CodeAction {
	// TODO: snippet support
	return lsp.CodeAction{
		Title: "Insert header",
		Edit: &lsp.WorkspaceEdit{
			Changes: map[string][]lsp.TextEdit{
				request.Params.TextDocument.URI: {
					{
						Range:   range_,
						NewText: "type(scope): description",
					},
				},
			},
		},
	}
}

// Removes all non-alphabetical characters in the specified range
func (s *State) removeNonAlphabeticalAction(request lsp.CodeActionRequest, range_ lsp.Range) lsp.CodeAction {
	text := GetText(s.Documents[request.Params.TextDocument.URI], range_)
	newText := ""
	for _, c := range text {
		if unicode.IsLetter(c) {
			newText += string(c)
		}
	}

	return lsp.CodeAction{
		Title: "Remove non-alphabetical characters",
		Edit: &lsp.WorkspaceEdit{
			Changes: map[string][]lsp.TextEdit{
				request.Params.TextDocument.URI: {
					{
						Range:   range_,
						NewText: newText,
					},
				},
			},
		},
	}
}
