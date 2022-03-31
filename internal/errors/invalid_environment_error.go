package errors

import "fmt"

type InvalidEnvironmentError struct {
	Selection string
}

func (m *InvalidEnvironmentError) Error() string {
	return fmt.Sprintf("Your selection %s is not a valid environment selection", m.Selection)
}

//NewInvalidEnvironmentError creates an error when the given environment flag is in valid.
//Note that this environment is not the environment in environment variable
func NewInvalidEnvironmentError(selection string) *InvalidEnvironmentError {
	return &InvalidEnvironmentError{
		Selection: selection,
	}
}

func (e *InvalidEnvironmentError) Is(tgt error) bool {
	_, ok := tgt.(*InvalidEnvironmentError)
	if !ok {
		return false
	}
	return ok
}
