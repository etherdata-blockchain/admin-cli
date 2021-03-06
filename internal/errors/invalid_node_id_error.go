package errors

import "fmt"

type InvalidNodeIdError struct {
	NodeId string
}

func (m *InvalidNodeIdError) Error() string {
	return fmt.Sprintf("Your node id %s is invalid", m.NodeId)
}

func NewInvalidNodeIdError(nodeId string) *InvalidNodeIdError {
	return &InvalidNodeIdError{
		NodeId: nodeId,
	}
}

func (e *InvalidNodeIdError) Is(tgt error) bool {
	_, ok := tgt.(*InvalidNodeIdError)
	if !ok {
		return false
	}
	return ok
}
