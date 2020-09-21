package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"go-gin-jwt-authorization-example/dtos/dto_request"
	"go-gin-jwt-authorization-example/middleware"
	"go-gin-jwt-authorization-example/models"
	"go-gin-jwt-authorization-example/pkg/app"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
	"go-gin-jwt-authorization-example/pkg/utils"
	"go-gin-jwt-authorization-example/services"
	"reflect"
)

// CourseEnroll godoc
// @Summary Course Enroll API
// @Tags course-enrolls
// @Accept  json
// @Produce  json
// @Param courseEnrollDto body dto_request.CourseEnrollDto true "courseEnrollDto"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /course-enrolls [post]
func CourseEnroll(c *gin.Context)  {
	appGin := app.AppGin{C : c}

	var courseEnrollDto dto_request.CourseEnrollDto
	if err := c.ShouldBindJSON(&courseEnrollDto); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var course models.Course
	if err := services.FindCourseByID(&course, courseEnrollDto.CourseID); err != nil {
		appGin.Response(http.StatusNotFound, app_response.COURSE_NOT_FOUND, nil)
		return
	}

	var courseEnroll models.CourseEnroll
	if err := services.FindCourseEnrollByCourseIDAndUserID(&courseEnroll, course.Model.ID, middleware.CurrentUserID); err == nil &&
			!reflect.ValueOf(courseEnroll).IsZero() {
		appGin.Response(http.StatusBadRequest, app_response.ALREADY_REGISTER_IN_THIS_COURSE, nil)
		return
	}

	if err := services.CourseEnroll(courseEnrollDto, middleware.CurrentUserID); err != nil {
		appGin.Response(http.StatusInternalServerError, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, nil)
}

// FindCourseEnrolls godoc
// @Summary Find Course Enrolls API
// @Tags course-enrolls
// @Accept  json
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /course-enrolls [get]
func FindCourseEnrolls(c *gin.Context) {
	appGin := app.AppGin{C : c}
	var courseEnrolls []models.CourseEnroll
	if err := services.FindCourseEnrolls(&courseEnrolls, middleware.CurrentUserID); err != nil {
		appGin.Response(http.StatusNotFound, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, courseEnrolls)
}

// FindCourseEnrollByID godoc
// @Summary Find Course By Find Course Enroll By ID API
// @Tags course-enrolls
// @Accept  json
// @Produce  json
// @Param id path int true "Course Enroll ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /course-enrolls/{courseEnrollId} [get]
func FindCourseEnrollByID(c *gin.Context)  {
	appGin := app.AppGin{C : c}
	enrollId, err := utils.CastStringToInt(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_DATA, err.Error())
		return
	}

	var courseEnroll models.CourseEnroll
	if err := services.FindCourseEnrollByEnrollIDAndUserID(&courseEnroll, enrollId, middleware.CurrentUserID); err != nil {
		appGin.Response(http.StatusNotFound, app_response.NOT_ENROLLED, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, courseEnroll)
}
