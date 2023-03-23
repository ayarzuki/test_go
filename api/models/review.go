package models

import (
	"beauty_store/api/database"
)

type Review struct {
	ID_REVIEW   int    `json:"id_review"`
	ID_PRODUCT  int    `json:"id_product"`
	ID_MEMBER   int    `json:"id_member"`
	USERNAME    string `json:"username"`
	GENDER      string `json:"gender"`
	SKINTYPE    string `json:"skintype"`
	SKINCOLOR   string `json:"skin_color"`
	DESC_REVIEW string `json:"desc_review"`
	LIKES_COUNT int    `json:"likes_count"`
}

func GetProductReviews(idProduct int) ([]Review, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
SELECT
  rp.ID_REVIEW,
  rp.ID_PRODUCT,
  rp.ID_MEMBER,
  m.USERNAME,
  m.GENDER,
  m.SKINTYPE,
  m.SKINCOLOR,
  rp.DESC_REVIEW,
  COUNT(lr.ID_MEMBER) as LIKES_COUNT
FROM
  review_product rp
JOIN
  member m ON rp.ID_MEMBER = m.ID_MEMBER
LEFT JOIN
  like_review lr ON rp.ID_REVIEW = lr.ID_REVIEW
WHERE
  rp.ID_PRODUCT = ?
GROUP BY
  rp.ID_REVIEW
	`, idProduct)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review

	for rows.Next() {
		var review Review
		err := rows.Scan(
			&review.ID_REVIEW,
			&review.ID_PRODUCT,
			&review.ID_MEMBER,
			&review.USERNAME,
			&review.GENDER,
			&review.SKINTYPE,
			&review.SKINCOLOR,
			&review.DESC_REVIEW,
			&review.LIKES_COUNT,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func InsertReview(review Review) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO review_product (ID_PRODUCT, ID_MEMBER, DESC_REVIEW) VALUES (?, ?, ?)", review.ID_PRODUCT, review.ID_MEMBER, review.DESC_REVIEW)

	return err
}

func UpdateReview(review Review) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE review_product SET ID_PRODUCT = ?, ID_MEMBER = ?, DESC_REVIEW = ? WHERE ID_REVIEW = ?", review.ID_PRODUCT, review.ID_MEMBER, review.DESC_REVIEW, review.ID_REVIEW)

	return err
}

func DeleteReview(idReview int) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM review_product WHERE ID_REVIEW = ?", idReview)

	return err
}
