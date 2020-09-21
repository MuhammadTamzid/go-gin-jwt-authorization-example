package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func InitPersistence() error {
	var err error
	DB, err = gorm.Open("mysql", getDbURL())
	if err != nil {
		log.Print(err)
		return err
	}

	//TODO: Handle migrations
	//DB.AutoMigrate( &models.Role{}, &models.User{}, &models.Course{}, &models.CourseContent{}, &models.CourseEnroll{})
	return nil
}

func getDbURL() string {
	databaseEnv := EnvVariables.Database
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseEnv.User, databaseEnv.Password, databaseEnv.Host, databaseEnv.Port, databaseEnv.DBName,
	)
}
