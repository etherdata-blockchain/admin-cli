package errors

import (
	"fmt"
)

type InvalidTemplateIdError struct {
	TemplateId string
}

func (e *InvalidTemplateIdError) Error() string {
	return fmt.Sprintf("Your template id %s is not valid", e.TemplateId)
}

func NewInvalidTemplateIdError(selection string) *InvalidTemplateIdError {
	return &InvalidTemplateIdError{
		TemplateId: selection,
	}
}

func (e *InvalidTemplateIdError) Is(tgt error) bool {
	_, ok := tgt.(*InvalidTemplateIdError)
	if !ok {
		return false
	}
	return ok
}
