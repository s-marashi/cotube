package main

import (
	"log"

	"github.com/s-marashi/cotube/internal/modules/meeting/handlers"
	"github.com/s-marashi/cotube/internal/modules/meeting/repository"
	"github.com/s-marashi/cotube/internal/modules/meeting/service"

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
