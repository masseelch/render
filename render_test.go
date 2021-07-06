package render

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func marshalResponse(t *testing.T, r Response) string {
	s, err := json.Marshal(r)
	assert.NoError(t, err)
	return string(s)
}

func TestBadRequest(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rq.Header.Set(HeaderAccept, ContentTypeJson)

	rr := httptest.NewRecorder()
	e := marshalResponse(t, NewResponse(http.StatusBadRequest, errors.New("test")))

	BadRequest(rr, rq, errors.New("test"))
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.JSONEq(t, e, rr.Body.String())
}

func TestBadRequest2(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rq.Header.Set(HeaderAccept, ContentTypeJson)
	rr := httptest.NewRecorder()

	ves := validator.ValidationErrors{
		&fe{"required", "required_field", ""},
		&fe{"len", "len_field", "6"},
	}
	e := marshalResponse(t, NewResponse(http.StatusBadRequest, ves))

	BadRequest(rr, rq, ves)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.JSONEq(t, e, rr.Body.String())
}

func TestForbidden(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	Forbidden(rr, rq, nil)
	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestCreated(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	Created(rr, rq, nil)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestInternalServerError(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	InternalServerError(rr, rq, nil)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestNoContent(t *testing.T) {
	rr := httptest.NewRecorder()

	NoContent(rr)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}

func TestNotFound(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	NotFound(rr, rq, nil)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUnauthorized(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	Unauthorized(rr, rq, nil)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestJSON(t *testing.T) {
	// Tested implicitly with TestBadRequest().
}

func TestRaw(t *testing.T) {
	rr := httptest.NewRecorder()

	Raw(rr, http.StatusOK, []byte("test this"))
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, ContentTypeTextPlain+CharsetSuffix, rr.Header().Get(HeaderContentType))
	assert.Equal(t, []byte("test this"), rr.Body.Bytes())
}

func TestHTML(t *testing.T) {
	rr := httptest.NewRecorder()

	HTML(rr, http.StatusOK, []byte("<html><body>test this</body></html>"))
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, ContentTypeTextHtml+CharsetSuffix, rr.Header().Get(HeaderContentType))
	assert.Equal(t, []byte("<html><body>test this</body></html>"), rr.Body.Bytes())
}

func TestOK(t *testing.T) {
	rr := httptest.NewRecorder()

	e, err := json.Marshal("test this")
	assert.NoError(t, err)
	OK(rr, httptest.NewRequest(http.MethodGet, "/", nil), "test this")
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, string(e), rr.Body.String())
}

func TestPartialContent(t *testing.T) {
	rr := httptest.NewRecorder()

	e, err := json.Marshal("test this")
	assert.NoError(t, err)
	PartialContent(rr, httptest.NewRequest(http.MethodGet, "/", nil), "test this")
	assert.Equal(t, http.StatusPartialContent, rr.Code)
	assert.JSONEq(t, string(e), rr.Body.String())
}

func TestRender(t *testing.T) {
	rqj := httptest.NewRequest(http.MethodGet, "/", nil)
	rqj.Header.Set(HeaderAccept, ContentTypeJson)

	rrj := httptest.NewRecorder()

	Render(rrj, rqj, http.StatusOK, nil)
	assert.Equal(t, ContentTypeJson+CharsetSuffix, rrj.Header().Get(HeaderContentType))

	rqx := httptest.NewRequest(http.MethodGet, "/", nil)
	rqx.Header.Set(HeaderAccept, ContentTypeXml)

	rrx := httptest.NewRecorder()

	Render(rrx, rqx, http.StatusOK, nil)
	assert.Equal(t, ContentTypeJson+CharsetSuffix, rrj.Header().Get(HeaderContentType))
}

func TestXML(t *testing.T) {
	// todo - implement
}
