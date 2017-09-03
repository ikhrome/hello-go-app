package main

import (
	"runtime"
	"time"

	"github.com/kataras/iris"
)

const (
	// DefaultTitle is the default website title
	DefaultTitle = "Tech Support GoLang App"
	// DefaultLayout is the default website layout
	DefaultLayout = "layouts/layout.gohtml"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./templates", ".gohtml").Reload(true))

	app.StaticWeb("/static", "./assets")

	app.Use(func(ctx iris.Context) {
		ctx.ViewData("Title", DefaultTitle)
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.ViewLayout(DefaultLayout)
		ctx.Header("X-Pechenki", "NO_PECHENKI_SUKA_BLYAT")

		ctx.Next()
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("BodyMessage", "Welcome to the world of Go!")
		ctx.ViewData("GoVersion", runtime.Version())
		if err := ctx.View("index.gohtml"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	})

	app.Run(iris.Addr(":8080"))
}
