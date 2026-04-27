# Chapter 4: Request Lines

Course focus: parse HTTP request lines from a stream of data.

## Lecture Goal

The request line is the first line of an HTTP request. It tells the server the operation, the target resource, and the protocol version. Parsing it correctly creates the first useful HTTP object from the raw TCP stream.

## Request Line Grammar

```text
METHOD SP request-target SP HTTP-version CRLF
```

Example:

```text
GET /coffee HTTP/1.1\r\n
```

The three fields are the method, request target, and HTTP version.

## Lecture Progression

1. Parse a complete request line from a full request string.
2. Validate that the line has exactly three sections.
3. Extract `Method`, `RequestTarget`, and `HttpVersion` into a `RequestLine` struct.
4. Refactor request-line parsing so it works on bytes and reports how many bytes were consumed.
5. Add parser states so incomplete data can wait for more bytes.
6. Connect the request-line parser to the live TCP listener.

## Core Learning

The request line ends at the first `\r\n`. Until that delimiter is present, the parser does not have enough data to decide whether the request line is valid.

The parser therefore has three possible outcomes:

1. A full valid request line was parsed.
2. More bytes are needed before a decision can be made.
3. The line is malformed and parsing should fail.

This is why request-line parsing needs to report both the parsed result and how much input it consumed.

## Practical Takeaways

- Returning `0` bytes consumed can mean the parser needs more data.
- Returning the consumed byte count lets the caller preserve unread bytes for later parser states.
- Supporting only `HTTP/1.1` is acceptable at this stage because the course is focused on HTTP/1.1.
- Incremental parsing matters because request lines can arrive split across multiple reads.
