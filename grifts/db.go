package grifts

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/toodo/models"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		return models.DB.Transaction(func(tx *pop.Connection) error {
			if err := tx.TruncateAll(); err != nil {
				return errors.WithStack(err)
			}

			u := models.User{
				Email:                "mark@example.com",
				Password:             "password",
				PasswordConfirmation: "password",
			}

			if err := tx.Create(&u); err != nil {
				return err
			}

			items := []models.Item{
				models.Item{
					Title:  "Buy Milk",
					UserID: u.ID,
				},
				models.Item{
					Title:  "Learn Go",
					Body:   nulls.NewString("What a great language"),
					UserID: u.ID,
				},
				models.Item{
					Title:     "Learn Buffalo",
					UserID:    u.ID,
					Completed: true,
				},
			}

			for _, i := range items {
				if err := tx.Create(&i); err != nil {
					return err
				}
			}

			return nil
		})
	})

})
