package models

import (
	"beauty_store/api/database"
)

type LikeReview struct {
	ID_REVIEW int `json:"id_review"`
	ID_MEMBER int `json:"id_member"`
}

func InsertLikeReview(like LikeReview) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close() // Memperbaiki kesalahan di sini

	_, err = db.Exec("INSERT INTO like_review (ID_REVIEW, ID_MEMBER) VALUES (?, ?)", like.ID_REVIEW, like.ID_MEMBER)

	return err
}

func DeleteLikeReview(like LikeReview) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM like_review WHERE ID_REVIEW = ? AND ID_MEMBER = ?", like.ID_REVIEW, like.ID_MEMBER)

	return err
}
