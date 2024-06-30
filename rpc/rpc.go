package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

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
