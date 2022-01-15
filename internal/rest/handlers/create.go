package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/saromanov/kigo/internal/models"
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
		ErrorResponse(nil, w)
		return
	}
	Response(nil, w)
	
}

