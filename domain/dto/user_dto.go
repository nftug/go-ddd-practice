package dto

import "domain/entity"

type UserDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserDto(u *entity.User) UserDto {
	return UserDto{
		ID:    u.ID.Get(),
		Name:  u.FullName.Get(),
		Email: u.Email.Get(),
	}
}
