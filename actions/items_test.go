package actions

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/gobuffalo/toodo/models"
	"github.com/markbates/going/randx"
)

type ItemsSuite struct {
	*suite.Action
	User *models.User
}

func (is *ItemsSuite) SetupTest() {
	is.Action.SetupTest()
	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}

	verrs, err := u.Create(is.DB)
	is.NoError(err)
	is.False(verrs.HasAny())

	is.Session.Set("current_user_id", u.ID)
	is.Session.Set("current_user", u)
	is.NoError(is.Session.Save())
	is.User = u
}

func (is *ItemsSuite) TableChange(table string, count int, fn func()) {
	scount, err := is.DB.Count(table)
	is.NoError(err)

	fn()

	ecount, err := is.DB.Count(table)
	is.NoError(err)
	is.Equal(count, ecount-scount)
}

func (is *ItemsSuite) CreateItem() *models.Item {
	item := &models.Item{
		Title:  randx.String(20),
		UserID: is.User.ID,
	}
	is.NoError(is.DB.Create(item))
	return item
}

func (as *ItemsSuite) Test_ItemsResource_List() {
	item := as.CreateItem()

	res := as.HTML("/items").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), item.Title)
}

func (as *ItemsSuite) Test_ItemsResource_Show() {
	item := as.CreateItem()

	res := as.HTML("/items/%s", item.ID).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), item.Title)
}

func (as *ItemsSuite) Test_ItemsResource_New() {
	res := as.HTML("/items/new").Get()
	as.Equal(200, res.Code)
}

func (as *ItemsSuite) Test_ItemsResource_Create() {
	item := &models.Item{Title: randx.String(20)}

	as.TableChange("items", 1, func() {
		res := as.HTML("/items").Post(item)
		as.Equal(302, res.Code)
	})
}

func (as *ItemsSuite) Test_ItemsResource_Edit() {
	item := as.CreateItem()
	res := as.HTML("/items/%s/edit", item.ID).Get()
	as.Equal(200, res.Code)
}

func (as *ItemsSuite) Test_ItemsResource_Update() {
	item := as.CreateItem()
	as.TableChange("items", 0, func() {
		res := as.HTML("/items/%s", item.ID).Put(&models.Item{
			ID:    item.ID,
			Title: "new title",
		})
		as.Equal(302, res.Code)

		as.NoError(as.DB.Reload(item))
		as.Equal("new title", item.Title)
	})
}

func (as *ItemsSuite) Test_ItemsResource_Destroy() {
	item := as.CreateItem()
	as.TableChange("items", -1, func() {
		res := as.HTML("/items/%s", item.ID).Delete()
		as.Equal(302, res.Code)
		as.Equal("/items", res.Location())
	})
}

func Test_ItemsSuite(t *testing.T) {
	as := suite.NewAction(App())
	suite.Run(t, &ItemsSuite{Action: as})
}
