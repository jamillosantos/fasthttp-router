package fasthttp_router

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"bytes"
)

type Middleware func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

func Middlewares(handler fasthttp.RequestHandler, middlewares ...Middleware) fasthttp.RequestHandler {
	result := handler
	for i := 0; i < len(middlewares); i++ {
		result = middlewares[i](result)
	}
	return result
}

type Router struct {
	children map[string]*node
	NotFound fasthttp.RequestHandler
}

func NewRouter() *Router {
	return &Router{
		children: make(map[string]*node),
	}
}

func (router *Router) handle(method, path string, handler fasthttp.RequestHandler) {
	root, ok := router.children[method]
	if !ok {
		root = newNode()
		router.children[method] = root
	}

	if path[0] == '/' {
		path = path[1:]
	}
	root.Add(path, handler, make([]string, 0))
}

func (router *Router) DELETE(path string, handler fasthttp.RequestHandler) {
	router.handle("DELETE", path, handler)
}

func (router *Router) GET(path string, handler fasthttp.RequestHandler) {
	router.handle("GET", path, handler)
}

func (router *Router) POST(path string, handler fasthttp.RequestHandler) {
	router.handle("POST", path, handler)
}

func (router *Router) PUT(path string, handler fasthttp.RequestHandler) {
	router.handle("PUT", path, handler)
}

func (router *Router) HEAD(path string, handler fasthttp.RequestHandler) {
	router.handle("HEAD", path, handler)
}

func (router *Router) OPTIONS(path string, handler fasthttp.RequestHandler) {
	router.handle("OPTIONS", path, handler)
}

func (router *Router) PATCH(path string, handler fasthttp.RequestHandler) {
	router.handle("PATCH", path, handler)
}

func (router *Router) Group(path string, middlewares ... Middleware) Routable {
	return &routerGroup{
		prefix:      path,
		router:      router,
		middlewares: middlewares,
	}
}

var (
	routerHandlerSep = []byte{'/'}
)

func (router *Router) Handler(ctx *fasthttp.RequestCtx) {
	method := string(ctx.Method())
	node, ok := router.children[method]
	if ok {
		var path [][]byte
		if ctx.Request.URI().Path()[0] == '/' {
			path = bytes.Split(ctx.Request.URI().Path()[1:], routerHandlerSep)
		} else {
			path = bytes.Split(ctx.Request.URI().Path(), routerHandlerSep)
		}
		found, node, values := node.Matches(path, make([][]byte, 0))
		if found {
			for i, v := range values {
				ctx.SetUserValue(node.names[i], string(v))
			}
			node.handler(ctx)
			return
		}
	}
	if router.NotFound != nil {
		router.NotFound(ctx)
	}
}

type routerGroup struct {
	prefix      string
	router      Routable
	middlewares []Middleware
}

func (group *routerGroup) DELETE(path string, handler fasthttp.RequestHandler) {
	group.router.DELETE(fmt.Sprintf("%s%s", group.prefix, path), Middlewares(handler, group.middlewares...))
}

func (group *routerGroup) GET(path string, handler fasthttp.RequestHandler) {
	group.router.GET(fmt.Sprintf("%s%s", group.prefix, path), Middlewares(handler, group.middlewares...))
}

func (group *routerGroup) POST(path string, handler fasthttp.RequestHandler) {
	group.router.POST(fmt.Sprintf("%s%s", group.prefix, path), Middlewares(handler, group.middlewares...))
}

func (group *routerGroup) PUT(path string, handler fasthttp.RequestHandler) {
	group.router.PUT(fmt.Sprintf("%s%s", group.prefix, path), Middlewares(handler, group.middlewares...))
}

func (group *routerGroup) HEAD(path string, handler fasthttp.RequestHandler) {
	group.router.HEAD(fmt.Sprintf("%s%s", group.prefix, path), Middlewares(handler, group.middlewares...))
}

func (group *routerGroup) OPTIONS(path string, handler fasthttp.RequestHandler) {
	group.router.OPTIONS(fmt.Sprintf("%s%s", group.prefix, path), Middlewares(handler, group.middlewares...))
}

func (group *routerGroup) PATCH(path string, handler fasthttp.RequestHandler) {
	group.router.PATCH(fmt.Sprintf("%s%s", group.prefix, path), Middlewares(handler, group.middlewares...))
}

func (group *routerGroup) Group(path string, middlewares ... Middleware) Routable {
	return &routerGroup{
		prefix:      path,
		router:      group,
		middlewares: middlewares,
	}
}