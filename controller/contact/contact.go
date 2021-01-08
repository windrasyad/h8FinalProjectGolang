// Package about displays the About page.
package contact

import (
	"net/http"

	"h8FinalProject/blueprint/lib/flight"

	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Get("/contact", Index)
}

// Index displays the About page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	c.View.New("contact/index").Render(w, r)
}
