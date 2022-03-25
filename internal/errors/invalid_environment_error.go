package errors

import "fmt"

type InvalidEnvironmentError struct {
	Selection string
}

func (m *InvalidEnvironmentError) Error() string {
	return fmt.Sprintf("Your selection %s is not a valid environment selection", m.Selection)
}

func NewInvalidEnvironmentError(selection string) *InvalidEnvironmentError {
	return &InvalidEnvironmentError{
		Selection: selection,
	}
}
