# Single Flight

> Cache avalanche: The large amount of cache expire or remove at the same time, cause database communication request explode.

> Cache breakdown: The key expires followed by large amounts of request to retrieve data. The cache needs to communicate with database but exceed database bottleneck hence cause DB down.

> Cache penetration: The key isn't cached but large amounts of request to retrieve data. The cache needs to communicate with database but exceed database bottleneck hence cause DB down.

## How to implement?

```
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}

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

Run `TestSingleFlight` in `single_flight_test.go`

`[SlowDB] search key Tom` should be displayed only once.