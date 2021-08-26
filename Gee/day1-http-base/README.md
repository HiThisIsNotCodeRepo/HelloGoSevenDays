# HTTP base

## Standard lib not support

1. Dynamic router: for instance `hello/:name`, `hello/*`
2. Authorisation
3. Simple way to handle HTML

## Core value of a web framework

1. Routing, specifically map request to function and support dynamic router.
2. Templates, using build in template engine to render.
3. Utilities, provide mechanism to deal with cookies and headers.
4. Plugin, with option to install globally or locally.

## Write a web server using system build in handler

```
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// nil means using standard lib instance
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// Router 1
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// Router 2
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

}

```

## Using customized handler

```
// Engine is the handler for all request
type Engine struct {
}

// Engine now implement Handler interface
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q \n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatalln(http.ListenAndServe(":9999", engine))
}

```

## Using package with relative import

### How to achieve relative import

Directory structure

```
gee/
  |--gee.go
  |--go.mod
main.go
go.mod
```

#### gee/go.mod

```
module gee

go 1.16

```

#### go.mod

```
module example

go 1.16

require gee v0.0.0

replace gee => ./gee

```

## Gee package content

```
// HandlerFunc defines the request handler used by gee or can just use http.Handler
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

// New return Engine instance
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// Helper function add handler to router map
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// Get is for user use, register Get method
func (e *Engine) Get(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST is for user use, register POST method
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// Run encapsulate ListenAndServe
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP implement http.handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

```

