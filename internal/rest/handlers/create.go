package handlers

import "net/http"

type createService struct {

}
func Create() http.Handler {
	return &createService{}
}

func (c *createService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
}

