package main

import (
	"encoding/json"
	"net/http"
)

// CreatePost create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Insert posts SET title=?, content=?")
	catch(err)

	_, er := query.Exec(post.Title, post.Content)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}