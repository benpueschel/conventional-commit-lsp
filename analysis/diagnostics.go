package analysis

import (
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

func GetDiagnostics(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}

	sections := strings.Split(text, "\n\n")
	totalSections := len(sections)
	currentLine := -1
	for index, section := range sections {
		lines := strings.Split(section, "\n")
		if index == 0 {
			// header
			isHeader := true
			for _, line := range lines {
				currentLine++
				// line is a comment, skip
				if strings.Index(line, "#") == 0 {
					continue
				}
				if isHeader {
					isHeader = false
					diagnostics = append(diagnostics, getHeaderDiagnostics(line, currentLine)...)
				} else if len(line) > 0 {
					diagnostics = append(diagnostics, lsp.Diagnostic{
						Severity: 1,
						Range:    LineRange(currentLine, 0, currentLine, len(line)),
						Source:   "conventional-commit-lsp",
						Message:  "Commit body must begin one blank line after the header. This line should be empty.",
						Data: &lsp.DiagnosticData{
							DiagnosticType: lsp.CommitMessageNoNewlineAfterHeader,
						},
					})
				}
			}
		} else if totalSections > 2 && index == totalSections-1 {
			// footer
			for _, line := range lines {
				currentLine++
				// line is a comment, skip
				if strings.Index(line, "#") == 0 {
					continue
				}
				diagnostics = append(diagnostics, getFooterDiagnostics(line, currentLine)...)
			}
		} else {
			// body
			for _, line := range lines {
				currentLine++
				if len(line) > 72 {
					diagnostics = append(diagnostics, lsp.Diagnostic{
						Severity: 2,
						Range:    LineRange(currentLine, 72, currentLine, len(line)),
						Source:   "conventional-commit-lsp",
						Message:  "Commit body shoyld be wrapped at 72 characters.",
					})
				}
			}
		}
		currentLine++
	}

	return diagnostics
}
