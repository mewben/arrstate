package main

import (
	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
)

func main() {
	// startup.InitEnvironment()
	db := startup.Init()
	app := pkg.SetupBackend(db)

	app.Listen("localhost:5000")
}
