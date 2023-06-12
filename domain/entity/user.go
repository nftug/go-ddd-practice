package entity

import "domain/vo"

type User struct {
	ID vo.Id
	vo.FullName
	vo.Email
}

func NewUser(firstName string, lastName string, email string) (*User, error) {
	fullName, err := vo.NewFullName(firstName, lastName)
	if err != nil {
		return nil, err
	}
	addr, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{FullName: *fullName, Email: *addr}, nil
}

func (u *User) Edit(firstName string, lastName string, email string) (*User, error) {
	fullName, err := vo.NewFullName(firstName, lastName)
	if err != nil {
		return nil, err
	}
	addr, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	u = &User{ID: u.ID, FullName: *fullName, Email: *addr}
	return u, nil
}
