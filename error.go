package libyggdrasil

import (
	"fmt"
)

type YggdrasilAuthenticationError struct {
	ResponseCode int
	ErrorName    string
	ErrorMessage string
	Cause        string
}

func (e *YggdrasilAuthenticationError) Error() string {
	return fmt.Sprintf("Response code %d: (%s) \"%s\"", e.ResponseCode, e.ErrorName, e.ErrorMessage)
}
