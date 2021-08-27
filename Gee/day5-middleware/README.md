# Middleware

## How it works

In this example the middles and handlers are stored in the context.

```
type Context struct {
        ...
	handlers   []HandlerFunc
	index      int
}
```

They will be added when router been triggered

```
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
```

Above snippet shows that in `c.handlers` the actual implementation of http request is put at the end of array.

```
func (c *Context) Next() {
	c.index++
	c.handlers[c.index](c)
}
```

`Next()` will ensure the middleware will run in sequence.

## What criteria make a framework successful

*It's my guessing*

1. Algorithm time complexity, smallest is the best.

```
O(1) < O(logn) < O(n) < O(nlogn) < O(n^2) < O(n^3) < O(2^n) < O(n!) < O(n^n)
```

2. Lines of code.
3. Features.