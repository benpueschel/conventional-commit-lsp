package analysis

import (
	"strings"
	"unicode"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

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

	commit_type, description, found := strings.Cut(header, ": ")
	if !found {
		line_range := LineRange(0, 0, 0, len(header))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Commit message must contain a type and description, separated by a colon and space. ",
		})
		return diagnostics
	}

	diagnostics = append(diagnostics, getTypeDiagnostic(commit_type)...)

	if strings.TrimSpace(description) == "" {
		line_range := LineRange(0, len(commit_type)+1, 0, len(header))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Commit message must contain a description",
		})
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

	return diagnostics
}

func getTypeDiagnostic(commit_type string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	if strings.Contains(commit_type, " ") {
		line_range := LineRange(0, 0, 0, len(commit_type))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Commit type must not contain spaces",
		})
	}

	commit_type, scope, found := strings.Cut(commit_type, "(")
	if !found {
		commit_type, breaking, _ := strings.Cut(commit_type, "!")
		idx := len(commit_type) + len(scope) + 1
		diagnostics := checkBreakingHeaderDiagnostic(breaking, idx, diagnostics)
		if !isAlphabetic(commit_type) {
			line_range := LineRange(0, 0, 0, len(commit_type))
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Severity: 1,
				Range:    line_range,
				Source:   "conventional-commit-lsp",
				Message:  "Commit type must be alphabetic",
			})
		}
		return diagnostics
	}

	scope, breaking, found := strings.Cut(scope, ")")
	if !found {
		idx := len(commit_type) + len(scope) + 1
		line_range := LineRange(0, idx, 0, idx+1)
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Unclosed scope. Insert a closing parenthesis",
		})
	}

	idx := len(commit_type) + len(scope) + 2
	diagnostics = checkBreakingHeaderDiagnostic(breaking, idx, diagnostics)

	if !isAlphabetic(commit_type) {
		line_range := LineRange(0, 0, 0, len(commit_type))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Commit type must be alphabetic",
		})
	}

	if !isAlphabetic(scope) {
		idx := len(commit_type) + len(scope) + 1
		line_range := LineRange(0, len(commit_type)+1, 0, idx)
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Scope must be alphabetic",
		})
	}

	return diagnostics
}

func checkBreakingHeaderDiagnostic(breaking string, idx int, diagnostics []lsp.Diagnostic) []lsp.Diagnostic {
	if breaking != "" && breaking != "!" {
		line_range := LineRange(0, idx, 0, idx+len(breaking))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   "conventional-commit-lsp",
			Message:  "Breaking change indicator must be '!'",
		})
	}
	return diagnostics
}

func isAlphabetic(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
