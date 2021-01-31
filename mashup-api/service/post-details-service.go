package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anaetrezve/golang-crash-course-projects/entity"
)

var (
	postService     PostService = NewPostService()
	userService     UserService = NewUserService()
	postDataChannel             = make(chan *http.Response)
	userDataChannel             = make(chan *http.Response)
)

type PostDetailsService interface {
	GetDetails() entity.PostDetails
}

type service struct{}

func NewPostDetailsService() PostDetailsService {
	return &service{}
}

func (*service) GetDetails() entity.PostDetails {
	// Have to implement
	// Goroutine to call endpoint 1
	go postService.FetchData()
	// Goroutine to call endpoint 2
	go userService.FetchData()

	// Create post channel to get the data
	post, _ := getPostData()
	// Create user channel to get the data
	user, _ := getUserData()

	return entity.PostDetails{
		ID:       post.ID,
		Title:    post.Title,
		Body:     post.Body,
		Username: user.Username,
		Email:    user.Email,
	}
}

func getPostData() (entity.Post, error) {
	readPost := <-postDataChannel
	var post entity.Post

	err := json.NewDecoder(readPost.Body).Decode(&post)
	if err != nil {
		fmt.Println(err.Error())
		return post, err
	}

	return post, nil
}

func getUserData() (entity.User, error) {
	readUser := <-userDataChannel
	var user entity.User

	err := json.NewDecoder(readUser.Body).Decode(&user)

	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	return user, nil
}
