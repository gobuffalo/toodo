package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	if c.Value("current_user") != nil {
		return c.Redirect(302, "/items")
	}
	return c.Render(200, r.HTML("index.plush.html"))
}
