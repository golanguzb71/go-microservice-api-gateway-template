package helpers

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be blank")
	}
	if len(password) < 4 || len(password) > 30 {
		return errors.New("password length should be 4 to 30 characters")
	}
	if validation.Validate(password, validation.Match(regexp.MustCompile("^[A-Za-z0-9$_@.#-]+$"))) != nil {
		return errors.New("password should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	return nil
}

func ValidateLogin(login string) error {
	if login == "" {
		return errors.New("login cannot be blank")
	}
	if len(login) < 6 && len(login) > 15 {
		return errors.New("login length should be 8 to 30 characters")
	}
	if validation.Validate(login, validation.Match(regexp.MustCompile("^[A-Za-z0-9$@_.#]+$"))) != nil {
		return errors.New("login should contain only alphabetic characters, numbers and special characters(@, $, _, ., #)")
	}
	return nil
}
