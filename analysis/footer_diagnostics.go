package analysis

import (
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func getFooterDiagnostics(line string, row int) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}

	if len(line) == 0 {
		return diagnostics
	}

	if len(line) > 72 {
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 2,
			Range:    LineRange(row, 72, row, len(line)),
			Source:   "conventional-commit-lsp",
			Message:  "Commit footer should be wrapped at 72 characters.",
		})
	}

	token, value, found := strings.Cut(line, ": ")
	if !found || len(token) == 0 || len(value) == 0 {
		// INFO: Footers can contain newlines, so we cannot enforce every line
		// to be in the format of 'Footer-key: string value'. Basically, we can't
		// be sure if the user is intending to write a multiline footer or just isn't
		// able to write a simple list of key-value pair footers correctly. 
		// So, we can't enforce this rule.
		/*
		diagnostics = append(diagnostics, lsp.Diagnostic{
			Severity: 1,
			Range:    LineRange(row, 0, row, len(line)),
			Source:   "conventional-commit-lsp",
			Message:  "Commit footer must be in the format of 'Footer-key: string value'.",
		})
		*/
		return diagnostics

	}

	if token != "BREAKING CHANGE" && token != "BREAKING-CHANGE" {
		line_range := LineRange(row, 0, row, len(token))
		if strings.ToUpper(token) == "BREAKING CHANGE" {
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Severity: 1,
				Range:    line_range,
				Source:   "conventional-commit-lsp",
				Message:  "BREAKING CHANGE must be uppercase.",
			})
		} else if strings.ToUpper(token) == "BREAKING-CHANGE" {
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Severity: 1,
				Range:    line_range,
				Source:   "conventional-commit-lsp",
				Message:  "BREAKING-CHANGE must be uppercase.",
			})
		} else if strings.Contains(token, " ") {
			// INFO: see above
			/* 
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Severity: 1,
				Range:    line_range,
				Source:   "conventional-commit-lsp",
				Message:  "Footer tokens must not contain spaces. Use hyphens instead.",
			})
			*/
		} else if strings.Contains(token, "_") {
			// INFO: see above
			/*
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Severity: 2,
				Range:    line_range,
				Source:   "conventional-commit-lsp",
				Message:  "Footer tokens should not contain underscores. Use hyphens instead.",
			})
			*/
		}
	}

	return diagnostics
}
