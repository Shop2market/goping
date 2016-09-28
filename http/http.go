package http

import (
	"encoding/json"
	"net/http"

	"github.com/Shop2market/goping/domain"
	"github.com/julienschmidt/httprouter"
)

func Ping(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	response := json.NewEncoder(w)
	pingResult := domain.Ping(map[string]interface{}{})

	w.WriteHeader(200)
	response.Encode(pingResult)
}
