package dto

import "domain/entity"

type UserDto struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserDto(u *entity.User) UserDto {
	return UserDto{
		Id:    u.Id.Get(),
		Name:  u.FullName.Get(),
		Email: u.Email.Get(),
	}
}
