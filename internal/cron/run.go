package cron

import (
	"log"

	cr "github.com/robfig/cron/v3"
)

// Run cron functions
func Run() {
	v := "hello cron"
	c := cr.New()

	c.AddFunc("0 0 * * *", func() { CalculateSails(v) }) // 12mn daily
	c.Start()
}

func CalculateSails(v string) {
	log.Println("vv: ", v)
}
