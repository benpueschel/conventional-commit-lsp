package rpc_test

import (
	"testing"
	"github.com/benpueschel/conventional-commit-lsp/rpc"
)

type EncodingExample struct {
	Method string `json:"method"`
}
var incommingMessage = "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"
func TestEncodeMessage(t *testing.T) {
	actual := rpc.EncodeMessage(EncodingExample{Method: "hi"})
	if incommingMessage != actual {
		t.Fatalf("Expected %s, got %s", incommingMessage, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	message, content, err := rpc.DecodeMessage([]byte(incommingMessage))
	if err != nil {
		t.Fatalf("Error decoding message: %s", err)
	}
	if message.Method != "hi" {
		t.Fatalf("Expected %s, got %s", "hi", message.Method)
	}
	if len(content) != 15 {
		t.Fatalf("Expected %d, got %d", 15, len(content))
	}
}
