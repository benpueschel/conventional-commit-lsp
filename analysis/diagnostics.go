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

