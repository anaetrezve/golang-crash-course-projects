package service

import (
	"fmt"
	"net/http"
)

const (
	userServiceURL string = "https://jsonplaceholder.typicode.com/users/1"
)

type UserService interface {
	FetchData()
}

type fetchUserDataService struct{}

func NewUserService() UserService {
	return &fetchUserDataService{}
}

func (*fetchUserDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s\n", userServiceURL)

	// Call the external API
	resp, _ := client.Get(userServiceURL)

	// Write response to the channel
	userDataChannel <- resp
}
