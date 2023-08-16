package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	w := httptest.NewRecorder()
	Handler(w, req)
	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	t.Log(string(body))
}
