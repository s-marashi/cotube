package main

import (
	"log"
	"log/slog"

	"github.com/s-marashi/cotube/internal/modules"
	meetingHdl "github.com/s-marashi/cotube/internal/modules/meeting/handlers"
	userHdl "github.com/s-marashi/cotube/internal/modules/user/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	// Initialize Echo
	e := echo.New()

	// Initialize handlers
	meetingHandler := meetingHdl.NewMeetingRestHandler(modules.InitMeetingService())
	userHandler := userHdl.NewUserRestHandler(modules.InitUserService())

	// Routes
	e.POST("/users", userHandler.RegisterUser)
	e.POST("/meetings/:meetingId/messages", meetingHandler.SendMessage)
	e.GET("/meetings/:meetingId/messages", meetingHandler.PollMessages)

	// Start server
	log.Fatal(e.Start(":8080"))
}
