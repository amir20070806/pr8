package main

import (
	"fmt"
	"strings"
)

type Validator interface {
	Validate(value string) bool
}

type RequiredValidator struct{}

func (r RequiredValidator) Validate(value string) bool {
	return strings.TrimSpace(value) != ""
}

type LengthValidator struct {
	Min, Max int
}

func (l LengthValidator) Validate(value string) bool {
	length := len(value)
	return length >= l.Min && length <= l.Max
}

type EmailValidator struct{}

func (e EmailValidator) Validate(value string) bool {
	return strings.Contains(value, "@") && strings.Contains(value, ".")
}

type Field struct {
	Name       string
	Value      string
	Validators []Validator
}

func (f *Field) Validate() bool {
	for _, validator := range f.Validators {
		if !validator.Validate(f.Value) {
			return false
		}
	}
	return true
}

func main() {
	fields := []Field{
		{
			Name:  "username",
			Value: "user123",
			Validators: []Validator{
				RequiredValidator{},
				LengthValidator{Min: 3, Max: 20},
			},
		},
		{
			Name:  "email",
			Value: "test@example.com",
			Validators: []Validator{
				RequiredValidator{},
				EmailValidator{},
			},
		},
	}

	for _, field := range fields {
		if field.Validate() {
			fmt.Printf("%s: valid\n", field.Name)
		} else {
			fmt.Printf("%s: invalid\n", field.Name)
		}
	}
}
