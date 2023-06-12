package vo

import (
	"fmt"
	"reflect"
	"regexp"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) (*FullName, error) {
	validateName := func(v string, fieldName string) error {
		if v == "" {
			return fmt.Errorf("%s is required", fieldName)
		}
		if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(v) {
			return fmt.Errorf("%s has an invalid character", fieldName)
		}
		return nil
	}

	if err := validateName(firstName, "firstName"); err != nil {
		return nil, err
	}
	if err := validateName(lastName, "lastName"); err != nil {
		return nil, err
	}

	return &FullName{firstName, lastName}, nil
}

func (f *FullName) Get() string {
	return f.firstName + " " + f.lastName
}

func (f *FullName) Equals(o *FullName) bool {
	return reflect.DeepEqual(f.firstName, o.firstName) &&
		reflect.DeepEqual(f.lastName, o.lastName)
}
