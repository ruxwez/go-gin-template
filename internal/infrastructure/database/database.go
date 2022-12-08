package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func Connection() error {
	var err error

	var db_host = os.Getenv("DB_HOST")
	var db_password = os.Getenv("DB_PASS")
	var db_user = os.Getenv("DB_USER")
	var db_name = os.Getenv("DB_NAME")

	dsnMySQL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_name)
	DB, err = gorm.Open(mysql.Open(dsnMySQL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		// CREAR FICHERO DE REGISTRO DE ERRORES
		return err
	}

	err = DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate()

	return err
}
