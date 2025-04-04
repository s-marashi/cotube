package repository

import (
	"errors"
	"sync"

	"github.com/s-marashi/cotube/internal/domain"
	"github.com/s-marashi/cotube/internal/modules/meeting"
)

type inMemoryMeetingRepository struct {
	meetings map[string]*domain.Meeting
	mutex    sync.RWMutex
}

func NewInMemoryMeetingRepository() meeting.MeetingRepository {
	return &inMemoryMeetingRepository{
		meetings: make(map[string]*domain.Meeting),
	}
}

func (r *inMemoryMeetingRepository) Create(meeting *domain.Meeting) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.meetings[meeting.ID]; exists {
		return errors.New("meeting already exists")
	}

	r.meetings[meeting.ID] = meeting
	return nil
}

func (r *inMemoryMeetingRepository) FindByID(id string) (*domain.Meeting, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	meeting, exists := r.meetings[id]
	if !exists {
		return nil, errors.New("meeting not found")
	}

	return meeting, nil
}

func (r *inMemoryMeetingRepository) AddUserToMeeting(meetingID string, userID string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	meeting, exists := r.meetings[meetingID]
	if !exists {
		return errors.New("meeting not found")
	}

	for _, existingUserID := range meeting.Users {
		if existingUserID == userID {
			return errors.New("user already in meeting")
		}
	}

	meeting.Users = append(meeting.Users, userID)
	meeting.Status = domain.MeetStatusActive
	return nil
}

func (r *inMemoryMeetingRepository) AddMessage(meetingID string, message *domain.Message) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	meeting, exists := r.meetings[meetingID]
	if !exists {
		return errors.New("meeting not found")
	}

	meeting.Messages = append(meeting.Messages, message)
	return nil
}
