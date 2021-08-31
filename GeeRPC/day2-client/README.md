# Codec

## What kind of function can be called remotely?

1. Method type is exported.
2. Method is exported.
3. Method has two arguments, both exported.
4. Method 2nd argument is a pointer.
5. Method return error.

```
func (t *T) MethodName(argType T1, replyType *T2) error
```

## Features of client

1. Send request.

`Call` will block until `call` receives result.

```
func (c *Client) Call(serviceMethod string, args, reply interface{}) error {
	call := <-c.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}
```

2. Receive response.

Once client is on the `receive()` goroutine init.

```
func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1,
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}
```

It reads response and extract call information then call its `done()`.

```
func (c *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = c.cc.ReadHeader(&h); err != nil {
			break
		}
		call := c.removeCall(h.Seq)
		switch {
		case call == nil:
			err = c.cc.ReadBody(nil)
		case h.Error != "":
			call.Error = fmt.Errorf(h.Error)
			err = c.cc.ReadBody(nil)
			call.done()
		default:
			err = c.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body" + err.Error())
			}
			call.done()
		}
	}
	c.terminateCalls(err)
}
```

```
func (c *Call) done() {
	c.Done <- c
}
```