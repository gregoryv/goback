package main

import (
	"net/http"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	go main()
	<-time.After(time.Millisecond) // bad test strategy

	cases := []struct {
		exp int // expected status code
		*http.Request
	}{
		{200, newRequest(t, "POST", "http://localhost:8080/room/1")},
		{200, newRequest(t, "GET", "http://localhost:8080/room")},
		{200, newRequest(t, "GET", "http://localhost:8080/")},
	}
	for _, c := range cases {
		resp, err := http.DefaultClient.Do(c.Request)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != c.exp {
			t.Error(resp.Status)
		}
	}
}

func newRequest(t *testing.T, method, path string) *http.Request {
	r, err := http.NewRequest(method, path, http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	return r
}
