package main

import (
	"encoding/json"
	"fmt"
	"go-ddd-practice/domain/dto"
	"go-ddd-practice/domain/entity"
	"log"
)

func main() {
	u, err := entity.CreateUser("Jane", "Doe", "test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	printDto(u)

	u, err = u.Edit("Taro", "Yamada", "test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	printDto(u)
}

func printDto(u *entity.User) {
	j, err := json.Marshal(dto.NewUserDto(u))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
