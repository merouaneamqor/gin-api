// internal/model/user.go
package model

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    gorm.Model
    Name      string    `gorm:"type:varchar(100);not null" validate:"required,alpha"`
    Email     string    `gorm:"type:varchar(100);unique;not null" validate:"required,email"`
    Password  string    `gorm:"not null" validate:"required,min=8"`
    Birthdate time.Time `validate:"required"`
    IsActive  bool      `gorm:"default:false"`

}
