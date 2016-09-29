package goping

import (
	"github.com/Shop2market/goping/http"
	"github.com/julienschmidt/httprouter"
)

func Ping() httprouter.Handle {
	return http.Ping
}
