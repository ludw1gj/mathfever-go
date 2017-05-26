package service

import (
	"net/http"
)

type Service interface {
	Execute() (string, error)
	HandleAPI(http.ResponseWriter, *http.Request)
}
