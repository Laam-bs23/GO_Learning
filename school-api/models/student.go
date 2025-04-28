package models

type Student struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	StudentName string `gorm:"not null" json:"student_name"`
	ClassId     int    `gorm:"not null" json:"class_id"`
	Secsion     string `gorm:"null" json:"student_section"`
}
