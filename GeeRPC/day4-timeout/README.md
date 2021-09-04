# Timeout

## Why need to handle timeout

When computation requires large amount of resource like http call ,database connection or reflection and in such
condition the computation may not 100% successful. We need to impose timeout control to prevent it from consuming
resources endlessly.

## Syntax to handle timeout

1. `ctx, _ := context.WithTimeout(context.Background(), time.Second)`.
2. `time.After()` and `select + chan`.

## Client side

*Connection timeout*

Add another function to handle timeout inside `Dial`.

```
func dialTimeout(f newClientFunc, network, address string, opts ...*Option) (client *Client, err error) {
	...
	// No timeout setting
	if opt.ConnectTimeout == 0 {
		result := <-ch
		return result.client, result.err
	}
	// Got timeout setting
	select {
	case <-time.After(opt.ConnectTimeout):
		return nil, fmt.Errorf("rpc client: connect timeout: expecte within %s", opt.ConnectTimeout)
	case result := <-ch:
		return result.client, result.err
	}
}
```

*Call timeout*

```
func (c *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	...
	select {
	case <-ctx.Done():
		c.removeCall(call.Seq)
		return errors.New("rpc client: call failed: " + ctx.Err().Error())
	case call := <-call.Done:
		return call.Error
	}
}
```

To set timeout limit just:

```
...
ctx, _ := context.WithTimeout(context.Background(), time.Second)
var reply int
err := client.Call(ctx, "Foo.Sum", &Args{1, 2}, &reply)
...
```

## Server side

*Handle timeout*

```
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup, timeout time.Duration) {
	...
	select {
	case <-time.After(timeout):
		req.h.Error = fmt.Sprintf("rpc server: request handle timeout: expect within %s", timeout)
		server.sendResponse(cc, req.h, invalidRequest, sending)
	case <-called:
		<-sent
	}
}
```