package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	app := NewServer()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// レスポンスボディの内容を読み取る
	body := new(strings.Builder)
	_, _ = io.Copy(body, resp.Body)
	assert.Equal(t, "Hello, World!", body.String())
}

func TestErrorRoute(t *testing.T) {
	app := NewServer()

	req := httptest.NewRequest(http.MethodGet, "/error", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// レスポンスボディの内容を読み取る
	body := new(strings.Builder)
	_, _ = io.Copy(body, resp.Body)
	assert.JSONEq(t, `{"error": "Bad Request Example"}`, body.String())
}

func TestNotFoundRoute(t *testing.T) {
	app := NewServer()

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	// レスポンスボディの内容を読み取る
	body := new(strings.Builder)
	_, _ = io.Copy(body, resp.Body)
	assert.JSONEq(t, `{"error": "Cannot GET /notfound"}`, body.String())
}
