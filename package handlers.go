package handlers

import (
	"fmt"
	"net/http"
)

func GetProject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Project handler is working!")
}
