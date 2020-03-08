package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
  Id uint `json:"id" gorm:"primary_key"`
  Name string `json:"name"`
}

var DB *gorm.DB
func GetDB() *gorm.DB {
  if DB == nil {
    db, err := gorm.Open("mysql", "root@/fiber-test")
    if err != nil {
      panic(err.Error())
    }
    DB = db
  }

  DB.AutoMigrate(&User{})

  return DB
}

