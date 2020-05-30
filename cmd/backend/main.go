package main

import (
	"log"

	"github.com/mewben/realty278/internal"
	"github.com/mewben/realty278/internal/enums"
	"github.com/mewben/realty278/pkg/models"
)

func main() {
	internal.InitEnvironment()

	business := models.NewBusinessModel()
	log.Println("business", business)
	log.Println("countries", enums.Countries)
}
