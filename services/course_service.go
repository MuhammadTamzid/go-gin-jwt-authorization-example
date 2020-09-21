package services

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"go-gin-jwt-authorization-example/configs"
	"go-gin-jwt-authorization-example/dtos/dto_request"
	"go-gin-jwt-authorization-example/models"
)

func CreateCourse(courseDto dto_request.CourseRequestDto) (err error) {
	var courseContents []models.CourseContent
	for _, k := range courseDto.CourseContents {
		courseContents = append(courseContents, models.CourseContent{Title: k.Title, Description: k.Description, LectureUrl: k.LectureUrl})
	}

	course := models.Course{Title: courseDto.Title, Description: courseDto.Description, CourseContents: courseContents}
	if err = configs.DB.Create(&course).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindCourses(courses *[]models.Course) (err error) {
	if err = configs.DB.Find(courses).Error; err != nil {
		log.Println(err)
		return err
	} 
	return nil
}

func FindCourseByID(course *models.Course, id int) (err error)  {
	if err = configs.DB.Where("id = ?", id).First(course).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteCourse(course *models.Course, id int) (err error)  {
	if err = configs.DB.Where("id = ?", id).Delete(course).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil;
}
