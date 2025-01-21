package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"hr-system/internal/attendance"
	"hr-system/internal/config"
	"hr-system/internal/database"
	"hr-system/internal/employee"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// TODO: use DI
	db := database.InitDB()
	cache := database.InitRedis()

	employeeRepo := employee.NewRepository(db)
	attendanceRepo := attendance.NewRepository(db)

	employeeService := employee.NewService(employeeRepo, cache, 7*24*time.Hour)
	attendanceService := attendance.NewService(attendanceRepo)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	employee.RegisterRoutes(v1, employeeService)
	attendance.RegisterRoutes(v1, attendanceService)

	fmt.Printf("Starting server at :%s\n", config.Get().Server.Port)
	log.Fatal(router.Run(":" + config.Get().Server.Port))

	// TODO: handle graceful shutdown
}
