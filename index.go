package main

import (
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(martini.Static("public"))

	m.Run()
}
