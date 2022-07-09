package main

import (
	"log"

	"github.com/yyh-gl/go-playground/src"
)

func main() {
	db := newMySQLClient()

	u := src.UserDTO{
		Name: "hoge",
		Age:  1,
	}

	if err := db.Create(&u).Error; err != nil {
		log.Fatal(err)
	}
}
