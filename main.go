package main

import (
	"fmt"

	"github.com/edison626/vblog/config"
	"github.com/edison626/vblog/domain"
	"github.com/edison626/vblog/vblog"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	greetertest := &domain.Greeter{Greeting: conf.Greeting}
	fmt.Println(greetertest)
	vblog.StartVblog(greetertest)
}
