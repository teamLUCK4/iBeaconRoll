package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ## ==== 프런트 디스크 메모리에 보낼 '오늘의 시간표 + UUID' 데이터 형식 ==== ##
type Beacon struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UUID      string             `bson:"uuid" json:"uuid"`
	Classroom string             `bson:"classroom" json:"classroom"`
	CreatedAt time.Time          `bson:"createdAt" json:"created_at"`
}

// TimetableWithBeacon은 시간표와 비콘 정보를 함께 포함하는 구조체입니다.
type TimetableWithBeacon struct {
	TimetableWithAttendance
	BeaconInfo *Beacon `json:"beacon_info,omitempty"`
}
