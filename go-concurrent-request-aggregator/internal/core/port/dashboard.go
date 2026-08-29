package port

import "github.com/yehezkiel1086/go-concurrent-request-aggregator/internal/core/domain"

type DashboardService interface {
	GetDashboard(userId string) (*domain.Dashboard, error)
}
