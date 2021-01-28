package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/short/XVIBz", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	testHandler := http.HandlerFunc(RedirectURL)

	testHandler.ServeHTTP(rr, req)

	fmt.Println(rr.Code)
	assert.Equal(t, 301, rr.Code)
}
