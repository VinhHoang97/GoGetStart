package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

// UpdatePost update a  spesific post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&post)

	query, err := db.Prepare("Update posts set title=?, content=? where id=?")
	catch(err)
	_, er := query.Exec(post.Title, post.Content, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

// DeletePost remove a specific post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete from posts where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

// AllPost remove a specific post
func AllPosts(w http.ResponseWriter, r *http.Request) {

	query, err := db.Prepare("Select * from posts")
	catch(err)
	_, er := query.Exec()
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}


// DetailPost remove a spesific post
func DetailPost(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	query, err := db.Prepare("Select * from posts where  id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

