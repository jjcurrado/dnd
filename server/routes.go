package server

import (
	"net/http"
)

func CharacterRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CharacterPostHandler(w, r)
	}
}
