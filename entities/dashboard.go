package entities

import "github.com/blueharvest-alterra/go-back-end/middlewares"

type EarningChart struct {
	Date  string
	Total float64
}

type Dashboard struct {
	EarningLastThirtyDays     float64
	ProductSoldLastThirtyDays uint
	FarmsInvestLastThirtyDays float64
	EarningCharts             []EarningChart
	LatestArticles            []Article
	LatestProducts            []Product
}

type DashboardRepositoryInterface interface {
	AdminDashboard(dashboard *Dashboard) error
	CustomerDashboard(dashboard *Dashboard) error
}

type DashboardUseCaseInterface interface {
	AdminDashboard(dashboard *Dashboard, userData *middlewares.Claims) (Dashboard, error)
	CustomerDashboard(dashboard *Dashboard, userData *middlewares.Claims) (Dashboard, error)
}
