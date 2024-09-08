package analysis

import (
	"strings"
	"unicode"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func getHeaderDiagnostics(header string, row int) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	if len(header) == 0 {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    LineRange(row, 0, row, 0),
			Source:   lsp.ServerName,
			Message:  "Commit message must not be empty",
			Data: &lsp.DiagnosticData{
				DiagnosticType: lsp.CommitMessageHeaderMissing,
			},
		})
		return diagnostics
	} else if len(header) > 72 {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    LineRange(row, 72, row, len(header)),
			Source:   lsp.ServerName,
			Message:  "Commit message must be less than 72 characters",
		})
	} else if len(header) > 50 {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 2,
			Range:    LineRange(row, 50, row, len(header)),
			Source:   lsp.ServerName,
			Message:  "Commit message should be less than 50 characters",
		})
	}

	commit_type, description, found := strings.Cut(header, ": ")
	if !found {
		line_range := LineRange(row, 0, row, len(header))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   lsp.ServerName,
			Message:  "Commit message must contain a type and description, separated by a colon and space. ",
		})
		return diagnostics
	}

	diagnostics = append(diagnostics, getTypeDiagnostic(commit_type, row)...)

	if strings.TrimSpace(description) == "" {
		line_range := LineRange(row, len(commit_type)+1, row, len(header))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   lsp.ServerName,
			Message:  "Commit message must contain a description",
		})
	}

	if commit_type == "" {
		line_range := LineRange(row, 0, row, len(header)-1)
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   lsp.ServerName,
			Message:  "Commit message must contain a type",
		})
	}

	return diagnostics
}

func getTypeDiagnostic(commit_type string, row int) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	if strings.Contains(commit_type, " ") {
		line_range := LineRange(row, 0, row, len(commit_type))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   lsp.ServerName,
			Message:  "Commit type must not contain spaces",
		})
	}

	commit_type, scope, found := strings.Cut(commit_type, "(")
	if !found {
		commit_type, breaking, _ := strings.Cut(commit_type, "!")
		idx := len(commit_type) + len(scope) + 1
		diagnostics := checkBreakingHeaderDiagnostic(breaking, idx, row, diagnostics)
		if !isAlphabetic(commit_type) {
			line_range := LineRange(row, 0, row, len(commit_type))
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Severity: 1,
				Range:    line_range,
				Source:   lsp.ServerName,
				Message:  "Commit type must be alphabetic",
				Data: &lsp.DiagnosticData{
					DiagnosticType: lsp.CommitMessageTypeNotAlphabetical,
				},
			})
		}
		return diagnostics
	}

	scope, breaking, found := strings.Cut(scope, ")")
	if !found {
		idx := len(commit_type) + len(scope) + 1
		line_range := LineRange(row, idx, row, idx+1)
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   lsp.ServerName,
			Message:  "Unclosed scope. Insert a closing parenthesis",
		})
	}

	idx := len(commit_type) + len(scope) + 2
	diagnostics = checkBreakingHeaderDiagnostic(breaking, idx, row, diagnostics)

	if !isAlphabetic(commit_type) {
		line_range := LineRange(row, 0, row, len(commit_type))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   lsp.ServerName,
			Message:  "Commit type must be alphabetic",
			Data: &lsp.DiagnosticData{
				DiagnosticType: lsp.CommitMessageTypeNotAlphabetical,
			},
		})
	}

	return diagnostics
}

func checkBreakingHeaderDiagnostic(breaking string, idx int, row int, diagnostics []lsp.Diagnostic) []lsp.Diagnostic {
	if breaking != "" && breaking != "!" {
		line_range := LineRange(row, idx, row, idx+len(breaking))
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    line_range,
			Source:   lsp.ServerName,
			Message:  "Breaking change indicator must be '!'",
			Data: &lsp.DiagnosticData{
				DiagnosticType: lsp.CommitMessageHeaderBreakingInvalid,
			},
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
