package errors

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/requestid"

	"github.com/mewben/realty278/pkg/services"
)

// HTTPError custom
type HTTPError struct {
	Ok        bool   `json:"ok"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestID"`
	Internal  error  `json:"-"`
}

// Error makes it compatible with `error` interface.
func (e *HTTPError) Error() string {
	if e.Internal == nil {
		return fmt.Sprintf("code=%d, message=%v, requestID=%v", e.Code, e.Message, e.RequestID)
	}
	return fmt.Sprintf("code=%d, message=%v, requestID=%v, internal=%v", e.Code, e.Message, e.RequestID, e.Internal)
}

// NewHTTPError error -
func NewHTTPError(c *fiber.Ctx, args ...interface{}) {
	e := &HTTPError{}
	// Defaults
	e.Ok = false
	e.Code = 400
	e.Message = ErrDefault
	e.RequestID = requestid.Get(c)

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
	c.Status(e.Code).JSON(e)
	// No need to send the error next
	c.Next()
}
