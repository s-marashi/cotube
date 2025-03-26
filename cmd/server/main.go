package main

import (
	"log"

	"colab-tube/internal/modules/meeting/handlers"
	"colab-tube/internal/modules/meeting/repository"
	"colab-tube/internal/modules/meeting/service"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Initialize repositories
	meetingRepo := repository.NewInMemoryMeetingRepository()

	// Initialize services
	meetingService := service.NewMeetingService(meetingRepo)

	// Initialize handlers
	meetingHandler := handlers.NewMeetingHandler(meetingService)

	// Routes
	//e.POST("/users/", )
	e.POST("/meetings/:meetingId/messages", meetingHandler.SendMessage)
	e.GET("/meetings/:meetingId/messages", meetingHandler.PollMessages)

	// Start server
	log.Fatal(e.Start(":8080"))
}
