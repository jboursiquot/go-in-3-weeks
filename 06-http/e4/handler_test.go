package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProverbsHandler(t *testing.T) {
	cases := []struct {
		scenario           string
		method             string
		endpoint           string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			scenario:           "ReturnsSpecificProverb",
			method:             http.MethodGet,
			endpoint:           "/proverbs/3",
			expectedStatusCode: http.StatusOK,
			expectedBody:       "Channels orchestrate; mutexes serialize.",
		},
		{
			scenario:           "ReturnsBadRequest",
			method:             http.MethodGet,
			endpoint:           "/proverbs/",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "",
		},
		{
			scenario:           "ReturnsMethodNotAllowed",
			method:             http.MethodPut,
			endpoint:           "/proverbs/3",
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedBody:       "",
		},
	}

	for _, c := range cases {
		r, _ := http.NewRequest(c.method, c.endpoint, nil)
		w := httptest.NewRecorder()
		h := newProverbsHandler()
		h.ServeHTTP(w, r)

		if w.Code != c.expectedStatusCode {
			t.Fatalf("expected %d but got %d for status code", c.expectedStatusCode, w.Code)
		}

		body := strings.TrimSpace(w.Body.String())
		if body != c.expectedBody {
			t.Fatalf("expected %s but got %s for body", c.expectedBody, body)
		}
	}
}
