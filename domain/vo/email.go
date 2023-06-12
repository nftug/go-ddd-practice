package vo

import (
	"net/mail"
	"reflect"
)

type Email struct {
	value string
}

func NewEmail(v string) (*Email, error) {
	addr, err := mail.ParseAddress(v)
	if err != nil {
		return nil, err
	}
	return &Email{value: addr.Address}, nil
}

func (e *Email) Get() string {
	return e.value
}

func (e *Email) Equals(o *Email) bool {
	return reflect.DeepEqual(e.value, o.value)
}
