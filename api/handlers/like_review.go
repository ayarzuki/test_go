package handlers

import (
	"beauty_store/api/models"
	"encoding/json"
	"net/http"
)

func InsertLikeReview(w http.ResponseWriter, r *http.Request) {
	var like models.LikeReview
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.InsertLikeReview(like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteLikeReview(w http.ResponseWriter, r *http.Request) {
	var like models.LikeReview
	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeleteLikeReview(like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
