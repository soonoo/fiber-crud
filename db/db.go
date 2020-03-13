package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/soonoo/committrs-server/models"
)

var DB *gorm.DB
func GetDB() *gorm.DB {
  if DB == nil {
    db, err := gorm.Open("mysql", "root@/fiber-test")
    db.LogMode(true)
    if err != nil {
      panic(err.Error())
    }
    DB = db
  }

  DB.AutoMigrate(&models.User{}).
    AutoMigrate(&models.Repo{}).
    AutoMigrate(&models.Commit{})

  return DB
}

