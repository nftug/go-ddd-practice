package main

import (
	"domain/dto"
	"domain/entity"
	"encoding/json"

	"fmt"
	"log"
)

func main() {
	u, err := entity.NewUser("Jane", "Doe", "test@example.com")
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
