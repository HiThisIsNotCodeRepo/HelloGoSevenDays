package gee

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:age", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/:address/:street", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	require.Equal(t, parsePattern("/p/:name"), []string{"p", ":name"})
	require.Equal(t, parsePattern("/p/*"), []string{"p", "*"})
	require.Equal(t, parsePattern("/x/*/*/*"), []string{"x", "*"})
	require.Equal(t, parsePattern("/p/*name/*"), []string{"p", "*name"})
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")
	require.NotNil(t, n)
	require.Equal(t, n.pattern, "/hello/:name")
	require.Equal(t, ps["name"], "geektutu")
	//require.Equal(t, ps["age"], "geektutu")
	n, ps = r.getRoute("GET", "/hello/test_address/test_street")
	require.NotNil(t, n)
	t.Log(ps)
}

func TestInsert(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/hello/:address/:street", nil)
	n, ps := r.getRoute("GET", "/hello/b/c")
	require.NotNil(t, n)
	require.Equal(t, ps["address"], "b")
	require.Equal(t, ps["street"], "c")
}
func TestRegisterStaticThenDynamic(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/hello/world", nil)
	r.addRoute("GET", "/:foo/:bar", nil)
	n, ps := r.getRoute("GET", "/hello/b")
	require.NotNil(t, n)
	require.Equal(t, ps["foo"], "hello")
	require.Equal(t, ps["bar"], "b")
}
func TestRegisterDynamicThenStatic(t *testing.T) {
	r := newRouter()
	r.addRoute("GET", "/:foo/:bar", nil)
	r.addRoute("GET", "/hello/world", nil)
	n, ps := r.getRoute("GET", "/hello/b")
	require.NotNil(t, n)
	require.Equal(t, ps["foo"], "hello")
	require.Equal(t, ps["bar"], "b")
}
