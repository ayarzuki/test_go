package models

import (
	"beauty_store/api/database"
)

type Member struct {
	ID_MEMBER int    `json:"id_member"`
	USERNAME  string `json:"username"`
	GENDER    string `json:"gender"`
	SKINTYPE  string `json:"skintype"`
	SKINCOLOR string `json:"skin_color"`
}

func GetAllMembers() ([]Member, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM member")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member

	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID_MEMBER, &member.USERNAME, &member.GENDER, &member.SKINTYPE, &member.SKINCOLOR)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

func InsertMember(member Member) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO member (USERNAME, GENDER, SKINTYPE, SKINCOLOR) VALUES (?, ?, ?, ?)", member.USERNAME, member.GENDER, member.SKINTYPE, member.SKINCOLOR)

	return err
}

func UpdateMember(member Member) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE member SET USERNAME = ?, GENDER = ?, SKINTYPE = ?, SKINCOLOR = ? WHERE ID_MEMBER = ?", member.USERNAME, member.GENDER, member.SKINTYPE, member.SKINCOLOR, member.ID_MEMBER)

	return err
}

func DeleteMember(id int) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM member WHERE ID_MEMBER = ?", id)

	return err
}
