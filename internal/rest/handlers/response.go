package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

type ResponseType interface {
	Data()(interface{})
}

type errorResponse struct {
	data interface{}
}

func (e errorResponse) Data() interface{}{
	return e.data
}

type dataResponse struct {
	data interface{}
}

func (e dataResponse) Data() interface{}{
	return e.data
}

// Response returns data for response
func Response[T ResponseType](a T, w http.ResponseWriter, r *http.Request) error {
	data := a.Data()
	render.JSON(w, r, data)
	return a.Run()
}