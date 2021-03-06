package scenario

import "time"

// Load
const (
	// initialStudentsCount 初期学生数
	initialStudentsCount = 50
	// initialCourseCount 初期科目数 30以上である必要がある
	initialCourseCount = 30
	// registerCourseLimitPerStudent は学生あたりの同時履修可能科目数の制限
	registerCourseLimitPerStudent = 20
	// StudentCapacityPerCourse は科目あたりの履修定員
	StudentCapacityPerCourse = 50
	// searchCountPerRegistration は履修登録前に実行する科目詳細取得の回数
	searchCountPerRegistration = 3
	// ClassCountPerCourse は科目あたりの講義数 -> same const exist in model/course.go
	ClassCountPerCourse = 5
	// AnnouncePagingStudentInterval はお知らせページングシナリオを開始する人数間隔
	AnnouncePagingStudentInterval = 10
	// waitCourseFullTimeout は最初の履修成功時刻から科目が満員に達するまでの待ち時間
	waitCourseFullTimeout = 2 * time.Second
	// waitReadClassAnnouncementTimeout は学生が講義課題のお知らせを確認するのを待つ最大時間
	waitReadClassAnnouncementTimeout = 5 * time.Second
	// waitGradeTimeout は成績取得がタイムアウトした際に再度確認しに行くまでの待ち
	waitGradeTimeout = 10 * time.Second
	// loadRequestTime はLoadシナリオ内でリクエストを送り続ける時間(Load自体のTimeoutより早めに終わらせる)
	loadRequestTime = 60 * time.Second
	// loginRetryCount はloadでloginがタイムアウトで失敗したときのリトライ回数
	loginRetryCount = 1
)

// Verify
const (
	assignmentsVerifyRate = 0.2
)

// Validation
const (
	validateGPAErrorTolerance              = 0.01
	validateTotalScoreErrorTolerance       = 0.01
	validateAnnouncementSampleStudentCount = 10
)
