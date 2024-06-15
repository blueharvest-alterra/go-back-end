package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"time"
)

type FarmDetailResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Picture     string    `json:"picture"`
}

type FarmMonitorDetailResponse struct {
	ID              uuid.UUID          `json:"id"`
	FarmID          uuid.UUID          `json:"farm_id"`
	Farm            FarmDetailResponse `json:"farm"`
	Temperature     float64            `json:"temperature"`
	PH              float64            `json:"ph"`
	DissolvedOxygen float64            `json:"dissolved_oxygen"`
	CreatedAt       time.Time          `json:"created_at"`
}

type CustomerDashboard struct {
	FarmMonitor    FarmMonitorDetailResponse `json:"farm_monitor"`
	LatestArticles []ArticleDetailResponse   `json:"latest_articles"`
	LatestProducts []ProductDetailResponse   `json:"latest_products"`
	AllProducts    []ProductDetailResponse   `json:"all_products"`
}

func CustomerDashboardFromUseCase(dashboard *entities.Dashboard) *CustomerDashboard {
	latestArticles := make([]ArticleDetailResponse, len(dashboard.LatestArticles))
	for i, _article := range dashboard.LatestArticles {
		latestArticles[i] = ArticleDetailResponse{
			ID:      _article.ID,
			Title:   _article.Title,
			Content: _article.Content,
			Picture: _article.Picture,
		}
	}

	latestProducts := make([]ProductDetailResponse, len(dashboard.LatestProducts))
	for i, _product := range dashboard.LatestProducts {
		latestProducts[i] = ProductDetailResponse{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      _product.Status,
			Price:       _product.Price,
		}
	}

	allProducts := make([]ProductDetailResponse, len(dashboard.LatestProducts))
	for i, _product := range dashboard.LatestProducts {
		allProducts[i] = ProductDetailResponse{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      _product.Status,
			Price:       _product.Price,
		}
	}

	return &CustomerDashboard{
		FarmMonitor: FarmMonitorDetailResponse{
			ID:     dashboard.FarmMonitor.ID,
			FarmID: dashboard.FarmMonitor.FarmID,
			Farm: FarmDetailResponse{
				ID:          dashboard.FarmMonitor.Farm.ID,
				Title:       dashboard.FarmMonitor.Farm.Title,
				Description: dashboard.FarmMonitor.Farm.Description,
				Picture:     dashboard.FarmMonitor.Farm.Picture,
			},
			Temperature:     dashboard.FarmMonitor.Temperature,
			PH:              dashboard.FarmMonitor.PH,
			DissolvedOxygen: dashboard.FarmMonitor.DissolvedOxygen,
			CreatedAt:       dashboard.FarmMonitor.CreatedAt,
		},
		LatestArticles: latestArticles,
		LatestProducts: latestProducts,
		AllProducts:    allProducts,
	}
}
