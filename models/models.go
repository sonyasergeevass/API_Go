package models

import "gorm.io/gorm"

type Segment struct {
	gorm.Model
	SegmentName string `json:"segment_name" gorm:"text; not null"`
	Users       []User `gorm:"many2many:user_segments;"`
}

type User struct {
	gorm.Model
	UserId   int64     `json:"user_id" gorm:"not null; default: null"`
	Segments []Segment `gorm:"many2many:user_segments;"`
}

type UserSegment struct {
	UserID    uint `json:"user_id"`
	SegmentID uint `json:"segment_id"`
}
