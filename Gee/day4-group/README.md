# Group

## What does group do

To enable group control we need `RouterGroup` struct

```
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}
```

In this struct `engine` is referenced the only engine instance. And we want to use engine to create group that means we
need engine to expose `RouterGroup` API, this is achieved by embedding

```
type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup
}
```