package modules

import (
	"github.com/s-marashi/cotube/internal/modules/meeting"
	meetingRepo "github.com/s-marashi/cotube/internal/modules/meeting/repository"
	meetingSvc "github.com/s-marashi/cotube/internal/modules/meeting/service"
	"github.com/s-marashi/cotube/internal/modules/user"
	userRepo "github.com/s-marashi/cotube/internal/modules/user/repository"
	userSvc "github.com/s-marashi/cotube/internal/modules/user/service"
)

func InitMeetingService() meeting.MeetingService {
	repo := meetingRepo.NewInMemoryMeetingRepository()
	return meetingSvc.NewMeetingService(repo)
}

func InitUserService() user.UserService {
	repo := userRepo.NewInMemUserRepository()
	return userSvc.NewUserService(repo)
}
