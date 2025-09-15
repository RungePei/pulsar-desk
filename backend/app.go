package backend

import (
	"context"
)

var app = &App{}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return app
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func Startup(ctx context.Context) {
	app.ctx = ctx
}
