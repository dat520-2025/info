# Remote Procedure Call (RPC)

Old technology, but recent implementations have improved.

- C RPC
- Go `net/rpc` -- only supports Go clients and servers
  - `gob` encoding
- CORBA -- Common Object Request Broker Architecture
  - Inspired gRPC
  - Interface Description Language (IDL)
- Java RMI -- Remote Method Invocation
- SOAP -- Simple Object Access Protocol
  - XML encoding

## gRPC

- Google RPC -- gRPC Remote Procedure Call
- Protocol Buffers -- `protobuf`

## How it works

### Client

- Client program invokes the local gRPC stub
- The stub takes care of marshaling (also called serialization) the arguments of a remote call (a request message) into a byte stream
- The stub sends the byte stream to the server
  - The method name to call on the server
  - The input arguments
- The stub also takes care of sending the marshaled msg over net.

### Server side

- The gRPC server receives the request msg from a client and unmarshals it:
  - method name
  - the input arguments
- The gRPC server then "finds" the method on the server
  - Often done using code generation or reflection
  - Given a method name as a string; use reflection to find actual method to invoke.
- The gRPC server stub then invokes the method on the server side
  - Gets the results of the method call
  - Marshals the result of the method call into a byte array
  - Send the byte array over the network.
- The client gRPC stub receives the result message and delivers the result to the client program.
