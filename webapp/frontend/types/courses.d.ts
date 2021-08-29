export type DayOfWeek =
  | 'monday'
  | 'tuesday'
  | 'wednesday'
  | 'thursday'
  | 'friday'

export type CourseType = 'liberal-arts' | 'major-subjects'

export type CourseStatus = 'registration' | 'in-progress' | 'closed'

export type User = {
  code: string
  name: string
  isAdmin: boolean
}

export type SearchCourseRequest = {
  keywords?: string
  type?: string
  credit?: number
  teacher?: string
  period?: number
  dayOfWeek?: string
  page?: number
}

export type Course = {
  id: string
  code: string
  type: CourseType
  name: string
  description: string
  credit: number
  period: number
  dayOfWeek: DayOfWeek
  teacher: string
  keywords: string
}

export type AddCourseRequest = Omit<Course, 'id' | 'teacher'>

export type SetCourseStatusRequest = {
  status: CourseStatus
}

export type Announcement = {
  id: string
  courseId: string
  courseName: string
  title: string
  unread: boolean
  createdAt: string
  message?: string
}

export type AnnouncementResponse = {
  id: string
  courseId: string
  courseName: string
  title: string
  unread: boolean
  createdAt: number
}

export type GetAnnouncementResponse = {
  unreadCount: number
  announcements: AnnouncementResponse[]
}

export type ClassInfo = {
  id: string
  part: number
  title: string
  description: string
  submissionClosed: boolean
  submitted: boolean
}

export type AddClassRequest = {
  part: number
  title: string
  description: string
  createdAt: number
}

type RegisterScoreRequestObject = {
  userCode: string
  score: number
}

export type RegisterScoreRequest = RegisterScoreRequestObject[]

type SummaryGrade = {
  gpa: number
  credits: number
  gpaAvg: number
  gpaTScore: number
  gpaMax: number
  gpaMin: number
}

type ClassScore = {
  title: string
  part: number
  score: number
  submitters: number
}

type CourseGrade = {
  name: string
  code: string
  totalScore: number
  totalScoreAvg: number
  totalScoreTScore: number
  totalScoreMax: number
  totalScoreMin: number
  classScores: ClassScore[]
}

export type Grade = {
  summary: SummaryGrade
  courses: CourseGrade[]
}
