package service

import "net/http"

type MathApi interface {
	Execute() (string, error)
	HandleAPI(http.ResponseWriter, *http.Request)
}
