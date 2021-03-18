package interfaces

import "net/http"

type Bolk interface {
	GetCounter(w http.ResponseWriter, r *http.Request)
	UpdateCounter(w http.ResponseWriter, r *http.Request)
}
