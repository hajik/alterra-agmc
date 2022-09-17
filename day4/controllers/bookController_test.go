package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetBooksDetail(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/books/1", nil)
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	e := echo.New()
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)

	return rec
}
