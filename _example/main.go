package main

import (
	"fmt"

	"github.com/aoisensi/go-gsi"
	"github.com/aoisensi/go-gsi/dota2"
)

func main() {
	s := gsi.NewServer(":58591")
	h := dota2.NewHandler(func(data *dota2.Data) {
		fmt.Println(data.Hero.Health)
	})
	s.AddHandler(h)
	s.Listen()
}
