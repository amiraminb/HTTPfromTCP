# IP and Network Layers

## Lecture Goal

TCP is not the bottom of the stack. TCP is built on IP, and IP is built on lower network layers that eventually become electrical signals, light pulses, or radio waves. The goal of this lecture is to understand how data moves downward through layers before it crosses the network, and then back upward when it reaches the destination.

## The Layered Model

Network protocols are layered so each layer can focus on one responsibility.

```text
HTTP
TCP or UDP
IP
Ethernet / Wi-Fi / cellular
Physical signals
```

Each layer uses the layer below it. HTTP does not need to know how radio waves work. TCP does not need to know how an Ethernet cable sends bits. IP does not need to know what an HTTP header means.

## What IP Does

IP stands for Internet Protocol. Its main job is to move packets from one machine to another using IP addresses.

IP provides:

- Source and destination IP addresses.
- A packet format that routers understand.
- Routing across networks.
- A hop limit so packets do not circle forever.
- Best-effort delivery.

Best-effort is important. IP tries to deliver packets, but it does not guarantee delivery. A packet can be lost, duplicated, delayed, or arrive out of order. IP also does not create a connection and does not know about ports. Ports belong to TCP and UDP.

## Packets and Routing

An IP packet contains a header and a payload.

```text
IP header
TCP segment or UDP datagram
```

The IP header includes information such as the source address, destination address, and hop limit. The payload is usually a TCP segment or UDP datagram.

Routers inspect the destination IP address and decide where to send the packet next. A packet may pass through many routers before reaching the destination network. Each router only needs to know the next hop, not the full application-level meaning of the packet.

## What IP Does Not Do

IP does not provide the guarantees that TCP provides.

IP does not guarantee:

- Delivery.
- Ordering.
- Deduplication.
- Retransmission.
- A continuous byte stream.
- Application message boundaries.

This is why TCP exists above IP. TCP uses IP packets, then adds reliability, ordering, retransmission, and stream behavior.

## What IP Is Built On

IP is built on a link layer. The link layer moves data across one local network segment.

Common link layers include:

- Ethernet for wired local networks.
- Wi-Fi for wireless local networks.
- Cellular links for mobile networks.

The link layer usually works with frames, not IP packets directly. An IP packet is placed inside a link-layer frame.

```text
Ethernet or Wi-Fi frame
IP packet
TCP segment
HTTP data
```

On a local network, devices often use hardware addresses such as MAC addresses. A router may receive a frame on one interface, extract the IP packet, decide the next hop, and then place that same IP packet into a new frame for the next network segment.

## What The Link Layer Is Built On

The link layer is built on the physical layer. This is the lowest layer: actual signals moving through physical media.

Examples include:

- Electrical signals over copper cables.
- Light pulses through fiber optic cables.
- Radio waves for Wi-Fi.
- Radio signals for cellular networks.

At this level, the concern is how bits are physically represented and transmitted. Higher layers see bytes and packets, but the physical layer ultimately sends changing voltages, light, or radio energy.

## Encapsulation

Each layer wraps the data from the layer above it. This is called encapsulation.

When sending an HTTP request, the data is wrapped like this:

```text
HTTP request
inside a TCP segment
inside an IP packet
inside an Ethernet or Wi-Fi frame
as physical signals
```

When receiving, the process is reversed:

```text
physical signals
become an Ethernet or Wi-Fi frame
which contains an IP packet
which contains a TCP segment
which contains HTTP data
```

Each layer removes its own wrapper and passes the remaining payload upward.

## Why This Matters For HTTP

HTTP depends on many lower layers, but each layer has a separate job.

- HTTP defines request and response meaning.
- TCP defines reliable ordered byte streams.
- IP defines addressing and routing between machines.
- Ethernet, Wi-Fi, and cellular define local delivery across one network segment.
- The physical layer defines how bits move through the real world.

This separation is why HTTP can work over Ethernet, Wi-Fi, or cellular without changing the HTTP protocol. HTTP sits at the top and trusts the lower layers to move bytes.

## Practical Takeaways

- TCP is built on IP.
- IP is built on link-layer technologies such as Ethernet and Wi-Fi.
- Link-layer technologies are built on physical signals.
- IP is connectionless and best-effort.
- TCP adds reliability and ordering above IP.
- HTTP adds application meaning above TCP.
