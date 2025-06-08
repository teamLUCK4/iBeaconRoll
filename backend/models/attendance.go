package models

// ##--- frontend에서 받는 데이터 형식 ---##
type AttendanceUpdateRequest struct {
	StudentID      int    `json:"student_id"`
	TimetableID    int    `json:"timetable_id"`
	Status         string `json:"status"`
	Classroom      string `json:"classroom"`
	AttendanceDate string `json:"attendance_date"`
}
