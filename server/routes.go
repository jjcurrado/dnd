package server

import (
	"fmt"
	"net/http"
)

func CharacterRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL)
	switch r.Method {
	case "POST":
		CharacterPostHandler(w, r)
	}
}
