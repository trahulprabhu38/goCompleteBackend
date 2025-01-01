package main

import "net/http"

func (app *application) healthCareHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("every this is good"))
}
