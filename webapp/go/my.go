package main

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var lock_1 sync.Mutex
var userCodeToUserId = map[string]string{}


func getUserIdFromCode(db *sqlx.DB, code string) string {
	lock_1.Lock()
	defer lock_1.Unlock()

	cachedId, ok := userCodeToUserId[code]
	if ok {
		return cachedId
	}

	var id string
	err := db.Get(&id, "SELECT `id` FROM `users` WHERE `code` = ?", code)
	if err != nil {
		panic("get user id from code error")
	}

	userCodeToUserId[code] = id

	return id
}
