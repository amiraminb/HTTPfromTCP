# Chapter 1: HTTP Streams

Course focus: read and process a stream of bytes as they arrive.

## Lecture Goal

HTTP starts as bytes. Before thinking about methods, headers, status codes, or bodies, the first lesson is that programs do not receive tidy protocol messages. They receive chunks from a stream, and those chunks must be assembled into meaningful units.

## Lecture Progression

1. Start with a plain file as the simplest stream source.
2. Read a fixed-size byte buffer repeatedly instead of reading the whole file at once.
3. Observe that a read boundary is not the same as a line boundary.
4. Keep unfinished text between reads until a delimiter is found.
5. Separate line extraction from stream reading so complete lines can be emitted as they become available.

## Core Learning

A stream is continuous data, not a sequence of already-parsed messages. If a program reads 8 bytes at a time, a line can be split in the middle, or one read can contain the end of one line and the start of another. The parser must carry state between reads.

```text
read 1: Do you h
read 2: ave what
read 3: it takes\nAre
```

The newline is the application-level boundary. The stream does not care about that boundary. The application must define it and detect it.

## Practical Takeaways

- A stream must be consumed progressively rather than assumed to arrive all at once.
- A read operation may return less data than expected, so parsers must be prepared for partial input.
- A delimiter such as `\n` marks a logical boundary only after the application looks for it.
- Partial data must be saved until the delimiter arrives.
- End-of-stream is a normal condition, not necessarily an error.
- Stream processing is the foundation for HTTP parsing because HTTP messages also arrive as bytes over time.
