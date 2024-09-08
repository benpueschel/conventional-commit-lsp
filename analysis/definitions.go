package analysis

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/benpueschel/conventional-commit-lsp/lsp"
)

// getRootDir returns the root directory of a git repository given a COMMIT_EDITMSG file.
func getRootDir(uri string) (string, error) {
	_, p, found := strings.Cut(uri, "file://")
	if !found {
		return "", fmt.Errorf("URI is not a file URI: %s", uri)
	}
	git_dir := path.Dir(p)
	return path.Dir(git_dir), nil
}

func (s *State) GetDefinitions(logger *log.Logger, request lsp.DefinitionRequest) ([]lsp.Location, error) {
	definitions := []lsp.Location{}
	uri := request.Params.TextDocument.URI

	position := request.Params.Position
	text := s.Documents[uri]
	line := GetLine(text, position.Line)

	_, trim_line, found := strings.Cut(line, "#\t")
	logger.Printf("Looking up definitions in line: %s\n", line)

	// In commit messages, all changed files are listed in comments
	// at the end of the message. If the line is not a comment,
	// we can't provide a definition.
	if !found {
		return definitions, nil
	}
	// The line is a comment, so we can check if it's part of the
	// auto-generated list of changed files.
	// NOTE: This is a very naive implementation. It assumes that
	// the line is a comment and that the comment is a list of
	// changed files. This is not always the case, but it's a good
	// starting point.

	logger.Printf("Looking up definitions in line: %s\n", trim_line)
	if strings.Index(text, "Changes to be committed:") < strings.Index(text, line) {
		trim_line = strings.ReplaceAll(trim_line, "modified: ", "")
		trim_line = strings.ReplaceAll(trim_line, "new file: ", "")
		trim_line = strings.ReplaceAll(trim_line, "deleted: ", "")
		trim_line = strings.TrimSpace(trim_line)
		rootDir, err := getRootDir(uri)
		if err != nil {
			return definitions, err
		}
		filePath := path.Join(rootDir, trim_line)
		fileURI := fmt.Sprintf("file://%s", filePath)
		fileURI = strings.ReplaceAll(fileURI, " ", "%20") // URI encode spaces
		logger.Printf("Checking for file: %s\n", filePath)
		logger.Printf("Checking for file URI: %s\n", fileURI)

		if _, err := os.Stat(filePath); err == nil {
			definitions = append(definitions, lsp.Location{
				URI: fileURI,
				Range: lsp.Range{
					Start: lsp.Position{
						Line:      0,
						Character: 0,
					},
					End: lsp.Position{
						Line:      0,
						Character: 1,
					},
				},
			})
		}

	}

	return definitions, nil
}
