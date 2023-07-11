package handlers

import (
	"net/http"
	"github.com/malekmajkl/go-learning/WebApp/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")
}
