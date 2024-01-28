package router

import "github.com/gofiber/fiber/v2"

type Router interface {
	Initialize(app *fiber.App)
}

func InstallRouter(app *fiber.App) {
	setup(app, NewProductPouter())
}

func setup(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.Initialize(app)
	}
}
