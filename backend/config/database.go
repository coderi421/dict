package config

import (
	"dict/model"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

// GetDB - call this method to get db
func GetDB() *gorm.DB {
	return db
}

// SetupDB - setup dabase for sharing to all api
func init() {
	_ = godotenv.Load(".env")

	var err error

	databaseInisialisation := "" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	//database := db
	//Switch Database
	if os.Getenv("DB_DRIVER") == "MYSQL" {
		db, err = gorm.Open(mysql.Open(databaseInisialisation), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}
	} else if os.Getenv("DB_DRIVER") == "POSTGRES" {
		db, err = gorm.Open(postgres.Open(databaseInisialisation), &gorm.Config{})
	} else {
		panic("database driver not supported")
	}

	//err = database.AutoMigrate(&model.User{})

	//if err != nil {
	//	panic("connection to database failed")
	//}

	// If data exist, not run seeder
	//err = database.First(&model.User{}).Error

	//if err != nil {
	//
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		Seeds(database)
	//	}
	//}

}

func Seeds(db *gorm.DB) bool { //https://gorm.io/docs/migration.html
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.MinCost)

	if err != nil {
		panic(err.Error())
	}

	var users = []model.User{
		model.User{
			ID:       1,
			Username: "user01",
			//Email:    "user@example.com",
			Password: string(passwordHash),
			//Phone:    "08123456789",
		},
		model.User{
			ID:       2,
			Username: "user02",
			//Email:    "user2@example.com",
			Password: string(passwordHash),
			//Phone:    "08123456789",
		},
	}

	err = db.Migrator().DropTable(&model.User{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.User{})

	if err != nil {
		panic(err.Error())
	}

	for i, _ := range users {
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error

		if err != nil {
			panic("Migration Failed")
		}
	}

	return true

}
