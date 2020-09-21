package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"go-gin-jwt-authorization-example/dtos/dto_request"
	"go-gin-jwt-authorization-example/models"
	"go-gin-jwt-authorization-example/pkg/app"
	"go-gin-jwt-authorization-example/pkg/app/app_response"
	"go-gin-jwt-authorization-example/pkg/utils"
	"go-gin-jwt-authorization-example/services"
)

// CreateCourse godoc
// @Summary Create Course API
// @Tags courses
// @Accept  json
// @Produce  json
// @Param courseDto body dto_request.CourseRequestDto true "courseDto"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /courses [post]
func CreateCourse(c *gin.Context) {
	appGin := app.AppGin{C : c}
	var courseDto dto_request.CourseRequestDto
	if err := c.ShouldBindJSON(&courseDto); err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_DATA, err.Error())
		return
	}

	if err := services.CreateCourse(courseDto); err != nil {
		appGin.Response(http.StatusInternalServerError, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, nil)
}

// FindCourses godoc
// @Summary Find Courses API
// @Tags courses
// @Accept  json
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /courses [get]
func FindCourses(c *gin.Context) {
	appGin := app.AppGin{C : c}
	var courses []models.Course
	if err := services.FindCourses(&courses); err != nil {
		appGin.Response(http.StatusNotFound, app_response.INTERNAL_SERVER_ERROR, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, courses)
}

// FindCourseByID godoc
// @Summary Find Course By CourseID API
// @Tags courses
// @Accept  json
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /courses/{courseId} [get]
func FindCourseByID(c *gin.Context) {
	appGin := app.AppGin{C : c}
	id, err := utils.CastStringToInt(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_DATA, err.Error())
		return
	}

	var course models.Course
	if err := services.FindCourseByID(&course, id); err != nil {
		appGin.Response(http.StatusNotFound, app_response.COURSE_NOT_FOUND, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, course)
}

// DeleteCourse godoc
// @Summary Delete Course API
// @Tags courses
// @Accept  json
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /courses/{courseId} [delete]
func DeleteCourse(c *gin.Context)  {
	appGin := app.AppGin{C : c}
	id, err := utils.CastStringToInt(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		appGin.Response(http.StatusBadRequest, app_response.INVALID_DATA, err.Error())
		return
	}

	var course models.Course
	if err := services.DeleteCourse(&course, id); err !=nil {
		appGin.Response(http.StatusNotFound, app_response.COURSE_NOT_FOUND, nil)
		return
	}

	appGin.Response(http.StatusOK, app_response.SUCCESS, nil)
}
