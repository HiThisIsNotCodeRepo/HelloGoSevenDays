package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	router *router
	*RouterGroup
	groups        []*RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
func (r *RouterGroup) Group(prefix string) *RouterGroup {
	engine := r.engine
	newGroup := &RouterGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (r *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := r.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	r.engine.router.addRoute(method, pattern, handler)
}

func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}

func (r *RouterGroup) Use(middlewares ...HandlerFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
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

func (r *RouterGroup) Static(relativePath string, root string) {
	handler := r.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	r.GET(urlPattern, handler)
}
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = e
	e.router.handle(c)
}

func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
