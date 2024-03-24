package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_sayHello(t *testing.T) {
	// prepare response recorder and request
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", http.NoBody)

	// call handler
	handler := sayHello()
	handler(w, r)

	// check result
	resp := w.Result()
	if v := resp.StatusCode; v != 200 {
		t.Error(resp.Status)
	}
}
