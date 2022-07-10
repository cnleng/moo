package http

import (
	"net/http"

	"github.com/moobu/moo/gateway"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (h handler) String() string {
	return "http"
}

func New() gateway.Handler {
	return &handler{}
}
