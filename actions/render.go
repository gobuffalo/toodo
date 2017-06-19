package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:   "application.html",
		TemplatesBox: packr.NewBox("../templates"),
		Helpers:      render.Helpers{},
	})
}

func init() {
	binding.RegisterCustomDecorder(func(vals []string) (interface{}, error) {
		fmt.Printf("### vals -> %+v\n", vals)
		// var ti nulls.Time
		//
		// t, err := parseTime(vals)
		// if err != nil {
		// 	return ti, errors.WithStack(err)
		// }
		// ti.Time = t
		// ti.Valid = true
		//
		// return ti, nil
		return nil, nil
	}, []interface{}{new(bool)}, nil)
}
