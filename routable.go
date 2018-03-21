package fasthttp_router

import "github.com/valyala/fasthttp"

type Routable interface {
	DELETE(path string, handler fasthttp.RequestHandler)
	GET(path string, handler fasthttp.RequestHandler)
	HEAD(path string, handler fasthttp.RequestHandler)
	OPTIONS(path string, handler fasthttp.RequestHandler)
	PATCH(path string, handler fasthttp.RequestHandler)
	POST(path string, handler fasthttp.RequestHandler)
	PUT(path string, handler fasthttp.RequestHandler)

	Group(path string, middlewares ...Middleware) Routable
}
