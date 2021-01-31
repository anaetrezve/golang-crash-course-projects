package controller

import (
	"encoding/json"
	"net/http"

	"github.com/anaetrezve/golang-crash-course-projects/service"
)

type controller struct {
	service service.PostDetailsService
}

type PostDetailsController interface {
	GetPostDetails(w http.ResponseWriter, r *http.Request)
}

func NewPostDetailsController(service service.PostDetailsService) PostDetailsController {
	return &controller{service}
}

func (c *controller) GetPostDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	postDetails := c.service.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postDetails)
}
