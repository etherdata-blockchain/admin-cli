package errors

import "fmt"

type InvalidTemplateIdError struct {
	TemplateId string
}

func (m *InvalidTemplateIdError) Error() string {
	return fmt.Sprintf("Your template id %s is not valid", m.TemplateId)
}

func NewInvalidTemplateIdError(selection string) *InvalidTemplateIdError {
	return &InvalidTemplateIdError{
		TemplateId: selection,
	}
}
