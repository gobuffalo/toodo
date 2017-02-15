package actions_test

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/gobuffalo/toodo/actions"
)

type ActionSuite struct {
	*suite.ActionSuite
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.New(actions.App())}
	suite.Run(t, as)
}
