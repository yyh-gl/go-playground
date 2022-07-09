package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMySQLClient() *gorm.DB {
	dsn := "root:password@tcp(mysql:3306)/playground?charset=utf8mb4,utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func newPostgreSQLClient() *gorm.DB {
	dsn := "host=postgresql user=root password=password dbname=playground port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
