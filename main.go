package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	tmpl := iris.HTML("./static", ".html")
	app.RegisterView(tmpl)

	app.StaticWeb("/", "./static")

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "Hello Iris",
		})
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":3000"))
}
