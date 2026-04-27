# Chapter 5: HTTP Headers

Course focus: understand how HTTP headers are structured, parsed, and handled.

## Lecture Goal

Headers are the metadata section of an HTTP message. They come after the request line and before the body. This chapter turns raw header lines into a structured collection that the request parser can use.

## Header Format

```text
field-name ":" optional-whitespace field-value CRLF
```

Example:

```text
Host: localhost:42069\r\n
User-Agent: curl/7.81.0\r\n
Accept: */*\r\n
\r\n
```

The final empty line marks the end of the header section.

## Lecture Progression

1. `L1 - Headers`: parse basic `Name: value` lines and detect the blank line that ends headers.
2. `L2 - Constraints`: enforce valid field-name rules and reject malformed spacing before the colon.
3. `L3 - Multiple Values`: handle repeated field names by combining values with commas.
4. `L4 - Add to Parse`: integrate header parsing into the request parser state machine.
5. `L5 - Live Headers`: print parsed headers from a real TCP request.

## Core Learning

HTTP headers are case-insensitive by name. `Host`, `host`, and `HOST` refer to the same field. The parser normalizes names to lowercase so callers can use consistent lookup behavior.

Header values are more permissive than header names. Values can contain surrounding whitespace that should be trimmed, but names must be valid tokens. A malformed field name means the HTTP message is malformed.

## Practical Takeaways

- Split each header line at the first colon, not every colon.
- Reject a header with whitespace before the colon.
- Trim surrounding whitespace from the value.
- Normalize field names to lowercase for storage.
- Distinguish missing headers from present headers with empty values.
- Treat the empty line as the transition out of header parsing.
- Headers determine how later parts of the message should be interpreted.
