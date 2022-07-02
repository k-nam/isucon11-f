package main

import "github.com/jmoiron/sqlx"

func getUserIdFromCode(db *sqlx.DB, code string) (string, error) {
	var id string
	err := db.Get(&id, "SELECT `id` FROM `users` WHERE `code` = ?", code)
	if err != nil {
		return "", err
	}
	return id, nil
}