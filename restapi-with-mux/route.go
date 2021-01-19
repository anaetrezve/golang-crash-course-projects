package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
	Text  string `json: "text"`
}

var posts []Post

func init() {
	posts = []Post{{ID: 1, Title: "Title One", Text: "Text One"}}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshaling the posts array"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshaling the post body"}`))
		return
	}
	post.ID = len(posts) + 1
	posts = append(posts, post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshaling the post"}`))
		return
	}
	w.Write(result)
}
