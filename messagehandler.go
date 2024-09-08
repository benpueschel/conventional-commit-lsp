package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/benpueschel/conventional-commit-lsp/analysis"
	"github.com/benpueschel/conventional-commit-lsp/lsp"
	"github.com/benpueschel/conventional-commit-lsp/rpc"
)

func handleBaseRequest(logger *log.Logger, contents []byte, request any, message string) error {
	if err := json.Unmarshal(contents, &request); err != nil {
		logger.Printf("%s: %s", message, err)
		return err
	}
	return nil
}

func initialize(logger *log.Logger, writer io.Writer, contents []byte) error {
	var request lsp.InitializeRequest
	if err := handleBaseRequest(logger, contents, &request, "initialize"); err != nil {
		return err
	}
	msg := lsp.NewInitializeResponse(request.ID)
	writeResponse(msg, writer)

	logger.Printf("Connected to client: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
	return nil
}

func textDocumentDidOpen(logger *log.Logger, writer io.Writer, state *analysis.State, contents []byte) error {
	var request lsp.TextDocumentDidOpenNotification
	if err := handleBaseRequest(logger, contents, &request, "textDocument/didOpen"); err != nil {
		return err
	}
	logger.Printf("Opened document: %s", request.Params.TextDocument.URI)
	diagnostics := state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	if diagnostics != nil {
		writeDiagnostics(request.Params.TextDocument.URI, diagnostics, writer)
	}
	return nil
}

func textDocumentDidChange(logger *log.Logger, writer io.Writer, state *analysis.State, contents []byte) error {
	var request lsp.TextDocumentDidChangeNotification
	if err := handleBaseRequest(logger, contents, &request, "textDocument/didChange"); err != nil {
		return err
	}
	logger.Printf("Changed document: %s", request.Params.TextDocument.URI)
	for _, change := range request.Params.ContentChanges {
		diagnostics := state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		if diagnostics != nil {
			writeDiagnostics(request.Params.TextDocument.URI, diagnostics, writer)
		}
	}
	return nil
}

func textDocumentCodeAction(logger *log.Logger, writer io.Writer, state *analysis.State, contents []byte) error {
	var request lsp.CodeActionRequest
	if err := handleBaseRequest(logger, contents, &request, "textDocument/codeAction"); err != nil {
		return err
	}

	response := state.GetCodeActions(request)
	writeResponse(response, writer)

	return nil
}

func textDocumentCompletion(logger *log.Logger, writer io.Writer, state *analysis.State, contents []byte) error {
	var request lsp.CompletionRequest
	if err := handleBaseRequest(logger, contents, &request, "textDocument/completion"); err != nil {
		return err
	}
	response := state.GetCompletions(request)
	writeResponse(response, writer)
	return nil
}

func textDocumentDefinition(logger *log.Logger, writer io.Writer, state *analysis.State, contents []byte) error {
	var request lsp.DefinitionRequest
	if err := handleBaseRequest(logger, contents, &request, "textDocument/definition"); err != nil {
		return err
	}
	definitions, err := state.GetDefinitions(logger, request)
	if err != nil {
		logger.Printf("Failed to get definitions: %s", err)
		return err
	}
	response := lsp.DefinitionResponse{
		Response: lsp.Response{
			ID: &request.ID,
			RPC: "2.0",
		},
		Result: definitions,
	}
	writeResponse(response, writer)
	return nil
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
