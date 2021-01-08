// Package controller loads the routes for each of the controllers.
package controller

import (
	"h8FinalProject/blueprint/controller/article"
	"h8FinalProject/blueprint/controller/contact"

	"h8FinalProject/blueprint/controller/debug"

	"h8FinalProject/blueprint/controller/about"
	"h8FinalProject/blueprint/controller/home"
	"h8FinalProject/blueprint/controller/login"
	"h8FinalProject/blueprint/controller/notepad"
	"h8FinalProject/blueprint/controller/register"
	"h8FinalProject/blueprint/controller/static"
	"h8FinalProject/blueprint/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
	contact.Load()
	article.Load()
}
