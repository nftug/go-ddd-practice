package entity

import "domain/vo"

type User struct {
	ID vo.Id
	vo.FullName
	vo.Email
}

func CreateUser(firstName string, lastName string, email string) (*User, error) {
	u, err := NewUser(0, firstName, lastName, email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Edit(firstName string, lastName string, email string) (*User, error) {
	u, err := NewUser(u.ID.Get(), firstName, lastName, email)
	if err != nil {
		return nil, err
	}
	return u, nil
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
	addr, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{ID: *_id, FullName: *fullName, Email: *addr}, nil
}
