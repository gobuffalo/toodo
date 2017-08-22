package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
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
		// Automatically save the session if the underlying
		// Handler does not return an error.
		app.Use(middleware.SessionSaver)

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		if ENV != "test" {
			// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
			// Remove to disable this.
			app.Use(csrf.Middleware)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))

		todoResource := TodosResource{&buffalo.BaseResource{}}
		app.Resource("/todos", todoResource)

		app.GET("/", todoResource.List)
	}

	return app
}
