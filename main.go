package main

import (
	"log"

	"github.com/mewben/arrstate/internal/startup"
	"github.com/mewben/arrstate/pkg"
)

func main() {
	log.Println("mainnnn")
	db := startup.Init()
	app := pkg.SetupBackend(db)

	app.Listen("localhost:5000")
}
