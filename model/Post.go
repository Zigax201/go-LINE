package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	ID        string `gorm:"primaryKey;autoIncrement"`
	ImagePath string `validator:"required"`
	Caption   string `form:caption validator:"required"`
	Author    int    `validator:"required"`
}
