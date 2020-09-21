package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name 	 string `json:"name"`
	Email 	 string `json:"email"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
	Role     Role
}
