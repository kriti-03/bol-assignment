package utility

import "fmt"

type (
	// NotFoundError happens when the request record is not found on the storage.
	NotFoundError struct {
		Msg string
	}
	// InvalidPlayError represents an invalid play error.
	InvalidPlayError struct {
		Msg string
	}
)

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("not found error: %v", e.Msg)
}

func (e *InvalidPlayError) Error() string {
	return fmt.Sprintf("invalid play: %v", e.Msg)
}
