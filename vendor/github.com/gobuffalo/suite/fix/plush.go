package fix

import (
	"io/ioutil"
	"time"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func renderWithContext(file packd.File, ctx *plush.Context) (string, error) {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.WithStack(err)
	}

	cm := map[string]interface{}{
		"uuid": func() uuid.UUID {
			u, _ := uuid.NewV4()
			return u
		},
		"uuidNamed": uuidNamed,
		"now":       now,
		"hash":      hash,
	}
	for k, v := range cm {
		if !ctx.Has(k) {
			ctx.Set(k, v)
		}
	}
	return plush.Render(string(b), ctx)
}

func render(file packd.File) (string, error) {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return plush.Render(string(b), plush.NewContextWith(map[string]interface{}{
		"uuid": func() uuid.UUID {
			u, _ := uuid.NewV4()
			return u
		},
		"uuidNamed": uuidNamed,
		"now":       now,
		"hash":      hash,
	}))
}

func hash(s string, opts map[string]interface{}, help plush.HelperContext) (string, error) {
	cost := bcrypt.DefaultCost
	if i, ok := opts["cost"].(int); ok {
		cost = i
	}
	ph, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	return string(ph), err
}

func now(help plush.HelperContext) string {
	return time.Now().Format(time.RFC3339)
}

func uuidNamed(name string, help plush.HelperContext) uuid.UUID {
	u, _ := uuid.NewV4()
	if ux, ok := help.Value(name).(uuid.UUID); ok {
		return ux
	}
	help.Set(name, u)
	return u
}
