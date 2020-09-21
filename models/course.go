package models

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Title string `join:"title"`
	Description string `join:"description"`
	CourseContents []CourseContent
}
