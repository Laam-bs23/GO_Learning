package models

type Class struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ClassName   string    `gorm:"not null" json:"class_name"`
	StudentCount int      `json:"student_count"`
} 