package service

import (
	"github.com/s-marashi/cotube/internal/domain"
	"github.com/s-marashi/cotube/internal/modules/meeting"

	"github.com/google/uuid"
)

type meetingService struct {
	repository meeting.MeetingRepository
}

func NewMeetingService(repo meeting.MeetingRepository) meeting.MeetingService {
	return &meetingService{
		repository: repo,
	}
}

func (s *meetingService) CreateMeeting(creatorID string) (*domain.Meeting, error) {
	return nil, nil
}

func (s *meetingService) SendMessage(meetingID string, content string, senderID string) error {
	_, err := s.repository.FindByID(meetingID)
	if err != nil {
		return err
	}

	message := domain.NewMessage(
		uuid.New().String(),
		content,
		senderID,
	)

	return s.repository.AddMessage(meetingID, message)
}

func (s *meetingService) GetMessages(meetingID string, lastMessageID string) ([]*domain.Message, error) {
	meeting, err := s.repository.FindByID(meetingID)
	if err != nil {
		return nil, err
	}

	// If no lastMessageID provided, return all messages
	if lastMessageID == "" {
		return meeting.Messages, nil
	}

	// Find messages after the last received message
	var newMessages []*domain.Message
	found := false
	for _, msg := range meeting.Messages {
		if found {
			newMessages = append(newMessages, msg)
		}
		if msg.ID == lastMessageID {
			found = true
		}
	}

	return newMessages, nil
}

func (s *meetingService) GetMeeting(meetingID string) (*domain.Meeting, error) {
	return s.repository.FindByID(meetingID)
}

// JoinMeeting implements ports.MeetingService.
func (s *meetingService) JoinMeeting(meetingID string, userID string) error {
	panic("unimplemented")
}
