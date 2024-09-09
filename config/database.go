package config

import (
	"fmt"
	"log"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"backend-onboard-task/helpers"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		helpers.GetEnv("MYSQL_ROOT_PASSWORD", "1234"),
		helpers.GetEnv("MYSQL_HOST", "localhost"),
		helpers.GetEnv("MYSQL_PORT", "3306"),
		helpers.GetEnv("MYSQL_DATABASE", "onycom"),
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // SQL 쿼리 로깅 활성화
	})
	if err != nil {
		log.Fatalf("데이터베이스 연결에 실패했습니다: %v", err)
	}

	log.Println("데이터베이스 연결에 성공했습니다.")

	return db
}