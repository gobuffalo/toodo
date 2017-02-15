package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/toodo/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_toodo_session",
		})

		app.Use(middleware.PopTransaction(models.DB))

		app.ServeFiles("/assets", assetsPath())

		var todoResource buffalo.Resource
		todoResource = TodosResource{&buffalo.BaseResource{}}
		app.Resource("/todos", todoResource)

		app.GET("/", todoResource.List)
	}

	return app
}
