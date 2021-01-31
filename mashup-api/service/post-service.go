package service

import (
	"fmt"
	"net/http"
)

const (
	postServiceURL = "https://jsonplaceholder.typicode.com/posts/1"
)

type PostService interface {
	FetchData()
}

type fetchPostDataService struct{}

func NewPostService() PostService {
	return &fetchPostDataService{}
}

func (*fetchPostDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s\n", postServiceURL)

	// Call the external API
	resp, _ := client.Get(postServiceURL)

	// Write response to the channel
	postDataChannel <- resp
}
