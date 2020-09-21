package models

import "github.com/jinzhu/gorm"

type CourseEnroll struct {
	gorm.Model
	CourseID int `join:"course_id"`
	UserID   int `join:"user_id"`
	Course 	 Course
}
