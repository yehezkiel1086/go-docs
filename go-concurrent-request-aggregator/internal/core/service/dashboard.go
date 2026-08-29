package service

import (
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/domain"
	"github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/port"
)

type DashboardService struct {
	userRepo port.UserRepository
	postRepo port.PostRepository
}

func NewDashboardService(
	userRepo port.UserRepository,
	postRepo port.PostRepository,
) *DashboardService {
	return &DashboardService{userRepo, postRepo}
}

type userResult struct {
	user *domain.User
	err  error
}

type postsResult struct {
	posts []*domain.Post
	err   error
}

func (s *DashboardService) GetDashboard(userId string) (*domain.Dashboard, error) {
	// fan-out: start concurrent requests
	userCh := make(chan userResult, 1)
	postsCh := make(chan postsResult, 1)

	go func() {
		user, err := s.userRepo.GetUser(userId)
		userCh <- userResult{user, err}
	}()

	go func() {
		posts, err := s.postRepo.GetPostsByUserId(userId)
		postsCh <- postsResult{posts, err}
	}()

	// fan-in: collect results
	var dashboard domain.Dashboard
	var dashboardErr error

	for range 2 {
		select {
		case res := <-userCh:
			if res.err != nil {
				dashboardErr = res.err
			} else {
				dashboard.User = res.user
			}
		case res := <-postsCh:
			if res.err != nil {
				dashboardErr = res.err
			} else {
				dashboard.Posts = res.posts
			}
		}
	}

	if dashboardErr != nil {
		return nil, dashboardErr
	}

	return &dashboard, nil
}
