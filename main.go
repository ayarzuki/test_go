package main

import (
	"beauty_store/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Member routes
	r.HandleFunc("/members", handlers.GetAllMembers).Methods("GET")
	r.HandleFunc("/members", handlers.InsertMember).Methods("POST")
	r.HandleFunc("/members", handlers.UpdateMember).Methods("PUT")
	r.HandleFunc("/members/{id}", handlers.DeleteMember).Methods("DELETE")

	// Product routes
	r.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}/reviews", handlers.GetProductReviews).Methods("GET")

	// LikeReview routes
	r.HandleFunc("/like_review", handlers.InsertLikeReview).Methods("POST")
	r.HandleFunc("/like_review", handlers.DeleteLikeReview).Methods("DELETE")

	// Review routes
	r.HandleFunc("/reviews", handlers.InsertReview).Methods("POST")
	r.HandleFunc("/reviews", handlers.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
