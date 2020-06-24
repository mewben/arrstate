package main

import (
	"log"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
)

func main() {
	log.Println("mainnnn")
	db := startup.Init()
	app := pkg.SetupBackend(db)

	app.Listen("localhost:5000")
}
