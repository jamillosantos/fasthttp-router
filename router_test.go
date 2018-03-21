package fasthttp_router

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
	"fmt"
)

func createRequestCtxFromPath(method, path string) *fasthttp.RequestCtx {
	result := &fasthttp.RequestCtx{}
	result.Request.Header.SetMethod(method)
	result.Request.URI().SetPath(path)
	return result
}

var _ = Describe("Router", func() {

	var emptyHandler fasthttp.RequestHandler = func(ctx *fasthttp.RequestCtx) {
	}

	Describe("Parse", func() {

		It("should parse a GET", func() {
			router := NewRouter()
			router.GET("/route", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(HaveKey("route"))
			Expect(router.children["GET"].wildcard).To(BeNil())
		})

		It("should parse a POST", func() {
			router := NewRouter()
			router.POST("/route", emptyHandler)

			Expect(router.children).To(HaveKey("POST"))
			Expect(router.children["POST"].children).To(HaveKey("route"))
			Expect(router.children["POST"].wildcard).To(BeNil())
		})

		It("should parse a PUT", func() {
			router := NewRouter()
			router.PUT("/route", emptyHandler)

			Expect(router.children).To(HaveKey("PUT"))
			Expect(router.children["PUT"].children).To(HaveKey("route"))
			Expect(router.children["PUT"].wildcard).To(BeNil())
		})

		It("should parse a DELETE", func() {
			router := NewRouter()
			router.DELETE("/route", emptyHandler)

			Expect(router.children).To(HaveKey("DELETE"))
			Expect(router.children["DELETE"].children).To(HaveKey("route"))
			Expect(router.children["DELETE"].wildcard).To(BeNil())
		})

		It("should parse a HEAD", func() {
			router := NewRouter()
			router.HEAD("/route", emptyHandler)

			Expect(router.children).To(HaveKey("HEAD"))
			Expect(router.children["HEAD"].children).To(HaveKey("route"))
			Expect(router.children["HEAD"].wildcard).To(BeNil())
		})

		It("should parse a OPTIONS", func() {
			router := NewRouter()
			router.OPTIONS("/route", emptyHandler)

			Expect(router.children).To(HaveKey("OPTIONS"))
			Expect(router.children["OPTIONS"].children).To(HaveKey("route"))
			Expect(router.children["OPTIONS"].wildcard).To(BeNil())
		})

		It("should parse a PATCH", func() {
			router := NewRouter()
			router.PATCH("/route", emptyHandler)

			Expect(router.children).To(HaveKey("PATCH"))
			Expect(router.children["PATCH"].children).To(HaveKey("route"))
			Expect(router.children["PATCH"].wildcard).To(BeNil())
		})

		It("should parse a POST", func() {
			router := NewRouter()
			router.POST("/route", emptyHandler)

			Expect(router.children).To(HaveKey("POST"))
			Expect(router.children["POST"].children).To(HaveKey("route"))
			Expect(router.children["POST"].wildcard).To(BeNil())
		})

		It("should parse a complete static route", func() {
			router := NewRouter()
			router.GET("/this/should/be/static", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(HaveKey("this"))
			Expect(router.children["GET"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children).To(HaveKey("should"))
			Expect(router.children["GET"].children["this"].children["should"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children).To(HaveKey("be"))
			Expect(router.children["GET"].children["this"].children["should"].children["be"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children["be"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children["be"].children).To(HaveKey("static"))
			Expect(fmt.Sprintf("%p", router.children["GET"].children["this"].children["should"].children["be"].children["static"].handler)).To(Equal(fmt.Sprintf("%p", emptyHandler)))
		})

		It("should parse multiple static routes related", func() {
			router := NewRouter()
			router.GET("/this/should/be/static", emptyHandler)
			router.GET("/this/should2/be/static", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(HaveKey("this"))
			Expect(router.children["GET"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].handler).To(BeNil())

			Expect(router.children["GET"].children["this"].children).To(HaveKey("should"))
			Expect(router.children["GET"].children["this"].children["should"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children).To(HaveKey("be"))
			Expect(router.children["GET"].children["this"].children["should"].children["be"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children["be"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children["be"].children).To(HaveKey("static"))
			Expect(fmt.Sprintf("%p", router.children["GET"].children["this"].children["should"].children["be"].children["static"].handler)).To(Equal(fmt.Sprintf("%p", emptyHandler)))

			Expect(router.children["GET"].children["this"].children).To(HaveKey("should2"))
			Expect(router.children["GET"].children["this"].children["should2"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].children).To(HaveKey("be"))
			Expect(router.children["GET"].children["this"].children["should2"].children["be"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].children["be"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].children["be"].children).To(HaveKey("static"))
			Expect(fmt.Sprintf("%p", router.children["GET"].children["this"].children["should2"].children["be"].children["static"].handler)).To(Equal(fmt.Sprintf("%p", emptyHandler)))
		})

		It("should parse a complete a route starting static and ending with a wildcard", func() {
			router := NewRouter()
			router.GET("/static/:wildcard", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].wildcard).To(BeNil())
			Expect(router.children["GET"].names).To(BeEmpty())
			Expect(router.children["GET"].children).To(HaveKey("static"))
			Expect(router.children["GET"].children["static"].children).To(BeEmpty())
			Expect(router.children["GET"].children["static"].handler).To(BeNil())
			Expect(router.children["GET"].children["static"].wildcard).NotTo(BeNil())
			Expect(router.children["GET"].children["static"].wildcard.handler).NotTo(BeNil())
			Expect(router.children["GET"].children["static"].wildcard.children).To(BeEmpty())
			Expect(router.children["GET"].children["static"].wildcard.names).To(Equal([]string{"wildcard"}))
		})

		It("should parse multiple static routes related and not", func() {
			router := NewRouter()
			router.GET("/this/should/be/static", emptyHandler)
			router.GET("/this/should2/be/static", emptyHandler)
			router.GET("/this2/should/be/static", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(HaveKey("this"))
			Expect(router.children["GET"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children).To(HaveKey("should"))
			Expect(router.children["GET"].children["this"].children["should"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children).To(HaveKey("be"))
			Expect(router.children["GET"].children["this"].children["should"].children["be"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children["be"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should"].children["be"].children).To(HaveKey("static"))
			Expect(fmt.Sprintf("%p", router.children["GET"].children["this"].children["should"].children["be"].children["static"].handler)).To(Equal(fmt.Sprintf("%p", emptyHandler)))

			Expect(router.children["GET"].children["this"].children).To(HaveKey("should2"))
			Expect(router.children["GET"].children["this"].children["should2"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].children).To(HaveKey("be"))
			Expect(router.children["GET"].children["this"].children["should2"].children["be"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].children["be"].handler).To(BeNil())
			Expect(router.children["GET"].children["this"].children["should2"].children["be"].children).To(HaveKey("static"))
			Expect(fmt.Sprintf("%p", router.children["GET"].children["this"].children["should2"].children["be"].children["static"].handler)).To(Equal(fmt.Sprintf("%p", emptyHandler)))

			Expect(router.children["GET"].children).To(HaveKey("this2"))
			Expect(router.children["GET"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this2"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this2"].handler).To(BeNil())
			Expect(router.children["GET"].children["this2"].children).To(HaveKey("should"))
			Expect(router.children["GET"].children["this2"].children["should"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this2"].children["should"].handler).To(BeNil())
			Expect(router.children["GET"].children["this2"].children["should"].children).To(HaveKey("be"))
			Expect(router.children["GET"].children["this2"].children["should"].children["be"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["this2"].children["should"].children["be"].handler).To(BeNil())
			Expect(router.children["GET"].children["this2"].children["should"].children["be"].children).To(HaveKey("static"))
			Expect(fmt.Sprintf("%p", router.children["GET"].children["this2"].children["should"].children["be"].children["static"].handler)).To(Equal(fmt.Sprintf("%p", emptyHandler)))
		})

		It("should parse a complete route with wildcard", func() {
			router := NewRouter()
			router.GET("/:account/detail/another", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(BeEmpty())
			Expect(router.children["GET"].wildcard).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.children).To(HaveKey("detail"))
			Expect(router.children["GET"].wildcard.handler).To(BeNil())
			Expect(router.children["GET"].wildcard.wildcard).To(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].children).To(HaveKey("another"))
			Expect(router.children["GET"].wildcard.children["detail"].handler).To(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].wildcard).To(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].children["another"].wildcard).To(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].children["another"].handler).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].children["another"].names).To(Equal([]string{"account"}))
		})

		It("should parse a complete route with a sequence of wildcards", func() {
			router := NewRouter()
			router.GET("/:account/:transaction/:invoice", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(BeEmpty())
			Expect(router.children["GET"].wildcard).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.handler).To(BeNil())
			Expect(router.children["GET"].wildcard.names).To(BeEmpty())
			Expect(router.children["GET"].wildcard.wildcard).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.wildcard.handler).To(BeNil())
			Expect(router.children["GET"].wildcard.wildcard.names).To(BeEmpty())
			Expect(router.children["GET"].wildcard.wildcard.wildcard).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.wildcard.wildcard.wildcard).To(BeNil())
			Expect(router.children["GET"].wildcard.wildcard.wildcard.handler).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.wildcard.wildcard.names).To(Equal([]string{"account", "transaction", "invoice"}))
		})

		It("should parse multiple routes starting with wildcards", func() {
			router := NewRouter()
			router.GET("/:account/detail", emptyHandler)
			router.GET("/:account/history", emptyHandler)
			router.GET("/:transaction/invoice", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(BeEmpty())
			Expect(router.children["GET"].wildcard).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.children).To(HaveKey("detail"))
			Expect(router.children["GET"].wildcard.handler).To(BeNil())
			Expect(router.children["GET"].wildcard.wildcard).To(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].children).To(BeEmpty())
			Expect(router.children["GET"].wildcard.children["detail"].handler).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].wildcard).To(BeNil())
			Expect(router.children["GET"].wildcard.children["detail"].names).To(Equal([]string{"account"}))
			Expect(router.children["GET"].wildcard.children["history"].names).To(Equal([]string{"account"}))
			Expect(router.children["GET"].wildcard.children["invoice"].names).To(Equal([]string{"transaction"}))
		})

		It("should parse multiple mixed routes", func() {
			router := NewRouter()
			router.GET("/accounts/:account/detail", emptyHandler)
			router.GET("/accounts/:account/history", emptyHandler)
			router.GET("/:transaction/invoice", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].children).To(HaveLen(1))
			Expect(router.children["GET"].children).To(HaveKey("accounts"))
			Expect(router.children["GET"].children["accounts"].children).To(BeEmpty())
			Expect(router.children["GET"].children["accounts"].wildcard).NotTo(BeNil())
			Expect(router.children["GET"].children["accounts"].handler).To(BeNil())
			Expect(router.children["GET"].children["accounts"].wildcard.children).To(HaveLen(2))
			Expect(router.children["GET"].children["accounts"].wildcard.children).To(HaveKey("detail"))
			Expect(router.children["GET"].children["accounts"].wildcard.children["detail"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["accounts"].wildcard.children["detail"].children).To(BeEmpty())
			Expect(router.children["GET"].children["accounts"].wildcard.children["detail"].handler).NotTo(BeNil())
			Expect(router.children["GET"].children["accounts"].wildcard.children["detail"].names).To(Equal([]string{"account"}))
			Expect(router.children["GET"].children["accounts"].wildcard.children).To(HaveKey("history"))
			Expect(router.children["GET"].children["accounts"].wildcard.children["history"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["accounts"].wildcard.children["history"].children).To(BeEmpty())
			Expect(router.children["GET"].children["accounts"].wildcard.children["history"].handler).NotTo(BeNil())
			Expect(router.children["GET"].children["accounts"].wildcard.children["history"].names).To(Equal([]string{"account"}))
			Expect(router.children["GET"].wildcard).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.handler).To(BeNil())
			Expect(router.children["GET"].wildcard.children).To(HaveKey("invoice"))
			Expect(router.children["GET"].wildcard.children["invoice"].wildcard).To(BeNil())
			Expect(router.children["GET"].wildcard.children["invoice"].children).To(BeEmpty())
			Expect(router.children["GET"].wildcard.children["invoice"].handler).NotTo(BeNil())
			Expect(router.children["GET"].wildcard.children["invoice"].names).To(Equal([]string{"transaction"}))
		})

		It("should panic due to conflicting empty tokens", func() {
			router := NewRouter()

			Expect(func() {
				router.GET("//detail", emptyHandler)
			}).To(Panic())

			Expect(func() {
				router.GET("/account/detail//", emptyHandler)
			}).To(Panic())

			Expect(func() {
				router.GET("/account//detail", emptyHandler)
			}).To(Panic())
		})

		It("should not panic with empty token at the end", func() {
			router := NewRouter()

			Expect(func() {
				router.GET("/account/", emptyHandler)
			}).NotTo(Panic())

			Expect(func() {
				router.GET("/account/detail/", emptyHandler)
			}).NotTo(Panic())

			Expect(func() {
				router.GET("/account/detail/:id/", emptyHandler)
			}).NotTo(Panic())
		})

		It("should panic due to conflicting static routes", func() {
			router := NewRouter()
			router.GET("/account/detail", emptyHandler)
			Expect(func() {
				router.GET("/account/detail", emptyHandler)
			}).To(Panic())
		})

		It("should panic due to conflicting 'wildcarded' routes", func() {
			router := NewRouter()
			router.GET("/:account", emptyHandler)
			Expect(func() {
				router.GET("/:transaction", emptyHandler)
			}).To(Panic())
		})

		It("should panic due to conflicting mixing routes", func() {
			router := NewRouter()
			router.GET("/:account/detail", emptyHandler)
			router.GET("/:account/id", emptyHandler)
			Expect(func() {
				router.GET("/:transaction/id", emptyHandler)
			}).To(Panic())
		})
	})

	Describe("Group", func() {
		It("should parse a GET", func() {
			router := NewRouter()
			group := router.Group("/group")
			group.GET("/route", emptyHandler)

			Expect(router.children).To(HaveKey("GET"))
			Expect(router.children["GET"].wildcard).To(BeNil())
			Expect(router.children["GET"].children).To(HaveLen(1))
			Expect(router.children["GET"].children).To(HaveKey("group"))
			Expect(router.children["GET"].children["group"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["group"].handler).To(BeNil())
			Expect(router.children["GET"].children["group"].children).To(HaveLen(1))
			Expect(router.children["GET"].children["group"].children).To(HaveKey("route"))
			Expect(router.children["GET"].children["group"].children["route"].wildcard).To(BeNil())
			Expect(router.children["GET"].children["group"].children["route"].handler).NotTo(BeNil())
			Expect(router.children["GET"].children["group"].children["route"].children).To(BeEmpty())
		})

		It("should parse a POST", func() {
			router := NewRouter()
			group := router.Group("/group")
			group.POST("/route", emptyHandler)

			Expect(router.children).To(HaveKey("POST"))
			Expect(router.children["POST"].wildcard).To(BeNil())
			Expect(router.children["POST"].children).To(HaveLen(1))
			Expect(router.children["POST"].children).To(HaveKey("group"))
			Expect(router.children["POST"].children["group"].wildcard).To(BeNil())
			Expect(router.children["POST"].children["group"].handler).To(BeNil())
			Expect(router.children["POST"].children["group"].children).To(HaveLen(1))
			Expect(router.children["POST"].children["group"].children).To(HaveKey("route"))
			Expect(router.children["POST"].children["group"].children["route"].wildcard).To(BeNil())
			Expect(router.children["POST"].children["group"].children["route"].handler).NotTo(BeNil())
			Expect(router.children["POST"].children["group"].children["route"].children).To(BeEmpty())
		})

		It("should parse a PUT", func() {
			router := NewRouter()
			group := router.Group("/group")
			group.PUT("/route", emptyHandler)

			Expect(router.children).To(HaveKey("PUT"))
			Expect(router.children["PUT"].wildcard).To(BeNil())
			Expect(router.children["PUT"].children).To(HaveLen(1))
			Expect(router.children["PUT"].children).To(HaveKey("group"))
			Expect(router.children["PUT"].children["group"].wildcard).To(BeNil())
			Expect(router.children["PUT"].children["group"].handler).To(BeNil())
			Expect(router.children["PUT"].children["group"].children).To(HaveLen(1))
			Expect(router.children["PUT"].children["group"].children).To(HaveKey("route"))
			Expect(router.children["PUT"].children["group"].children["route"].wildcard).To(BeNil())
			Expect(router.children["PUT"].children["group"].children["route"].handler).NotTo(BeNil())
			Expect(router.children["PUT"].children["group"].children["route"].children).To(BeEmpty())
		})

		It("should parse a DELETE", func() {
			router := NewRouter()
			group := router.Group("/group")
			group.DELETE("/route", emptyHandler)

			Expect(router.children).To(HaveKey("DELETE"))
			Expect(router.children["DELETE"].wildcard).To(BeNil())
			Expect(router.children["DELETE"].children).To(HaveLen(1))
			Expect(router.children["DELETE"].children).To(HaveKey("group"))
			Expect(router.children["DELETE"].children["group"].wildcard).To(BeNil())
			Expect(router.children["DELETE"].children["group"].handler).To(BeNil())
			Expect(router.children["DELETE"].children["group"].children).To(HaveLen(1))
			Expect(router.children["DELETE"].children["group"].children).To(HaveKey("route"))
			Expect(router.children["DELETE"].children["group"].children["route"].wildcard).To(BeNil())
			Expect(router.children["DELETE"].children["group"].children["route"].handler).NotTo(BeNil())
			Expect(router.children["DELETE"].children["group"].children["route"].children).To(BeEmpty())
		})

		It("should parse a HEAD", func() {
			router := NewRouter()
			group := router.Group("/group")
			group.HEAD("/route", emptyHandler)

			Expect(router.children).To(HaveKey("HEAD"))
			Expect(router.children["HEAD"].wildcard).To(BeNil())
			Expect(router.children["HEAD"].children).To(HaveLen(1))
			Expect(router.children["HEAD"].children).To(HaveKey("group"))
			Expect(router.children["HEAD"].children["group"].wildcard).To(BeNil())
			Expect(router.children["HEAD"].children["group"].handler).To(BeNil())
			Expect(router.children["HEAD"].children["group"].children).To(HaveLen(1))
			Expect(router.children["HEAD"].children["group"].children).To(HaveKey("route"))
			Expect(router.children["HEAD"].children["group"].children["route"].wildcard).To(BeNil())
			Expect(router.children["HEAD"].children["group"].children["route"].handler).NotTo(BeNil())
			Expect(router.children["HEAD"].children["group"].children["route"].children).To(BeEmpty())
		})

		It("should parse a OPTIONS", func() {
			router := NewRouter()
			group := router.Group("/group")
			group.OPTIONS("/route", emptyHandler)

			Expect(router.children).To(HaveKey("OPTIONS"))
			Expect(router.children["OPTIONS"].wildcard).To(BeNil())
			Expect(router.children["OPTIONS"].children).To(HaveLen(1))
			Expect(router.children["OPTIONS"].children).To(HaveKey("group"))
			Expect(router.children["OPTIONS"].children["group"].wildcard).To(BeNil())
			Expect(router.children["OPTIONS"].children["group"].handler).To(BeNil())
			Expect(router.children["OPTIONS"].children["group"].children).To(HaveLen(1))
			Expect(router.children["OPTIONS"].children["group"].children).To(HaveKey("route"))
			Expect(router.children["OPTIONS"].children["group"].children["route"].wildcard).To(BeNil())
			Expect(router.children["OPTIONS"].children["group"].children["route"].handler).NotTo(BeNil())
			Expect(router.children["OPTIONS"].children["group"].children["route"].children).To(BeEmpty())
		})

		It("should parse a PATCH", func() {
			router := NewRouter()
			group := router.Group("/group")
			group.PATCH("/route", emptyHandler)

			Expect(router.children).To(HaveKey("PATCH"))
			Expect(router.children["PATCH"].wildcard).To(BeNil())
			Expect(router.children["PATCH"].children).To(HaveLen(1))
			Expect(router.children["PATCH"].children).To(HaveKey("group"))
			Expect(router.children["PATCH"].children["group"].wildcard).To(BeNil())
			Expect(router.children["PATCH"].children["group"].handler).To(BeNil())
			Expect(router.children["PATCH"].children["group"].children).To(HaveLen(1))
			Expect(router.children["PATCH"].children["group"].children).To(HaveKey("route"))
			Expect(router.children["PATCH"].children["group"].children["route"].wildcard).To(BeNil())
			Expect(router.children["PATCH"].children["group"].children["route"].handler).NotTo(BeNil())
			Expect(router.children["PATCH"].children["group"].children["route"].children).To(BeEmpty())
		})

		It("should check the subgroup", func() {
			router := NewRouter()
			group := router.Group("/group")
			group2 := group.Group("/subgroup").(*routerGroup)

			Expect(group2).NotTo(BeNil())
			Expect(group2.router).To(Equal(group))
			Expect(group2.prefix).To(Equal("/subgroup"))
		})
	})

	Describe("Handle", func() {
		var router *Router

		BeforeEach(func() {
			router = NewRouter()
		})

		It("should resolve a static route", func() {
			value := 1
			router.GET("/static", func(ctx *fasthttp.RequestCtx) {
				value = 2
			})

			router.Handler(createRequestCtxFromPath("GET", "/static"))

			Expect(value).To(Equal(2))
		})

		It("should resolve a static route not starting with /", func() {
			value := 1
			router.GET("static", func(ctx *fasthttp.RequestCtx) {
				value = 2
			})

			router.Handler(createRequestCtxFromPath("GET", "/static"))

			Expect(value).To(Equal(2))
		})

		It("should resolve multiple static routes", func() {
			value1 := 1
			value2 := 1
			value3 := 1

			router.GET("/static", func(ctx *fasthttp.RequestCtx) {
				value1 = 2
			})

			router.GET("/static/second", func(ctx *fasthttp.RequestCtx) {
				value2 = 2
			})

			router.GET("/another", func(ctx *fasthttp.RequestCtx) {
				value3 = 2
			})

			router.Handler(createRequestCtxFromPath("GET", "/static"))
			router.Handler(createRequestCtxFromPath("GET", "/static/second"))
			router.Handler(createRequestCtxFromPath("GET", "/another"))

			Expect(value1).To(Equal(2))
			Expect(value2).To(Equal(2))
			Expect(value3).To(Equal(2))
		})

		It("should resolve a wildcard route", func() {
			value := 1
			router.GET("/:wildcard", func(ctx *fasthttp.RequestCtx) {
				Expect(ctx.UserValue("wildcard")).To(Equal("value"))
				value = 2
			})

			router.Handler(createRequestCtxFromPath("GET", "/value"))

			Expect(value).To(Equal(2))
		})

		It("should resolve a multiple wildcard routes", func() {
			value1 := 1
			value2 := 1
			value3 := 1
			router.GET("/:account/transactions", func(ctx *fasthttp.RequestCtx) {
				Expect(ctx.UserValue("account")).To(Equal("value1"))
				value1 = 2
			})
			router.GET("/:account/profile", func(ctx *fasthttp.RequestCtx) {
				Expect(ctx.UserValue("account")).To(Equal("value2"))
				value2 = 2
			})
			router.GET("/:user/roles", func(ctx *fasthttp.RequestCtx) {
				Expect(ctx.UserValue("user")).To(Equal("value3"))
				value3 = 2
			})

			router.Handler(createRequestCtxFromPath("GET", "/value1/transactions"))
			router.Handler(createRequestCtxFromPath("GET", "/value2/profile"))
			router.Handler(createRequestCtxFromPath("GET", "/value3/roles"))

			Expect(value1).To(Equal(2))
			Expect(value2).To(Equal(2))
			Expect(value3).To(Equal(2))
		})

		It("should resolve a multiple wildcard routes", func() {
			value1 := 1
			value2 := 1
			value3 := 1
			router.GET("/:account/transactions", func(ctx *fasthttp.RequestCtx) {
				Expect(ctx.UserValue("account")).To(Equal("value1"))
				value1 = 2
			})
			router.GET("/:account/profile", func(ctx *fasthttp.RequestCtx) {
				Expect(ctx.UserValue("account")).To(Equal("value2"))
				value2 = 2
			})
			router.GET("/:user/roles", func(ctx *fasthttp.RequestCtx) {
				Expect(ctx.UserValue("user")).To(Equal("value3"))
				value3 = 2
			})

			router.Handler(createRequestCtxFromPath("GET", "/value1/transactions"))
			router.Handler(createRequestCtxFromPath("GET", "/value2/profile"))
			router.Handler(createRequestCtxFromPath("GET", "/value3/roles"))

			Expect(value1).To(Equal(2))
			Expect(value2).To(Equal(2))
			Expect(value3).To(Equal(2))
		})

		It("should call the not found callback for static routes", func() {
			value1 := 1

			router.GET("/account/transactions", func(ctx *fasthttp.RequestCtx) {
				Fail("should not be called")
			})

			router.NotFound = func(ctx *fasthttp.RequestCtx) {
				value1 = 2
			}
			router.Handler(createRequestCtxFromPath("GET", "/account/transactions_notfound"))

			Expect(value1).To(Equal(2))
		})

		It("should call the not found callback for static routes half path", func() {
			value1 := 1

			router.GET("/account/transactions", func(ctx *fasthttp.RequestCtx) {
				Fail("should not be called")
			})

			router.NotFound = func(ctx *fasthttp.RequestCtx) {
				value1 = 2
			}
			router.Handler(createRequestCtxFromPath("GET", "/account"))

			Expect(value1).To(Equal(2))
		})

		It("should call the not found callback for wildcard routes", func() {
			value1 := 1
			value2 := 1
			value3 := 1

			router.GET("/:account/transactions", func(ctx *fasthttp.RequestCtx) {
				Fail("should not be called")
			})
			router.GET("/:account/profile", func(ctx *fasthttp.RequestCtx) {
				Fail("should not be called")
			})
			router.GET("/:user/roles", func(ctx *fasthttp.RequestCtx) {
				Fail("should not be called")
			})

			router.NotFound = func(ctx *fasthttp.RequestCtx) {
				value1 = 2
			}
			router.Handler(createRequestCtxFromPath("GET", "/value1/transactions_notfound"))

			router.NotFound = func(ctx *fasthttp.RequestCtx) {
				value2 = 2
			}
			router.Handler(createRequestCtxFromPath("GET", "/value2/profile_notfound"))

			router.NotFound = func(ctx *fasthttp.RequestCtx) {
				value3 = 2
			}
			router.Handler(createRequestCtxFromPath("GET", "/value3/roles_notfound"))

			Expect(value1).To(Equal(2))
			Expect(value2).To(Equal(2))
			Expect(value3).To(Equal(2))
		})

		It("should call the not found callback for wildcard half path", func() {
			value1 := 1

			router.GET("/:account/transactions", func(ctx *fasthttp.RequestCtx) {
				Fail("should not be called")
			})

			router.NotFound = func(ctx *fasthttp.RequestCtx) {
				value1 = 2
			}
			router.Handler(createRequestCtxFromPath("GET", "/value1"))

			Expect(value1).To(Equal(2))
		})

		It("should call the not found callback for wrong method", func() {
			value1 := 1

			router.GET("/:account/transactions", func(ctx *fasthttp.RequestCtx) {
				Fail("should not be called")
			})

			router.NotFound = func(ctx *fasthttp.RequestCtx) {
				value1 = 2
			}
			router.Handler(createRequestCtxFromPath("POST", "/value1"))

			Expect(value1).To(Equal(2))
		})
	})
})
