# Protocol Buffers
> Protocol Buffers are a language-neutral, platform-neutral, extensible way of
  serialising structured data for use in communications protocols, data storage,
  and more, originally designed at Google.

## Hello Protocol Buffers 👋
You define a **data structure** (schema) and then compile it to source code
(using the Protocol Buffer compiler) to **write/read your structured data**
to/from **data streams**, using a variety of languages.

So, Protocol Buffers can be described as a mechanism for enforcing schematic
shape for language-agnostic communication with data streams (or, sequences of
data packets used to transmit or receive information).

Protocol Buffers can be seen as an alternative to XML-based data serialisation.
However, XML is notoriously space intensive, and encoding/decoding it can impose
a huge performance penalty on applications. Also, navigating an XML DOM tree is
considerably more complicated than navigating simple fields in a class.

### Define a data structure in a `.proto` description
```
message Person {
  required string name = 1;
  required int32 id = 2;
  optional string email = 3;
}
```

The Protocol Buffer compiler creates a class that implements automatic
encoding and parsing of the protocol buffer data with an efficient binary format.

### Write structured data to data stream
```
Person john = Person.newBuilder()
    .setId(1234)
    .setName("John Doe")
    .setEmail("jdoe@example.com")
    .build();
output = new FileOutputStream(args[0]);
john.writeTo(output);
```

The generated class provides getters and setters for the fields that make up a
protocol buffer and takes care of the details of reading and writing the
protocol buffer as a unit.

### Read structured data from data stream
```
Person john;
fstream input(argv[1], ios::in | ios::binary);
john.ParseFromIstream(&input);
id = john.id();
name = john.name();
email = john.email();
```

The protocol buffer format supports the idea of extending the format over time
in such a way that the code can still read data encoded with the old format.

## Tutorials
- [Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [JS](https://github.com/dcodeIO/ProtoBuf.js/)

## Reference
- [`.proto` Language Guide](https://developers.google.com/protocol-buffers/docs/proto3)

## Setup
- Install the Protocol Buffer compiler
```
brew install protobuf
```

- Install the Go Protocol Buffer plugin
```
go get github.com/golang/protobuf/protoc-gen-go
```

- Run the compiler
```
protoc -I=. --go_out=. ./addressbook.proto
```

The `--go_out` option is used to generate [Go classes](https://developers.google.com/protocol-buffers/docs/reference/go-generated).
Similar options are provided for other supported languages, like `--js_out` for
[JavaScript output](https://developers.google.com/protocol-buffers/docs/reference/javascript-generated).
