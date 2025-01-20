package validate

import (
	"errors"
)

var (
	BlankFullName      = errors.New("Full name is required")
	LengthFullName     = errors.New("Full name must be less than 50 characters")
	InvalidPhoneNumber = errors.New("Invalid phone number. Please enter a 10-digit phone number.")
	InvalidGender      = errors.New("Invalid gender. Please enter Male or Female.")
)

func ValidateFullName(name string) error {
	if name == "" {
		return BlankFullName
	}
	if len(name) > 50 {
		return LengthFullName
	}
	return nil
}
func ValidatePhoneNumber(phoneNumber string) error {
	if len(phoneNumber) != 10 {
		return InvalidPhoneNumber
	}
	return nil
}
func ValidateGender(gender string) error {
	if gender != "Male" && gender != "Female" {
		return InvalidGender
	}
	return nil
}
