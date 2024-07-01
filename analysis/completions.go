package analysis

import (
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func (s *State) GetCompletions(request lsp.CompletionRequest) lsp.CompletionResponse {
	results := []lsp.CompletionItem{}

	text := s.Documents[request.Params.TextDocument.URI]
	position := request.Params.Position
	line := GetLine(text, position.Line)
	colonIndex := strings.Index(line, ":")
	isBeforeColon := colonIndex == -1 || colonIndex > position.Character

	if IsHeader(text, position.Line) {
		if isBeforeColon {
			results = []lsp.CompletionItem{{
				Label:         "feat",
				Detail:        "Indicates a commit implements a new feature.",
				Documentation: "",
			}, {
				Label:         "fix",
				Detail:        "Indicates a commit that fixes a bug.",
				Documentation: "",
			}, {
				Label:         "docs",
				Detail:        "Indicates a commit that changes documentation.",
				Documentation: "",
			}, {
				Label:         "ci",
				Detail:        "Indicates a commit that changes CI configuration.",
				Documentation: "",
			}, {
				Label:         "style",
				Detail:        "Indicates a commit that changes code style (whitespace, formatting, etc.,.",
				Documentation: "",
			}, {
				Label:         "refactor",
				Detail:        "Indicates a commit that refactors code.",
				Documentation: "",
			}, {
				Label:         "revert",
				Detail:        "Indicates a commit that reverts one ore more previous commits.",
				Documentation: "",
			}}
		}
	} else {
		if isBeforeColon {
			results = []lsp.CompletionItem{{
				Label:         "BREAKING CHANGE",
				Detail:        "Explain a breaking change this commit introduces.",
				Documentation: "",
			}, {
				Label:         "Closes",
				Detail:        "A list of issues this commit closes.",
				Documentation: "",
			}, {
				Label:         "Fixes",
				Detail:        "A list of issues this commit fixes.",
				Documentation: "",
			}, {
				Label:         "Refs",
				Detail:        "A list of issues this commit references.",
				Documentation: "",
			}, {
				Label:         "Signed-off-by",
				Detail:        "A list of sign-offs for the commit.",
				Documentation: "",
			}, {
				Label:         "Reviewed-by",
				Detail:        "A list of reviewers for the commit.",
				Documentation: "",
			}}
		}
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &request.ID,
		},
		Result: results,
	}
}
