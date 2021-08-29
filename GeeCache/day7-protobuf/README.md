# Protocol Buffer

> Protocol Buffer like XML, but smaller, faster, and simpler.

## How to use it?

*Step 1*

Download [binary](https://github.com/protocolbuffers/protobuf/releases)

*Step 2*

`geecachepb.proto`

```
syntax = "proto3";

option go_package = "./geecachepb";
package geecachepb;

message Request{
  string group = 1;
  string key = 2;
}

message Response {
  bytes value = 1;
}

service GroupCache{
  rpc Get(Request) returns (Response);
}
```

*Step 3*

In `geecachepb.proto` folder.

```
protoc --go_out=.. *proto
```

## How to test?

Open 3 terminals.

*Terminal 1*

```
go run . -port=8001
```

*Terminal 2*

```
go run . -port=8002
```

*Terminal 3*

```
go run . -port=8003 -api=1
```

Run `TestPB` in `protobuf_test.go`
