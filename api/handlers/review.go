package handlers

import (
	"beauty_store/api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InsertReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.InsertReview(review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.UpdateReview(review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idReview, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = models.DeleteReview(idReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
