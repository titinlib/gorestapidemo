package database

import (
	"fmt"

	"restapidemo/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

var DB_SAMPLE_CONFIG = map[string]string{
	"DB_USER":          "root",
	"DB_PASSWORD":      "root",
	"DB_DATABASE_NAME": "test_db",
}

func init() {
	viper.AddConfigPath("../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// fmt.Println(
	// 	viper.GetString("DB_USER"),
	// 	viper.GetString("DB_PASSWORD"),
	// 	viper.GetString("DB_DATABASE_NAME"),
	// )

	// connectToDB(
	// 	viper.GetString("DB_USER"),
	// 	viper.GetString("DB_PASSWORD"),
	// 	viper.GetString("DB_DATABASE_NAME"),
	// )

	connectToDB(
		DB_SAMPLE_CONFIG["DB_USER"],
		DB_SAMPLE_CONFIG["DB_PASSWORD"],
		DB_SAMPLE_CONFIG["DB_DATABASE_NAME"],
	)
}

// connecting to database
func connectToDB(DB_USER, DB_PASSWORD, DB_DATABASE_NAME string) {

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_DATABASE_NAME)

	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err: ", err)
		panic("failed to connect database")
	}

	// Migrate the schemas
	migrate()
}

func migrate() {
	// TODO:
	database.AutoMigrate(&models.Product{}, &models.User{})
}

func GetDB() *gorm.DB {
	return database
}
