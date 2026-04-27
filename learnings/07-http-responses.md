# Chapter 7: HTTP Responses

Course focus: construct and send HTTP server responses.

Known lesson page checked: [Content Length](https://www.boot.dev/lessons/5639db38-6e0d-41f0-9275-9bf3b97693e4)

## Lecture Goal

Requests are what the client sends. Responses are what the server sends back. This chapter begins the server side of the protocol: accept a TCP connection, write bytes that follow HTTP response syntax, and make the client understand where the response body ends.

## Response Structure

An HTTP response has the same broad shape as a request: start line, headers, blank line, optional body. The start line is different because the server is reporting an outcome rather than asking for a resource.

```text
HTTP/1.1 200 OK\r\n
Content-Type: text/plain\r\n
Content-Length: 13\r\n
\r\n
Hello World!
```

The fields mean:

- `HTTP/1.1`: protocol version.
- `200`: status code.
- `OK`: human-readable reason phrase.
- `Content-Type`: how to interpret the body.
- `Content-Length`: how many body bytes follow.

## Lecture Progression So Far

1. Create a dedicated HTTP server entry point.
2. Listen on a TCP port and accept client connections.
3. Write a hard-coded HTTP response directly to the connection.
4. Close the connection after the response is written.
5. Use the `Content Length` lesson to observe why response body framing matters.

## Core Learning

HTTP responses are just bytes over TCP, but the bytes must follow the protocol. A browser or `curl` does not know that `Hello World!` is complete unless the response provides framing information.

The course lesson highlights the RFC rule behind `Content-Length`: for messages with content, the field provides the framing information needed to determine where the data and message end.

If `Content-Length` is missing, the client can fall back to connection close. That works for simple examples, but it is weaker because the end of the message is no longer described by the message itself.

## Practical Takeaways

- A response status line is not the same as a request line.
- `Content-Type` describes the body format.
- `Content-Length` describes the body size.
- A valid-looking response can still have poor framing if the body length is omitted.
- Closing the TCP connection can signal the end of a response, but explicit length is clearer and expected for normal fixed-size bodies.
- Response framing is as important as request framing because clients also need to know when a message is complete.
