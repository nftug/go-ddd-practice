package entity

import "domain/vo"

type User struct {
	Id       vo.Id
	FullName vo.FullName
	Email    vo.Email
}

func CreateUser(firstName string, lastName string, email string) (*User, error) {
	return NewUser(0, firstName, lastName, email)
}

func (u *User) Edit(firstName string, lastName string, email string) (*User, error) {
	return NewUser(u.Id.Get(), firstName, lastName, email)
}

func NewUser(id int, firstName string, lastName string, email string) (*User, error) {
	_id, err := vo.NewId(id)
	if err != nil {
		return nil, err
	}
	fullName, err := vo.NewFullName(firstName, lastName)
	if err != nil {
		return nil, err
	}
	_email, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{Id: *_id, FullName: *fullName, Email: *_email}, nil
}
