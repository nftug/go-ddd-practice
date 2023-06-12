package vo

import "fmt"

type Id struct {
	value uint
}

func NewId(v int) (*Id, error) {
	if v <= 0 {
		return nil, fmt.Errorf("Id is invalid")
	}
	return &Id{value: uint(v)}, nil
}

func (id *Id) Get() uint {
	return id.value
}

func (id *Id) Equals(o *Id) bool {
	return id.value == o.value
}
