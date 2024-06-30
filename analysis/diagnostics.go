package analysis

import (
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func GetDiagnostics(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}

	isHeader := true
	isBody := false

	for row, line := range strings.Split(text, "\n") {
		if strings.Index(line, "#") == 0 {
			continue
		}

		if isHeader {
			isHeader = false
			diagnostics = append(diagnostics, getHeaderDiagnostics(line, row)...)
			continue
		}

		if !isBody {
			isBody = true
			if len(line) > 0 {
				diagnostics = append(diagnostics, lsp.Diagnostic{
					Severity: 1,
					Range:    LineRange(row, 0, row, len(line)),
					Source:   "conventional-commit-lsp",
					Message:  "Commit body must begin one blank line after the header. This line should be empty.",
				})
			}
		}

	}
	return diagnostics
}

