package errors

import (
	"fmt"

	"github.com/mewben/realty278/pkg/services"
)

// HTTPError custom
type HTTPError struct {
	Code      int    `json:"-"`
	Message   string `json:"message"`
	RequestID string `json:"requestID"`
	Internal  error  `json:"-"`
}

// Error makes it compatible with `error` interface.
func (e *HTTPError) Error() string {
	if e.Internal == nil {
		return fmt.Sprintf("code=%d, message=%v, requestID=%v", e.Message)
	}
	return fmt.Sprintf("message=%v, internal=%v", e.Message, e.Internal)
}

// NewHTTPError error -
func NewHTTPError(args ...interface{}) error {
	e := &HTTPError{}
	// Defaults
	e.Message = ErrDefault

	for _, arg := range args {
		switch arg := arg.(type) {
		case int:
			e.Code = arg
		case string:
			e.Message = arg
		case error:
			e.Internal = arg
		}
	}

	// TODO: log to external service or file
	// log.Println("-err", e)

	// TODO: Translate Message
	e.Message = services.T(e.Message)
	return e
}
