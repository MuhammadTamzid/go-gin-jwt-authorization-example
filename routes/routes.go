package routes

import (
	"log"
	"go-gin-jwt-authorization-example/controllers"
	"go-gin-jwt-authorization-example/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// TODO: load policy from DB. Ref: https://github.com/Blank-Xu/sql-adapter
	e, err := casbin.NewEnforcer("authz/model.conf", "authz/policy.csv")
	log.Print(err)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/users/register", controllers.Register)

		auth := v1.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/refresh-token", controllers.RefreshToken)
		}

		courses := v1.Group("/courses")
		{
			courses.GET("/", controllers.FindCourses)
			courses.GET("/:id", controllers.FindCourseByID)
		}
	}

	v1.Use(middleware.AuthMiddleware())
	v1.Use(middleware.Authorizer(e))
	{
		courses := v1.Group("/courses")
		{
			courses.POST("", controllers.CreateCourse)
			courses.DELETE("/:id", controllers.DeleteCourse)
		}

		courseEnrolls := v1.Group("/course-enrolls")
		{
			courseEnrolls.POST("", controllers.CourseEnroll)
			courseEnrolls.GET("", controllers.FindCourseEnrolls)
			courseEnrolls.GET("/:id", controllers.FindCourseEnrollByID)
		}

		v1.POST("/auth/logout", controllers.Logout)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
