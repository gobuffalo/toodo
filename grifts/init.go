package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/toodo/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
