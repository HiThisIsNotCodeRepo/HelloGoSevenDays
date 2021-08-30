# Codec

## A typical RPC call

```
err = client.Call("Arith.Multiply", args, &reply)
```

*Request from client*

1. Service name: `Arith`
2. Method name: `Multiply`
3. Arguments: `args`

*Response from server*

1. Error: `error`
2. Reply: `reply`

To simplify model we use the `body` to store the `args` from request and the `reply` from response and `error`, `Arith`
,`Multiply` we use `Header` to store.

```
type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}
```

## Encode and decode body

`Codec` is the interface describe how we encode/decode body

```
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}
```

## Customize protocol

`Option` is like `Header` in HTML it describes how message encode and decode.

```
type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

```

So when the server receives the request it will decode `Option` then retrieve the information of `CodecType`, then
decode message.

## Ensure the server is on successfully then start client

```
func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	addr := make(chan string)
	go startServer(addr)
	conn, _ := net.Dial("tcp", <-addr)
	...
}
```