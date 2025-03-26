package handlers

import (
	"colab-tube/internal/modules/meeting"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MeetingHandler struct {
	meetingService meeting.MeetingService
}

func NewMeetingHandler(service meeting.MeetingService) *MeetingHandler {
	return &MeetingHandler{
		meetingService: service,
	}
}

type SendMessageRequest struct {
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

func (h *MeetingHandler) SendMessage(c echo.Context) error {
	meetingID := c.Param("meetingId")

	var req SendMessageRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	err := h.meetingService.SendMessage(meetingID, req.Content, req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "message sent"})
}

func (h *MeetingHandler) PollMessages(c echo.Context) error {
	meetingID := c.Param("meetingId")
	lastMessageID := c.QueryParam("lastMessageId")

	messages, err := h.meetingService.GetMessages(meetingID, lastMessageID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, messages)
}

func (h *MeetingHandler) CreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "user created"})
}
