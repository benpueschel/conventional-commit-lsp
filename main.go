package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/benpueschel/conventional-commit-lsp/analysis"
	"github.com/benpueschel/conventional-commit-lsp/lsp"
	"github.com/benpueschel/conventional-commit-lsp/rpc"
)

func main() {
	logger := getLogger("log.txt")
	logger.Println("Starting server")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		message, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Failed to decode message: %s", err)
			continue
		}
		handleMessage(logger, writer, &state, message, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state *analysis.State, msg rpc.BaseMessage, contents []byte) {
	logger.Printf("Received message with method: %s", msg.Method)

	switch msg.Method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("initialize: %s", err)
			return
		}
		logger.Printf("Connected to client: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(msg, writer)
		logger.Println("Sent initialize response")
	case "textDocument/didOpen":
		var request lsp.TextDocumentDidOpenNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpen: %s", err)
			return
		}
		logger.Printf("Opened document: %s", request.Params.TextDocument.URI)
		diagnostics := state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
		if diagnostics != nil {
			writeDiagnostics(request.Params.TextDocument.URI, diagnostics, writer)
		}
	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}
		logger.Printf("Changed document: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			diagnostics := state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
			if diagnostics != nil {
				writeDiagnostics(request.Params.TextDocument.URI, diagnostics, writer)
			}
		}
	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/codeAction: %s", err)
			return
		}
		response := state.GetCodeActions(request)
		writeResponse(response, writer)
	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/completion: %s", err)
			return
		}
		response := state.GetCompletions(request)
		writeResponse(response, writer)
	}
}

func writeDiagnostics(uri string, diagnostics []lsp.Diagnostic, writer io.Writer) {
	writeResponse(lsp.PublishDiagnosticsNotification{
		Notification: lsp.Notification{
			RPC:    "2.0",
			Method: "textDocument/publishDiagnostics",
		},
		Params: lsp.PublishDiagnosticsParams{
			URI:         uri,
			Diagnostics: diagnostics,
		},
	}, writer)
}

func writeResponse(msg any, writer io.Writer) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	return log.New(file, "[conventional-commit-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
	if !found {
		// wait for more data, so do nothing
		return 0, nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		// wait for more data, so do nothing
		return 0, nil, nil
	}

	// header + content + \r\n\r\n
	totalLength := len(header) + len(content[:contentLength]) + 4
	return totalLength, data[:totalLength], nil
}
