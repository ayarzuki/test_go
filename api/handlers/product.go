package handlers

import (
	"beauty_store/api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func GetProductReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idProduct, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reviews, err := models.GetProductReviews(idProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reviews)
}
