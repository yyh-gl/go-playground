package src

import "gorm.io/gorm"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{
		name: name,
		age:  age,
	}
}

type UserDTO struct {
	gorm.Model
	Name string
	Age  int
}

func (u *UserDTO) TableName() string {
	return "users"
}
