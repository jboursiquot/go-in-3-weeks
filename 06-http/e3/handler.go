package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type proverb struct {
	id    int
	value string
}

type proverbsHandler struct {
	proverbs []proverb
}

func (ph *proverbsHandler) lookup(id int) (*proverb, error) {
	for _, p := range ph.proverbs {
		if id == p.id {
			return &p, nil
		}
	}
	return nil, errUnknownProverb
}

func newProverbsHandler() *proverbsHandler {
	return &proverbsHandler{
		proverbs: []proverb{
			{id: 1, value: "Don't communicate by sharing memory, share memory by communicating."},
			{id: 2, value: "Concurrency is not parallelism."},
			{id: 3, value: "Channels orchestrate; mutexes serialize."},
			{id: 4, value: "The bigger the interface, the weaker the abstraction."},
			{id: 5, value: "Make the zero value useful."},
		},
	}
}

func (ph *proverbsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Path[len("/proverbs/"):])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := ph.lookup(id)
		if err == errUnknownProverb {
			http.Error(w, errUnknownProverb.Error(), http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, p.value)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

var errUnknownProverb = errors.New("Unknown Proverb")
