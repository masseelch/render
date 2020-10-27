package render

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Response struct {
	Code   int         `json:"code" xml:"code"`                         // http response status code
	Status string      `json:"status" xml:"status"`                     // user-level status message
	Errors interface{} `json:"errors,omitempty" xml:"errors,omitempty"` // application-level error messages, for debugging
}

func NewResponse(code int, msg interface{}) Response {
	r := Response{
		Code:   code,
		Status: http.StatusText(code),
	}
	switch t := msg.(type) {
	case validator.ValidationErrors:
		m := make(map[string]string, len(t))
		for _, err := range t {
			switch err.Tag() {
			case "required":
				m[err.Field()] = "This value is required."
			case "email":
				m[err.Field()] = "This is not a valid email."
			default:
				s := fmt.Sprintf("This value failed validation on '%s", err.Tag())
				if err.Param() != "" {
					s += fmt.Sprintf(":%s'.", err.Param())
				} else {
					s += "'."
				}
				m[err.Field()] = s
			}
		}
		r.Errors = m
	case error:
		r.Errors = t.Error()
	default:
		r.Errors = t
	}
	return r
}
