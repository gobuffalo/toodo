package fix

import (
	"path/filepath"
	"sync"

	"github.com/gobuffalo/plush"

	"github.com/BurntSushi/toml"
	"github.com/gobuffalo/packd"
	"github.com/pkg/errors"
)

var scenes = map[string]Scenario{}
var moot = &sync.RWMutex{}

func InitWithContext(box packd.Walkable, ctx *plush.Context) error {
	err := box.Walk(func(path string, file packd.File) error {
		if filepath.Ext(path) != ".toml" {
			return nil
		}

		x, err := renderWithContext(file, ctx)
		if err != nil {
			return errors.Wrap(errors.WithStack(err), path)
		}

		sc := Scenarios{}
		_, err = toml.Decode(x, &sc)
		if err != nil {
			return errors.WithStack(err)
		}

		moot.Lock()
		for _, s := range sc.Scenarios {
			scenes[s.Name] = s
		}
		moot.Unlock()
		return nil
	})
	return err
}

func Init(box packd.Walkable) error {
	err := box.Walk(func(path string, file packd.File) error {
		if filepath.Ext(path) != ".toml" {
			return nil
		}

		x, err := render(file)
		if err != nil {
			return errors.Wrap(errors.WithStack(err), path)
		}

		sc := Scenarios{}
		_, err = toml.Decode(x, &sc)
		if err != nil {
			return errors.WithStack(err)
		}

		moot.Lock()
		for _, s := range sc.Scenarios {
			scenes[s.Name] = s
		}
		moot.Unlock()
		return nil
	})
	return err
}

func Find(name string) (Scenario, error) {
	moot.RLock()
	s, ok := scenes[name]
	moot.RUnlock()
	if !ok {
		return Scenario{}, errors.Errorf("could not find a scenario named %q", name)
	}
	return s, nil
}
