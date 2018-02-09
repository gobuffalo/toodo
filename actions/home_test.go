package actions

import "github.com/gobuffalo/toodo/models"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Set("current_user_id", u.ID)

	res := as.HTML("/").Get()
	as.Equal(302, res.Code)
	as.Equal("/items", res.Location())

	as.Session.Clear()
	res = as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}
