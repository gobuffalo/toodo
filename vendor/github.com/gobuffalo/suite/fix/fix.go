package fix

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/gobuffalo/plush"

	"github.com/BurntSushi/toml"
	"github.com/gobuffalo/packd"
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
			return err
		}

		sc := Scenarios{}
		_, err = toml.Decode(x, &sc)
		if err != nil {
			return err
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
			return err
		}

		sc := Scenarios{}
		_, err = toml.Decode(x, &sc)
		if err != nil {
			return err
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
		return Scenario{}, fmt.Errorf("could not find a scenario named %q", name)
	}
	return s, nil
}
