package db

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/soonoo/committrs-server/models"
	// _. "github.com/volatiletech/sqlboiler/queries/qm"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

var DB *sql.DB

func GetDB() *sql.DB {
	if DB == nil {
		db, err := sql.Open("postgres", "dbname=boiler user=kakao_ent sslmode=disable")
		// boil.SetDB(db)
		boil.DebugMode = true
		if err != nil {
			panic(err.Error())
		}

		// db, err := gorm.Open("mysql", "root@/fiber-test")
		// if err != nil {
		//   panic(err.Error())
		// }
		DB = db
	}

	// DB.AutoMigrate(&models.User{}).
	// AutoMigrate(&models.Repo{}).
	// AutoMigrate(&models.Commit{})

	return DB
}
