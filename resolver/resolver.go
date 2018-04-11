package resolver

import "github.com/jehaby/webapp102/service"

type Resolver struct {
	app *service.App
}

func NewRootResolver(app *service.App) *Resolver {
	return &Resolver{
		app: app,
	}
}
