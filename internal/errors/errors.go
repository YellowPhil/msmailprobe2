package errors

import "fmt"

type (
	ErrSomethingWrong struct {
		StatusCode int
	}
	ErrUnauthorized struct{}
)

func (errAD *ErrUnauthorized) Error() string {
	return "got 401 response status code"
}

func (errSW *ErrSomethingWrong) Error() string {
	return fmt.Sprintf("Status code: %d", errSW.StatusCode)
}
