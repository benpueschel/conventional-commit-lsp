package analysis

import (
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func GetDiagnostics(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		if row == 0 {
			diagnostics = append(diagnostics, getHeaderDiagnostics(line)...)
		}
	}
	return diagnostics
}

func getHeaderDiagnostics(header string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	if len(header) == 0 {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    LineRange(0, 0, 0, 0),
			Source:   "conventional-commit-lsp",
			Message:  "Commit message must not be empty",
		})
		return diagnostics
	} else if len(header) > 72 {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    LineRange(0, 72, 0, len(header)),
			Source:   "conventional-commit-lsp",
			Message:  "Commit message must be less than 72 characters",
		})
		return diagnostics
	} else if len(header) > 50 {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 2,
			Range:    LineRange(0, 50, 0, len(header)),
			Source:   "conventional-commit-lsp",
			Message:  "Commit message should be less than 50 characters",
		})
	}

	commit_type, description, found := strings.Cut(header, ":")
	if !found {
		line_range := LineRange(0, 0, 0, len(header))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Commit message must contain a type and description, separated by a colon. ",
		})
		return diagnostics
	}

	if strings.Contains(commit_type, " ") {
		line_range := LineRange(0, 0, 0, len(commit_type))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Commit type must not contain spaces",
		})
		return diagnostics
	}

	if commit_type == "" {
		line_range := LineRange(0, 0, 0, len(header)-1)
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Commit message must contain a type",
		})
	}
	_ = description

	return diagnostics
}
