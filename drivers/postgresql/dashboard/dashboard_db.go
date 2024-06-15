package dashboard

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/article"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/entities"
)

type EarningChart struct {
	Date  string
	Total float64
}

type Dashboard struct {
	EarningThisMonth     float64
	ProductSoldThisMonth uint
	FarmsInvestThisMonth float64
	EarningCharts        []EarningChart
	LatestArticles       []article.Article
	LatestProducts       []product.Product
}

func FromUseCase(dashboard *entities.Dashboard) *Dashboard {
	topArticles := make([]article.Article, len(dashboard.LatestArticles))
	for i, _article := range dashboard.LatestArticles {
		topArticles[i] = article.Article{
			ID:        _article.ID,
			Title:     _article.Title,
			Content:   _article.Content,
			Picture:   _article.Picture,
			CreatedAt: _article.CreatedAt,
		}
	}

	topProducts := make([]product.Product, len(dashboard.LatestProducts))
	for i, _product := range dashboard.LatestProducts {
		topProducts[i] = product.Product{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      _product.Status,
			Price:       _product.Price,
			CreatedAt:   _product.CreatedAt,
		}
	}

	earningChart := make([]EarningChart, len(dashboard.EarningCharts))
	for i, _earningChart := range dashboard.EarningCharts {
		earningChart[i] = EarningChart{
			Date:  _earningChart.Date,
			Total: _earningChart.Total,
		}
	}

	return &Dashboard{
		EarningThisMonth:     dashboard.EarningLastThirtyDays,
		ProductSoldThisMonth: dashboard.ProductSoldLastThirtyDays,
		FarmsInvestThisMonth: dashboard.FarmsInvestLastThirtyDays,
		EarningCharts:        earningChart,
		LatestArticles:       topArticles,
		LatestProducts:       topProducts,
	}
}

func (u *Dashboard) ToUseCase() *entities.Dashboard {
	topArticles := make([]entities.Article, len(u.LatestArticles))
	for i, _article := range u.LatestArticles {
		topArticles[i] = entities.Article{
			ID:        _article.ID,
			Title:     _article.Title,
			Content:   _article.Content,
			Picture:   _article.Picture,
			CreatedAt: _article.CreatedAt,
		}
	}

	topProducts := make([]entities.Product, len(u.LatestProducts))
	for i, _product := range u.LatestProducts {
		topProducts[i] = entities.Product{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      _product.Status,
			Price:       _product.Price,
			CreatedAt:   _product.CreatedAt,
		}
	}

	earningChart := make([]entities.EarningChart, len(u.EarningCharts))
	for i, _earningChart := range u.EarningCharts {
		earningChart[i] = entities.EarningChart{
			Date:  _earningChart.Date,
			Total: _earningChart.Total,
		}
	}

	return &entities.Dashboard{
		EarningLastThirtyDays:     u.EarningThisMonth,
		ProductSoldLastThirtyDays: u.ProductSoldThisMonth,
		FarmsInvestLastThirtyDays: u.FarmsInvestThisMonth,
		EarningCharts:             earningChart,
		LatestArticles:            topArticles,
		LatestProducts:            topProducts,
	}
}
