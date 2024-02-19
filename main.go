package main

import (
	"dnd/app"
)

func main() {

	s := app.NewServer()
	s.Start(":8000")
	s.Close()
}
