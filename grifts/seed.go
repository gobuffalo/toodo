package grifts

import (
	"github.com/gobuffalo/toodo/models"
	. "github.com/markbates/grift/grift"
)

var _ = Add("seed", func(c *Context) error {
	err := models.DB.RawQuery("delete from todos").Exec()
	if err != nil {
		return err
	}
	todo := &models.Todo{Title: "Buy Milk"}
	verrs, err := models.DB.ValidateAndCreate(todo)
	if verrs.HasAny() {
		return verrs
	}
	return err
})
