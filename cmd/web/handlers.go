package main

import (
	"fmt"
	"net/http"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world")
}
