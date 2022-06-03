package RESTful

// import (
// 	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/framework"
// 	routing "github.com/go-ozzo/ozzo-routing"
// 	"github.com/go-ozzo/ozzo-routing/file"
// 	"github.com/valyala/fasthttp"
// )

// /*with auth query
// batch(...)
// batches(...)
// bts
// withauthmutation
// cancelorder()
// markbatchoutfordelivery("--")
// markorderbatchoutfordelivery
// updatebatchdeliverystatus
// updateorderbatchdelvierystatus

// */

// func init() {
// 	appCtx := framework.GetCurrentAppContext()
// 	router := appCtx.Router

// 	// serve RESTful APIs
// 	api := router.Group("/api")

// 	api.Get("/SampleObject1{id}", func(c *routing.Context) error {
// 		return c.Write("user list")
// 	})
// 	api.Post("/users", func(c *routing.Context) error {
// 		return c.Write("create a new user")
// 	})
// 	api.Put(`/users/<id:\d+>`, func(c *routing.Context) error {
// 		return c.Write("update user " + c.Param("id"))
// 	})

// 	// serve index file
// 	router.Get("/", file.Content("ui/index.html"))
// 	// serve files under the "ui" subdirectory
// 	router.Get("/*", file.Server(file.PathMap{
// 		"/": "/ui/",
// 	}))

// 	// the corresponding fasthttp code
// 	m := func(ctx *fasthttp.RequestCtx) {
// 		switch string(ctx.Path()) {
// 		case "/foo":
// 			fooHandlerFunc(ctx)
// 		case "/bar":
// 			barHandlerFunc(ctx)
// 		case "/baz":
// 			bazHandler.HandlerFunc(ctx)
// 		default:
// 			ctx.Error("not found", fasthttp.StatusNotFound)
// 		}
// 	}

// 	fasthttp.ListenAndServe(":80", m)

// 	fasthttp.ListenAndServe(":8080", router.HandleRequest)

// }
