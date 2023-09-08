package vblog

import (
	"fmt"

	"github.com/edison626/vblog/domain"
)

func StartVblog(greeter *domain.Greeter) {
	fmt.Println("RESULT: ", greeter.Greet())
}
