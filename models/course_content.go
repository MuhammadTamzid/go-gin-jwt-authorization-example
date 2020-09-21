package models

import "github.com/jinzhu/gorm"

type CourseContent struct {
	gorm.Model
	Title       string `join:"title"`
	Description string `join:"description"`
	LectureUrl  string `join:"lecture_url"`
	CourseID    int `join:"course_id"`
}
