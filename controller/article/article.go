package article

import (
	"h8FinalProject/blueprint/lib/flight"
	"h8FinalProject/blueprint/middleware/acl"
	"h8FinalProject/blueprint/model/article"
	"net/http"

	"github.com/blue-jay/core/pagination"
	"github.com/blue-jay/core/router"
)

var (
	uri = "/article"
)

// Load the routes.
func Load() {
	c := router.Chain(acl.DisallowAnon)
	router.Get(uri, Index, c...)
	router.Get(uri+"/create", Create, c...)
	router.Post(uri+"/create", Store, c...)
	router.Get(uri+"/view/:id", Show, c...)
	router.Get(uri+"/edit/:id", Edit, c...)
	router.Patch(uri+"/edit/:id", Update, c...)
	router.Delete(uri+"/:id", Destroy, c...)
}

// Index displays the articles.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	// Create a pagination instance with a max of 10 results.
	p := pagination.New(r, 10)

	articles, _, err := article.ByUserIDPaginate(c.DB, c.UserID, p.PerPage, p.Offset)
	if err != nil {
		c.FlashErrorGeneric(err)
		articles = []article.Article{}
	}

	count, err := article.ByUserIDCount(c.DB, c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
	}

	// Calculate the number of pages.
	p.CalculatePages(count)

	v := c.View.New("article/index")
	v.Vars["articles"] = articles
	v.Vars["pagination"] = p
	v.Render(w, r)
}

// Create displays the create form.
func Create(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("article/create")
	c.Repopulate(v.Vars, "tittle")
	v.Render(w, r)
}

// Store handles the create form submission.
func Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	if !c.FormValid("tittle", "issian") {
		Create(w, r)
		return
	}

	_, err := article.Create(c.DB, r.FormValue("tittle"), r.FormValue("issian"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		Create(w, r)
		return
	}

	c.FlashSuccess("Article added.")
	c.Redirect(uri)
}

// Show displays a single Article.
func Show(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	article, _, err := article.ByID(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("article/show")
	v.Vars["article"] = article
	v.Render(w, r)
}

// Edit displays the edit form.
func Edit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	article, _, err := article.ByID(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("article/edit")
	c.Repopulate(v.Vars, "tittle")
	v.Vars["article"] = article
	v.Render(w, r)
}

// Update handles the edit form submission.
func Update(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	if !c.FormValid("tittle", "issian") {
		Edit(w, r)
		return
	}

	_, err := article.Update(c.DB, r.FormValue("tittle"), r.FormValue("issian"), c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
		Edit(w, r)
		return
	}

	c.FlashSuccess("Article updated.")
	c.Redirect(uri)
}

// Destroy handles the delete form submission.
func Destroy(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	_, err := article.DeleteSoft(c.DB, c.Param("id"), c.UserID)
	if err != nil {
		c.FlashErrorGeneric(err)
	} else {
		c.FlashNotice("Article deleted.")
	}

	c.Redirect(uri)
}
