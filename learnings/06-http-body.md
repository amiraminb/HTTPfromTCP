# Chapter 6: HTTP Body

Course focus: read and process the body of an HTTP request.

## Lecture Goal

The body is the optional payload after the request line and headers. It is where a client sends data such as form submissions, JSON payloads, or uploaded bytes. The important lesson is that the body needs framing: the parser must know exactly where it ends.

## Body Framing

Example request with a body:

```text
POST /submit HTTP/1.1\r\n
Host: localhost:42069\r\n
Content-Length: 13\r\n
\r\n
hello world!\n
```

The blank line only says the headers are finished. It does not define the size of the body. In this implementation, `Content-Length` provides that size.

## Lecture Progression

1. Extend the request structure with a `Body` field.
2. Read the `Content-Length` header after the header section is complete.
3. Enter `StateBody` only when the content length is greater than zero.
4. Append body bytes incrementally until the declared length has been reached.
5. Treat an early end of stream as an incomplete request.

## Core Learning

The body is not line-based. It can contain text, JSON, binary data, or newline characters. The parser should not look for `\n` to decide when the body is done. It should count bytes according to the framing metadata.

This is the reason `Content-Length` matters:

```text
Content-Length: 13
```

The parser must read exactly 13 body bytes after the header terminator. If fewer bytes arrive, the request is incomplete. If more bytes are already buffered, those bytes belong to whatever comes next and should not be swallowed as part of this body.

## Practical Takeaways

- Parse headers before attempting to parse the body.
- Use `Content-Length` as byte-count framing.
- Do not use line delimiters to parse a body.
- Body bytes can arrive split across multiple reads.
- `chunked` transfer encoding is separate framing and is not implemented yet.
- The body should be treated as bytes, not as lines.
