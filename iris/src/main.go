package main

import "github.com/kataras/iris/v12"

func get(ctx iris.Context) {
	ctx.WriteString("Hello, World!")
}

func main() {
	app := iris.New()

	router := app.Party("/")
	{
		router.Use(iris.Compression)
		router.Get("/", get)
	}

	app.Listen("0.0.0.0:8000")
}
