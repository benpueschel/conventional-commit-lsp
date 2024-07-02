package analysis

import (
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func (s *State) getHeaderCompletions(request lsp.CompletionRequest, line string) []lsp.CompletionItem {
	results := []lsp.CompletionItem{}

	position := request.Params.Position
	colonIndex := strings.Index(line, ":")
	parenIndex := strings.Index(line, "(")
	isBeforeColon := colonIndex == -1 || colonIndex > position.Character
	isBeforeParen := parenIndex == -1 || parenIndex > position.Character

	// If the cursor is before the colon and before an optional paren 
	// (indicating a scope), we can suggest commit types.
	if isBeforeColon && isBeforeParen {
		results = []lsp.CompletionItem{{
			Label:         "feat",
			Detail: 	   "Commit adds or removes a new feature.",
			Documentation: "",
		}, {
			Label:         "fix",
			Detail:        "Commit fixes a bug.",
			Documentation: "",
		}, {
			Label:         "docs",
			Detail:        "Commit only affects documentation.",
			Documentation: "",
		}, {
			Label:         "ci",
			Detail:        "Commit changes CI configuration.",
			Documentation: "",
		}, {
			Label:         "style",
			Detail:        "Commit changes code style and does not affect the code semantics.",
			Documentation: "",
		}, {
			Label:         "refactor",
			Detail:        "Commit rewrites/refactors code without changing any API behavior.",
			Documentation: "",
		}, {
			Label:         "revert",
			Detail:        "Commit reverts previous commit(s).",
			Documentation: "",
		}, {
			Label:         "perf",
			Detail:        "Commit improves performance without changing any API behavior.",
			Documentation: "",
		}, {
			Label:         "test",
			Detail:        "Commit adds or modifies tests.",
		}}
	}
	return results
}

func (s *State) getBodyCompletions(request lsp.CompletionRequest, line string) []lsp.CompletionItem {
	results := []lsp.CompletionItem{}

	position := request.Params.Position
	colonIndex := strings.Index(line, ":")
	spaceIndex := strings.Index(line, " ")
	isBeforeColon := colonIndex == -1 || colonIndex > position.Character
	isBeforeSpace := spaceIndex == -1 || spaceIndex > position.Character

	// If the cursor is before the colon and before a space, 
	// we can suggest footers.
	if isBeforeColon && isBeforeSpace {
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
	return results
}

func (s *State) GetCompletions(request lsp.CompletionRequest) lsp.CompletionResponse {
	results := []lsp.CompletionItem{}

	text := s.Documents[request.Params.TextDocument.URI]
	position := request.Params.Position
	line := GetLine(text, position.Line)

	if IsHeader(text, position.Line) {
		results = s.getHeaderCompletions(request, line)
	} else {
		results = s.getBodyCompletions(request, line)
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &request.ID,
		},
		Result: results,
	}
}
