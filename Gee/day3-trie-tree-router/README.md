# Trie tree

## The limitation of the map

Map is suitable for static route path `/hello/world`, it doesn't support dynamic route path `/hello/:name`. To do that
we can use `Trie tree` data structure.

```
type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

```

However, the difference adding order will result in different behavior.

*Success*

```

func TestRegisterStaticThenDynamic(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/hello/world", nil)
	r.addRoute("GET", "/:foo/:bar", nil)
	n, ps := r.getRoute("GET", "/hello/b")
	require.NotNil(t, n)
	require.Equal(t, ps["foo"], "hello")
	require.Equal(t, ps["bar"], "b")
}
```

*Fail*

```
func TestRegisterDynamicThenStatic(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/:foo/:bar", nil)
	r.addRoute("GET", "/hello/world", nil)
	n, ps := r.getRoute("GET", "/hello/b")
	require.NotNil(t, n)
	require.Equal(t, ps["foo"], "hello")
	require.Equal(t, ps["bar"], "b")
}
```