package render

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ErrResponse struct {
	Code   int         `json:"code"`             // http response status code
	Status string      `json:"status"`           // user-level status message
	Errors interface{} `json:"errors,omitempty"` // application-level error messages, for debugging
}

func JSON(w http.ResponseWriter, d interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}

func render(w http.ResponseWriter, resp ErrResponse) {
	w.WriteHeader(resp.Code)
	JSON(w, resp)
}

func Unauthorized(w http.ResponseWriter, err ...interface{}) {
	resp := ErrResponse{
		Code:   http.StatusUnauthorized,
		Status: http.StatusText(http.StatusUnauthorized),
	}
	if len(err) > 0 {
		resp.Errors = err[0]
	}
	render(w, resp)
}

func InternalServerError(w http.ResponseWriter, err ...interface{}) {
	resp := ErrResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
	}
	if len(err) > 0 {
		resp.Errors = err[0]
	}
	render(w, resp)
}
