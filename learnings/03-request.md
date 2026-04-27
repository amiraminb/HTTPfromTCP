# Chapter 3: Requests

Course focus: understand HTTP requests and how data is sent over the web.

## Lecture Goal

A TCP connection gives the server bytes. An HTTP request gives those bytes meaning. This chapter is the conceptual bridge from transport to protocol: a client sends a structured message that asks the server to do something with a resource.

## Request Structure

An HTTP request is ordered. The parser should expect the pieces in this sequence:

```text
request line\r\n
header line\r\n
header line\r\n
\r\n
optional body bytes
```

The empty line is not decoration. It is the separator between headers and the optional message body.

## Lecture Progression

1. Treat the whole incoming byte stream as one HTTP request.
2. Identify that the first meaningful section is the request line.
3. Add headers as the metadata section after the request line.
4. Add the optional body after the blank line.
5. Refactor toward a parser that can consume partial data incrementally.

## Core Learning

An HTTP request is a protocol message with framing rules. The parser needs to know when each section starts and ends. It cannot parse the body before parsing the headers because the headers explain whether a body exists and how long it is.

State is important because a stream parser may receive only part of a request at a time. The parser needs to remember whether it is currently reading the request line, the headers, or the body.

## Practical Takeaways

- HTTP request parsing is layered on top of TCP stream reading.
- The line ending for HTTP is `\r\n`.
- The blank line `\r\n` after headers marks the end of the header section.
- The parser has to preserve unread bytes between calls.
- The request should be understood as structured protocol data rather than as one unorganized string.
