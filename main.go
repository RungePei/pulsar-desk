package main

import (
	"context"
	"embed"
	"pulsar-desk/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()
	dbService := &backend.DbService{}
	backend.InitDB()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "pulsar desk",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			backend.Startup(ctx)
			dbService.Startup(ctx)
		},
		OnShutdown: func(ctx context.Context) {
			dbService.Shutdown()
		},
		Bind: []interface{}{
			app,
			dbService,
		},
		EnumBind: []interface{}{
			backend.AllLogLevel,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
