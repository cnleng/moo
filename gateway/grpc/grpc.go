package grpc

import (
	"net/http"

	"github.com/moobu/moo/gateway"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (h handler) String() string {
	return "grpc"
}

func New() gateway.Proxy {
	return &handler{}
}