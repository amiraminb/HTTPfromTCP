package request

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HTTPVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

var (
	ErrBadRequestLine = errors.New("bad requestline")
	HTTPPrefix        = "HTTP/"
	HTTPSeparator     = "\r\n"
)

func ParseRequestLine(reqLine string) (*RequestLine, error) {
	sections := strings.Fields(reqLine)
	if len(sections) != 3 {
		return nil, ErrBadRequestLine
	}

	hasPrefix := strings.HasPrefix(sections[2], HTTPPrefix)
	if !hasPrefix {
		return nil, ErrBadRequestLine
	}

	return &RequestLine{
		HTTPVersion:   strings.TrimPrefix(sections[2], HTTPPrefix),
		RequestTarget: sections[1],
		Method:        sections[0],
	}, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read request: %w", err)
	}

	reqParts := strings.Split(string(req), HTTPSeparator)
	reqLine, err := ParseRequestLine(reqParts[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse request: %w", err)
	}

	request := Request{
		RequestLine: *reqLine,
	}

	return &request, nil
}
