package handlers

import (
	"net/http"
	"encoding/json"
	models "github.com/saromanov/kigo/internal/models/rest"
)

type createService struct {

}
func Create() http.Handler {
	return &createService{}
}

func (c *createService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var t models.ServiceCreate
    err := decoder.Decode(&t)
    if err != nil {
		e := errorResponse{
			data: models.ErrorResponse {
				Message: "unable to decode response",
			},
		}
		Response[errorResponse](e, w, r)
		return
	}
	d := dataResponse{
		data: t,
	}
	Response[dataResponse](d, w, r)
	
}

