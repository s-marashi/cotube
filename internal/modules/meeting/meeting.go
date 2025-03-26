package meeting

import "github.com/s-marashi/cotube/internal/domain"

// MeetingRepository defines the methods for meeting storage
type MeetingRepository interface {
	Create(meeting *domain.Meeting) error
	FindByID(id string) (*domain.Meeting, error)
	AddUserToMeeting(meetingID string, userID string) error
	AddMessage(meetingID string, message *domain.Message) error
}

// MeetingService defines the business logic for meetings
type MeetingService interface {
	CreateMeeting(creatorID string) (*domain.Meeting, error)
	JoinMeeting(meetingID string, userID string) error
	SendMessage(meetingID string, content string, senderID string) error
	GetMeeting(meetingID string) (*domain.Meeting, error)
	GetMessages(meetingID string, lastMessageID string) ([]*domain.Message, error)
}
