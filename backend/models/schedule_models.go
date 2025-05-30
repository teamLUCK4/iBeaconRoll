package models

import (
	"database/sql"
)

// Timetable 구조체는 시간표 정보를 나타냅니다.
type Timetable struct {
	ID          int    `json:"id" db:"id"`
	StudentID   int    `json:"student_id" db:"student_id"`
	Semester    int    `json:"semester" db:"semester"`
	SubjectName string `json:"subject_name" db:"subject_name"`
	DayOfWeek   string `json:"day_of_week" db:"day_of_week"`
	StartTime   string `json:"start_time" db:"start_time"`
	EndTime     string `json:"end_time" db:"end_time"`
	Classroom   string `json:"classroom" db:"classroom"`
}

// TimetableWithAttendance 구조체는 출석 정보가 포함된 시간표를 나타냅니다.
type TimetableWithAttendance struct {
	ID               int            `json:"id" db:"id"`
	StudentID        int            `json:"student_id" db:"student_id"`
	Semester         int            `json:"semester" db:"semester"`
	SubjectName      string         `json:"subject_name" db:"subject_name"`
	DayOfWeek        string         `json:"day_of_week" db:"day_of_week"`
	StartTime        string         `json:"start_time" db:"start_time"` // "10:00:00" 형식
	EndTime          string         `json:"end_time" db:"end_time"`     // "11:30:00" 형식
	Classroom        string         `json:"classroom" db:"classroom"`
	Status           sql.NullString `json:"status,omitempty" db:"status"`
	AttendanceTime   sql.NullString `json:"attendance_time,omitempty" db:"attendance_time"`
	AttendanceStatus string         `json:"attendance_status" db:"attendance_status"`
}

// DailySchedule 구조체는 특정 날짜의 일정 목록을 나타냅니다.
type DailySchedule struct {
	Date      string                    `json:"date"` // ISO8601 format: "2025-05-30T07:09:12Z"
	StudentID int                       `json:"student_id"`
	DayOfWeek string                    `json:"day_of_week"`
	Classes   []TimetableWithAttendance `json:"classes"`
	UpdatedAt string                    `json:"updated_at"` // ISO8601 format: "2025-05-30T07:09:12Z"
}
