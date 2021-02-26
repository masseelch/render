package render

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	ContentTypeJson      = "application/json"
	ContentTypeTextPlain = "text/plain"
	ContentTypeTextHtml  = "text/html"
	ContentTypeXml       = "application/xml"
	CharsetSuffix        = "; charset=utf-8"
	HeaderAccept         = "Accept"
	HeaderContentType    = "Content-Type"
)

func BadRequest(w http.ResponseWriter, r *http.Request, msg interface{}) {
	resp := NewResponse(http.StatusBadRequest, msg)
	Render(w, r, resp.Code, resp)
}

func Created(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Render(w, r, http.StatusCreated, msg)
}

func Forbidden(w http.ResponseWriter, r *http.Request, msg interface{}) {
	resp := NewResponse(http.StatusForbidden, msg)
	Render(w, r, resp.Code, resp)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, msg interface{}) {
	resp := NewResponse(http.StatusInternalServerError, msg)
	Render(w, r, resp.Code, resp)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func NotFound(w http.ResponseWriter, r *http.Request, msg interface{}) {
	resp := NewResponse(http.StatusNotFound, msg)
	Render(w, r, resp.Code, resp)
}

func Unauthorized(w http.ResponseWriter, r *http.Request, msg interface{}) {
	resp := NewResponse(http.StatusUnauthorized, msg)
	Render(w, r, resp.Code, resp)
}

func OK(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Render(w, r, http.StatusOK, msg)
}

func PartialContent(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Render(w, r, http.StatusPartialContent, msg)
}

func Render(w http.ResponseWriter, r *http.Request, code int, d interface{}) {
	switch r.Header.Get(HeaderAccept) {
	case ContentTypeXml:
		XML(w, code, d)
	default:
		JSON(w, code, d)
	}
}

func JSON(w http.ResponseWriter, code int, d interface{}) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(HeaderContentType, ContentTypeJson+CharsetSuffix)
	w.WriteHeader(code)
	_, _ = w.Write(buf.Bytes())
}

func XML(w http.ResponseWriter, code int, d interface{}) {
	buf := new(bytes.Buffer)
	if err := xml.NewEncoder(w).Encode(d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set(HeaderContentType, ContentTypeXml+CharsetSuffix)
	w.WriteHeader(code)
	_, _ = w.Write(buf.Bytes())
}

func Raw(w http.ResponseWriter, code int, d []byte) {
	w.Header().Set(HeaderContentType, ContentTypeTextPlain+CharsetSuffix)
	w.WriteHeader(code)
	_, _ = w.Write(d)
}

func HTML(w http.ResponseWriter, code int, d []byte) {
	w.Header().Set(HeaderContentType, ContentTypeTextHtml+CharsetSuffix)
	w.WriteHeader(code)
	_, _ = w.Write(d)
}
