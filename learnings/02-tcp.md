# Chapter 2: TCP

Course focus: understand reliable TCP streams and compare them with UDP.

## Lecture Goal

After learning that a stream is just bytes over time, the next step is to replace the file with the network. HTTP/1.1 normally runs over TCP, so the program needs to listen for connections and read bytes from a socket.

## Lecture Progression

1. Open a TCP port and wait for clients to connect.
2. Accept a client connection and treat it as a two-way byte stream.
3. Read from the connection using the same stream-reading idea from Chapter 1.
4. Close the connection when the conversation is finished.
5. Compare TCP with UDP to understand why HTTP/1.1 is built on top of TCP.

## Core Learning

TCP gives a reliable, ordered byte stream. Reliable means bytes are retransmitted if necessary. Ordered means the receiver observes bytes in the same order the sender wrote them. Stream means there are no message boundaries built into TCP.

TCP is connection-oriented. Before application data is exchanged, the client and server establish a connection. Once the connection exists, both sides can send and receive bytes. TCP tracks the state of that connection until it is closed.

TCP also uses sequence numbers and acknowledgements. At a high level, each side can tell which bytes have been received and which bytes may need to be sent again. If packets are lost or arrive out of order at the network level, TCP hides much of that complexity from the application and reconstructs the original ordered byte stream.

This is why TCP is useful for protocols like HTTP. HTTP messages need to arrive complete and in order. A request line such as `GET / HTTP/1.1` is only meaningful if the bytes arrive in the right order.

The tradeoff is that TCP does not preserve application message boundaries. It guarantees the order of bytes, not the shape of the messages inside those bytes.

That last point is essential. TCP can deliver these writes in many read patterns:

```text
sender writes: "GET / HTTP/1.1\r\n"
receiver read: "GET / HT"
receiver read: "TP/1.1\r\n"
```

The data is correct and ordered, but the application still has to parse it.

## TCP vs UDP

TCP and UDP are both transport protocols, but they make different promises.

TCP is like an ongoing phone call. A connection is established, both sides can speak, and the protocol works to keep the conversation ordered and reliable. If bytes are lost, TCP can retransmit them. If packets arrive out of order, TCP reorders them before the application reads them.

UDP is more like sending individual postcards. Each datagram is its own separate message. There is no long-lived connection, no built-in retry, and no guarantee that messages arrive in order or arrive at all. The upside is that UDP has less overhead and can be useful when speed matters more than perfect delivery.

| Concept            | TCP                                   | UDP                                                   |
| ---                | ---                                   | ---                                                   |
| Connection         | Connection-oriented                   | Connectionless                                        |
| Data model         | Continuous byte stream                | Individual datagrams                                  |
| Delivery           | Reliable delivery with retransmission | Best-effort delivery                                  |
| Ordering           | Preserves byte order                  | No ordering guarantee                                 |
| Message boundaries | Not preserved                         | Preserved per datagram                                |
| Common use cases   | HTTP/1.1, SSH, databases              | DNS, games, voice/video, custom low-latency protocols |

The most important contrast for this course is message boundaries. UDP preserves datagram boundaries, so one send maps to one received datagram if it arrives. TCP does not do that. TCP turns all writes into one ordered stream of bytes, so HTTP must define its own boundaries using `\r\n`, blank lines, `Content-Length`, and later chunked encoding.

## Practical Takeaways

- TCP connections are stateful conversations between two endpoints.
- UDP sends discrete datagrams, while TCP exposes a continuous stream.
- HTTP parsing belongs above TCP because TCP only transports bytes.
- TCP handles reliable delivery and ordering, but it does not understand HTTP message boundaries.
- Protocols such as HTTP define their own structure on top of the TCP stream.
- Reliability and ordering are transport guarantees; request lines, headers, and bodies are HTTP concepts layered above those guarantees.
