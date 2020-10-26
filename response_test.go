package render

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	// correct response
	e := Response{
		Code:   http.StatusForbidden,
		Status: http.StatusText(http.StatusForbidden),
	}
	assert.Equal(t, e, NewResponse(http.StatusForbidden, nil))

	// false positive
	e1 := Response{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusOK),
	}
	assert.NotEqual(t, e1, NewResponse(http.StatusNotFound, nil))

	ev := "internal server error test"
	e2 := Response{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: ev,
	}
	assert.Equal(t, e2, NewResponse(http.StatusInternalServerError, ev))

	// validation errors
	ves := validator.ValidationErrors{
		&fe{"required", "required_field", ""},
		&fe{"len", "len_field", "6"},
	}
	e3 := Response{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: map[string]string{
			"required_field": "This value is required.",
			"len_field": "This value failed validation on 'len:6'.",
		},
	}
	assert.Equal(t, e3, NewResponse(http.StatusInternalServerError, ves))

	// msg of type error
	e4 := Response{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: "this is an error",
	}
	assert.Equal(t, e4, NewResponse(http.StatusInternalServerError, errors.New("this is an error")))
}

type fe struct {
	tag   string
	field string
	param string
}

func (fe *fe) Tag() string { return fe.tag }

func (fe *fe) ActualTag() string { return fe.tag }

func (fe *fe) Namespace() string { return "" }

func (fe *fe) StructNamespace() string { return "" }

func (fe *fe) Field() string { return fe.field }

func (fe *fe) StructField() string { return "" }

func (fe *fe) Value() interface{} { return "" }

func (fe *fe) Param() string { return fe.param }

func (fe *fe) Kind() reflect.Kind { return 0 }

func (fe *fe) Type() reflect.Type { return nil }

func (fe *fe) Translate(ut ut.Translator) string { return "" }
