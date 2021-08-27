# Static file server

## How it works

```
func main() {
	r := gee.New()
	r.Static("/assets","./static_asset")
	r.Run(":9999")
}

func (r *RouterGroup) Static(relativePath string, root string) {
	handler := r.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	r.GET(urlPattern, handler)
}

func (r *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(r.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}
```

Imagine we want to use url `http://localhost:9999/assets/foo.js` to retrieve `foo.js`

1. In `main()` call `Static` with relative url path `/assets` and the file root path `./static_asset`, it means the url
   prefix with `/assets` will be handled by static file server and the file we need is stored in the `./static_asset`
2. To create file server handler we need to cast type from `string` to `http.Dir` because it has
   implemented `http.FileSystem` interface. This step can be considered as to load files from `./static_asset`.

```
type FileSystem interface {
	Open(name string) (File, error)
}
```

3. To map file name correctly on file system we need to strip the absolute path for example convert `assets/foo.js`
   to `foo.js`, therefore we call `http.StripPrefix(absolutePath, http.FileServer(fs))` otherwise `fs.Open()` may fail
   if file name is not correct. This step can be considered as make the file name correct.
4. Define the Handler function , inside it will retrieve file name and open it then return.
5. To retrieve the file name we have defined the url pattern, so any url match `assets/...` will be added `/*filepath`
   behind so that the `{file_name}`inside `assets/{file_name}` can be read. 