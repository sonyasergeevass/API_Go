package models

import "gorm.io/gorm"

type Segment struct {
	gorm.Model
	SegmentName string `json:"segment_name" gorm:"text; not null;default:null"`
}

type User struct {
	gorm.Model
	UserId int64 `json:"user_id" gorm:"; not null;"`
}
