package errors

import "fmt"

type InvalidPasswordError struct {
	Password string
}

func (m *InvalidPasswordError) Error() string {
	return fmt.Sprintf("Your password %s is invalid", m.Password)
}

func NewInvalidPasswordError(password string) *InvalidPasswordError {
	return &InvalidPasswordError{
		Password: password,
	}
}

func (e *InvalidPasswordError) Is(tgt error) bool {
	_, ok := tgt.(*InvalidPasswordError)
	if !ok {
		return false
	}
	return ok
}
