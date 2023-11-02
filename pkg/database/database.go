package database

import (
	"employee-manager/internal/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDBFromViper() {
	driver := viper.GetString("database.driver")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")

	fmt.Println("============================")
	fmt.Println("DataBase Type :: ", driver)
	fmt.Println("============================")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, dbname)
	//dsn := "root:S@lim786$@tcp(mysql-container:3600)/employee-manager"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Auto-migrate the schema
	DB.AutoMigrate(&model.Employee{})
}
