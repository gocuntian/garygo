package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

type Visitor struct {
	Username string
	Mail     string
	Data     []string `form:"mydata"`
}

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())
	// set the view html template engine
	app.Adapt(view.HTML("./templates", ".html"))

	app.Get("/", func(ctx *iris.Context) {
		if err := ctx.Render("form.html", nil); err != nil {
			ctx.Log(iris.DevMode, err.Error())
		}
	})

	app.Post("/form_action", func(ctx *iris.Context) {
		visitor := Visitor{}
		err := ctx.ReadForm(&visitor)
		if err != nil {
			ctx.Log(iris.DevMode, "Error when reading form: "+err.Error())
		}

		ctx.Writef("Visitor: %#v", visitor)
	})

	app.Listen(":8089")
}