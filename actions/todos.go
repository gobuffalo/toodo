package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/toodo/models"
	"github.com/markbates/pop"
)

type TodosResource struct {
	buffalo.Resource
}

func (v TodosResource) List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	todos := &models.Todos{}
	err := tx.Order("(case when completed then 1 else 2 end) desc, lower(title) asc").All(todos)
	if err != nil {
		return err
	}
	c.Set("todos", todos)
	return c.Render(200, r.HTML("todos/index.html"))
}

func (v TodosResource) New(c buffalo.Context) error {
	c.Set("todo", &models.Todo{})
	return c.Render(200, r.HTML("todos/new.html"))
}

func (v TodosResource) Create(c buffalo.Context) error {
	t := &models.Todo{}
	err := c.Bind(t)
	if err != nil {
		return err
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := tx.ValidateAndCreate(t)
	if err != nil {
		return err
	}
	if verrs.HasAny() {
		c.Set("todo", t)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("todos/new.html"))
	}
	c.Flash().Add("success", "Todo was created successfully")
	return c.Redirect(301, "/todos")
}

func (v TodosResource) Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	t := &models.Todo{}
	err := tx.Find(t, c.Param("todo_id"))
	if err != nil {
		return err
	}
	err = c.Bind(t)
	if err != nil {
		return err
	}
	err = tx.Update(t)
	if err != nil {
		return err
	}

	c.Set("todo", t)
	return c.Render(200, r.Template("text/javascript", "todos/update.js"))
}
