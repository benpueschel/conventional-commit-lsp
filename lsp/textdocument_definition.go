package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition
type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#definitionParams
type DefinitionParams struct {
	TextDocumentPositionParams
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition
type DefinitionResponse struct {
	Response
	Result []Location `json:"result"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#location
// Represents a location inside a resource, such as a line inside a text file.
type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}
