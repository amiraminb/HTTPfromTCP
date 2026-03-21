package request

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HTTPVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	reqParts := strings.Split(string(req), "\r\n")
	reqLine := reqParts[0]
	reqLineParts := strings.Split(string(reqLine), " ")
	if len(reqLineParts) != 3 {
		return nil, fmt.Errorf("Request line does not have 3 sections")
	}

	request := Request{
		RequestLine{
			HTTPVersion:   strings.TrimPrefix(reqLineParts[2], "HTTP/"),
			RequestTarget: reqLineParts[1],
			Method:        reqLineParts[0],
		},
	}

	return &request, nil
}
