package main

import (
	"log"

	"github.com/gobuffalo/toodo/actions"
)

func main() {
	log.Fatal(actions.App().Serve())
}
