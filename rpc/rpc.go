package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Encodes any json-serializable value into a string that can be sent over a stream.
// The format as demanded by the LSP specification is:
//
// Content-Length: <length>\r\n\r\n<content>
// where '<length>'  is the length of the content in bytes
// and   '<content>' is the json-serialized value.
func EncodeMessage(message any) string {
	content, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

// Decodes a slice of bytes in the JSON-RPC 2.0 format into a BaseMessage and returns the message content.
// Note that this function does not touch the content itself, just the message method (e.g. "textDocument/didOpen").
// The content is returned as a slice of bytes, the caller is responsible for unmarshalling it.
// The format as demanded by the LSP specification is:
//
// "Content-Length: <length>\r\n\r\n<content>"
// where '<length>'  is the length of the content in bytes
// and   '<content>' is the json-serialized value.
func DecodeMessage(msg []byte) (BaseMessage, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return BaseMessage{}, nil, errors.New("No separator found")
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return BaseMessage{}, nil, err
	}
	content = content[:contentLength]
	var message BaseMessage
	err = json.Unmarshal(content, &message)
	if err != nil {
		return message, nil, err
	}

	return message, content, nil
}
