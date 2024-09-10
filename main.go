package main

import (
	"github.com/gin-gonic/gin"
	"backend-onboard-task/config"
	"backend-onboard-task/models"
	"backend-onboard-task/routes"
)

func main() {
	// 데이터베이스 설정 및 연결
	db := config.ConnectDatabase()

	// 모델 기반으로 테이블을 자동 생성 및 업데이트
	db.AutoMigrate(
		&models.User{}, 
		&models.Board{},
	) 

	// 기본 설정으로 Gin 인스턴스 생성
    r := gin.Default()

	// 라우팅 경로 설정
    routes.SetupRoutes(r, db)

	// HTTP 서버 실행 (포트 8080)
	r.Run()
}