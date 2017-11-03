package actions

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var assetsBox = packr.NewBox("../public/assets")
var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:   "application.html",
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,
		Helpers:      render.Helpers{},
	})
}
