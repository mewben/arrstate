package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/mewben/realty278/internal/startup"
	"github.com/mewben/realty278/pkg"
)

func main2() {
	fn("3")
	fn(3)
	fn("4.3")
	fn(4.3)
	fn(-43.7)
	fn("something")
}

func fn(n interface{}) interface{} {
	var x big.Float
	f, b, err := x.Parse(fmt.Sprint(n), 10)
	log.Println("f", f)
	fl, acc := f.Float32()
	log.Println("float32", fl)
	log.Println("float32acc", acc)
	log.Println("b", b)
	log.Println("err", err)
	return f
}

func main() {
	log.Println("baccckend")
	db := startup.Init()
	app := pkg.SetupBackend(db)

	app.Listen("localhost:5000")
}
