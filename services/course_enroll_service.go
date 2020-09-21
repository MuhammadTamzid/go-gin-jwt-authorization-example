package services

import (
	"log"
	"go-gin-jwt-authorization-example/configs"
	"go-gin-jwt-authorization-example/dtos/dto_request"
	"go-gin-jwt-authorization-example/models"
)

func CourseEnroll(courseEnrollDto dto_request.CourseEnrollDto, userID int) (err error) {
	courseEnroll := models.CourseEnroll{CourseID: courseEnrollDto.CourseID, UserID: userID}

	if err = configs.DB.Create(&courseEnroll).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindCourseEnrolls(courseEnrolls *[]models.CourseEnroll, userID int) (err error) {
	if err = configs.DB.Preload("Course").Where("user_id = ?", userID).Find(courseEnrolls).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindCourseEnrollByEnrollIDAndUserID(courseEnroll *models.CourseEnroll, enrollID int, userID int) (err error) {
	if err = configs.DB.Preload("Course.CourseContents").Where("id = ? AND user_id = ?", enrollID, userID).Find(courseEnroll).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindCourseEnrollByCourseIDAndUserID(courseEnroll *models.CourseEnroll, courseID uint, userID int) (err error) {
	if err = configs.DB.Where("course_id = ? AND user_id = ?", courseID, userID).Find(courseEnroll).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
