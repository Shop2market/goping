package http

import(
  "github.com/Shop2market/goping/domain"
)

func Ping(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	response := json.NewEncoder(w)
	pingResult := domain.Ping(map[string]interface{})

  w.WriteHeader(200)
	response.Encode(pingResult)
}
