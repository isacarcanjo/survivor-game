package errors

import "fmt"

type UserAlreadyExistsError struct {
	Email  string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("%s: User already exists", e.Email)
}

type InvalidNameError struct {}

func (e *InvalidNameError) Error() string {
	return fmt.Sprint("the name must be mandatory and greater than 3")
}

type InvalidPasswordError struct {}

func (e *InvalidPasswordError) Error() string {
	return fmt.Sprint("the password is invalid. Must have at least 6 characters")
}

type InvalidEmailError struct {}

func (e *InvalidEmailError) Error() string {
	return fmt.Sprint("the email is invalid")
}