# Gee

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