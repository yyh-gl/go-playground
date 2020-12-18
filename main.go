package main

import (
	"fmt"

	"github.com/yyh-gl/playground/src"
)

func main() {
	u := src.NewUser("yyh-gl", 20)
	fmt.Printf("%+v\n", *u)
}
