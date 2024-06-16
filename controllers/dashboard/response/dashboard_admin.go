package response

import (
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type EarningChartResponse struct {
	Date  string  `json:"date"`
	Total float64 `json:"value"`
}

type ArticleDetailResponse struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Picture string    `json:"picture"`
}

type ProductDetailResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	Thumbnail   string    `json:"thumbnail"`
}

type Dashboard struct {
	EarningLastThirtyDays     float64                 `json:"earning_last_thirty_days"`
	ProductSoldLastThirtyDays uint                    `json:"product_sold_last_thirty_days"`
	FarmsInvestLastThirtyDays float64                 `json:"farms_invest_last_thirty_days"`
	EarningCharts             []EarningChartResponse  `json:"earning_charts"`
	LatestArticles            []ArticleDetailResponse `json:"latest_articles"`
	LatestProducts            []ProductDetailResponse `json:"latest_products"`
}

func AdminDashboardFromUseCase(dashboard *entities.Dashboard) *Dashboard {
	topArticles := make([]ArticleDetailResponse, len(dashboard.LatestArticles))
	for i, _article := range dashboard.LatestArticles {
		topArticles[i] = ArticleDetailResponse{
			ID:      _article.ID,
			Title:   _article.Title,
			Content: _article.Content,
			Picture: _article.Picture,
		}
	}

	topProducts := make([]ProductDetailResponse, len(dashboard.LatestProducts))
	for i, _product := range dashboard.LatestProducts {
		topProducts[i] = ProductDetailResponse{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      string(_product.Status),
			Price:       _product.Price,
		}
	}

	earningCharts := make([]EarningChartResponse, len(dashboard.EarningCharts))
	for i, _earningChart := range dashboard.EarningCharts {
		earningCharts[i] = EarningChartResponse{
			Date:  _earningChart.Date,
			Total: _earningChart.Total,
		}
	}

	return &Dashboard{
		EarningLastThirtyDays:     dashboard.EarningLastThirtyDays,
		ProductSoldLastThirtyDays: dashboard.ProductSoldLastThirtyDays,
		FarmsInvestLastThirtyDays: dashboard.FarmsInvestLastThirtyDays,
		EarningCharts:             earningCharts,
		LatestArticles:            topArticles,
		LatestProducts:            topProducts,
	}
}
