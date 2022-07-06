package main

import (
	"fmt"
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
	// start := time.Now()
	err := db.Get(&id, "SELECT `id` FROM `users` WHERE `code` = ?", code)
	// fmt.Printf("get user id. code: %s,  took: %d\n", code, time.Now().Sub(start).Milliseconds())
	if err != nil {
		panic("get user id from code error")
	}

	userCodeToUserId[code] = id

	return id
}

var lock_2 sync.Mutex

var userGpaInfo = map[string]UserGpaInfo{}

// var gpaMin float64 = 0
// var gpaMax float64 = 0
// var gpaSum float64 = 0
// var gpaAgv float64 = 0
// var numUsers int = 0

type GpaInfo struct {
	tScore float64
	gpaAgv float64
	gpaMin float64
	gpaMax float64
}

type UserGpaInfo struct {
	UserId         string `db:"user_id"`
	CreditSum      int    `db:"credits"`
	CreditScoreSum int    `db:"score_sum"`
	Gpa            float64
}

func addScores(scores []UserGpaInfo, courseId string) {
	lock_2.Lock()
	defer lock_2.Unlock()

	// user id -> (credit, score)
	for _, score := range scores {
		info, ok := userGpaInfo[score.UserId]
		if !ok {
			info = UserGpaInfo{}
		}
		info.CreditSum += score.CreditSum
		info.CreditScoreSum += score.CreditScoreSum
		newGpa := float64(info.CreditScoreSum) / float64(info.CreditSum) / 100.0
		info.Gpa = newGpa
		fmt.Printf("new gpa %d, %d, %f, %s, %s\n", info.CreditSum, info.CreditScoreSum, newGpa, score.UserId, courseId)
		userGpaInfo[score.UserId] = info
	}
}

var gpaLoaded = false

func loadGpa(db *sqlx.DB) {
	// lock_2.Lock()
	// defer lock_2.Unlock()
	fmt.Printf("Load gpa\n")

	userCreditSums := []UserGpaInfo{}

	query := "SELECT `users`.`id` AS `user_id`, SUM(`courses`.`credit`) AS `credits`" +
		"     FROM `users`" +
		"     JOIN `registrations` ON `users`.`id` = `registrations`.`user_id`" +
		"     JOIN `courses` ON `registrations`.`course_id` = `courses`.`id` AND `courses`.`status` = ?" +
		"     WHERE `users`.`type` = ? GROUP BY `users`.`id` ORDER BY `users`.`id`"
	if err := db.Select(&userCreditSums, query, StatusClosed, Student); err != nil {
		panic(err)
	}
	fmt.Printf("len 1 %d\n", len(userCreditSums))
	userCreditSums2 := []UserGpaInfo{}
	query = "SELECT `users`.`id` AS `user_id`, IFNULL(SUM(`submissions`.`score` * `courses`.`credit`), 0) AS `score_sum`" +
		" FROM `users`" +
		" JOIN `registrations` ON `users`.`id` = `registrations`.`user_id`" +
		" JOIN `courses` ON `registrations`.`course_id` = `courses`.`id` AND `courses`.`status` = ?" +
		" LEFT JOIN `classes` ON `courses`.`id` = `classes`.`course_id`" +
		" LEFT JOIN `submissions` ON `users`.`id` = `submissions`.`user_id` AND `submissions`.`class_id` = `classes`.`id`" +
		" WHERE `users`.`type` = ?" +
		" GROUP BY `users`.`id` ORDER BY `users`.`id`"
	if err := db.Select(&userCreditSums2, query, StatusClosed, Student); err != nil {
		panic(err)
	}
	fmt.Printf("len 2 %d\n", len(userCreditSums2))

	for i := range userCreditSums {
		fmt.Printf("initial gpa %s, %d, %d\n", userCreditSums[i].UserId, userCreditSums[i].CreditSum, userCreditSums2[i].CreditScoreSum)
		userGpaInfo[userCreditSums[i].UserId] = UserGpaInfo{
			CreditSum:      userCreditSums[i].CreditSum,
			CreditScoreSum: userCreditSums2[i].CreditScoreSum,
			Gpa:            float64(userCreditSums2[i].CreditScoreSum) / float64(userCreditSums[i].CreditSum) / 100,
		}
	}
}

func getGpaInfo(db *sqlx.DB, myGpa float64) []float64 {
	lock_2.Lock()
	defer lock_2.Unlock()

	fmt.Printf("get gpa info\n")
	if !gpaLoaded {
		gpaLoaded = true
		loadGpa(db)
	}

	gpas := []float64{}
	for _, info := range userGpaInfo {
		gpas = append(gpas, info.Gpa)
	}

	return gpas
}

func getCourseScoreSums(db *sqlx.Tx, courseId string) []UserGpaInfo {
	userScoreSums := []UserGpaInfo{}
	query := "SELECT `users`.`id` AS `user_id`, IFNULL(SUM(`submissions`.`score` * `courses`.`credit`), 0) AS `score_sum`" +
		" FROM `users`" +
		" JOIN `registrations` ON `users`.`id` = `registrations`.`user_id`" +
		" JOIN `courses` ON `registrations`.`course_id` = `courses`.`id` AND `courses`.`status` = ?" +
		" LEFT JOIN `classes` ON `courses`.`id` = `classes`.`course_id`" +
		" LEFT JOIN `submissions` ON `users`.`id` = `submissions`.`user_id` AND `submissions`.`class_id` = `classes`.`id`" +
		" WHERE `users`.`type` = ? AND `courses`.`id` = ? " +
		" GROUP BY `users`.`id` ORDER BY `users`.`id`"
	if err := db.Select(&userScoreSums, query, StatusClosed, Student, courseId); err != nil {
		fmt.Println("getCourseScoreSums error 1")
		panic(err)
	}

	var credit int = 0
	query = "SELECT `courses`.`credit` FROM `courses` WHERE `courses`.`id` = ?"
	err := db.Get(&credit, query, courseId)
	if err != nil {
		fmt.Println("getCourseScoreSums error 2")
		panic(err)
	}

	for i := range userScoreSums {
		userScoreSums[i].CreditSum = credit
	}

	return userScoreSums
}

var submissionLock sync.Mutex

// {userId_classId} -> filename
var submissions = map[string]Submission{}

const bulkSize = 10

func addSubmission(db *sqlx.DB, sub Submission) {
	submissionLock.Lock()
	defer submissionLock.Unlock()
	key := fmt.Sprintf("%s_%s", sub.UserID, sub.ClassId)
	submissions[key] = sub

	// fmt.Printf("add submission: %s, %s\n", sub.UserID, sub.ClassId)
	// if len(submissions) >= bulkSize {
	// 	fmt.Printf("db bulk: %d\n", len(submissions))
	// 	arr := []Submission{}
	// 	for _, sub := range submissions {
	// 		arr = append(arr, sub)
	// 	}

	// 	db.NamedExec("INSERT INTO `submissions` (`user_id`, `class_id`, `file_name`) VALUES (:user_id, :class_id, :file_name) ON DUPLICATE KEY UPDATE `file_name` = VALUES(`file_name`)", arr)
	// }
}

// func bulkInsertSubs(db *sqlx.DB) {
// 	submissionLock.Lock()
// 	defer submissionLock.Unlock()

// 	arr := []Submission{}
// 	for _, sub := range submissions {
// 		arr = append(arr, sub)
// 	}

// 	db.NamedExec("INSERT INTO `submissions` (`user_id`, `class_id`, `file_name`) VALUES (:user_id, :class_id, :file_name) ON DUPLICATE KEY UPDATE `file_name` = VALUES(`file_name`)", arr)
// }

func bulkInsertSubsTx(tx *sqlx.Tx, classID string) {
	submissionLock.Lock()
	defer submissionLock.Unlock()

	arr := []Submission{}
	for key, sub := range submissions {
		if sub.ClassId == classID {
			arr = append(arr, sub)
			delete(submissions, key)
			// fmt.Printf("tx bulk: %s \n", classID)
		}
	}
	// fmt.Printf("tx bulk: %s , %d\n", classID, len(arr))
	_, err := tx.NamedExec("INSERT INTO `submissions` (`user_id`, `class_id`, `file_name`) VALUES (:user_id, :class_id, :file_name) ON DUPLICATE KEY UPDATE `file_name` = VALUES(`file_name`)", arr)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
