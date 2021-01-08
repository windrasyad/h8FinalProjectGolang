// Package home displays the Home page.
package home

import (
	"net/http"

	"h8FinalProject/blueprint/lib/flight"
	"h8FinalProject/blueprint/model/home"

	"github.com/blue-jay/core/pagination"
	"github.com/blue-jay/core/router"
)

var (
	uri = "/"
)

// Load the routes.
func Load() {
	// c := router.Chain(acl.DisallowAnon)
	router.Get(uri, Index)
	router.Get(uri+"view/:id", Show)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Create a pagination instance with a max of 10 results.
	p := pagination.New(r, 10)

	articles, _, err := home.ByUserIDPaginate(c.DB, c.UserID, p.PerPage, p.Offset)
	if err != nil {
		c.FlashErrorGeneric(err)
		articles = []home.Article{}
	}

	count, err := home.ByUserIDCount(c.DB, c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
	}

	// Calculate the number of pages.
	p.CalculatePages(count)

	v := c.View.New("home/index")
	v.Vars["articles"] = articles
	v.Vars["pagination"] = p
	v.Render(w, r)
}

func Show(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	article, _, err := home.ByID(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("home/show")
	v.Vars["article"] = article
	v.Render(w, r)
}
