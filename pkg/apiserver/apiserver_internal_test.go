package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handel_hello(t *testing.T) {
	s := New(&Config{bind: ":8000"})
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handle_hello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")

}
