package dashboard

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/article"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farm"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmMonitor"
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
	// customer dashboard
	FarmMonitor farmMonitor.FarmMonitor
	AllProducts []product.Product
}

func FromUseCase(dashboard *entities.Dashboard) *Dashboard {
	latestArticles := make([]article.Article, len(dashboard.LatestArticles))
	for i, _article := range dashboard.LatestArticles {
		latestArticles[i] = article.Article{
			ID:        _article.ID,
			Title:     _article.Title,
			Content:   _article.Content,
			Picture:   _article.Picture,
			CreatedAt: _article.CreatedAt,
		}
	}

	latestProducts := make([]product.Product, len(dashboard.LatestProducts))
	for i, _product := range dashboard.LatestProducts {
		latestProducts[i] = product.Product{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      product.Status(_product.Status),
			Price:       _product.Price,
			CreatedAt:   _product.CreatedAt,
		}
	}

	allProducts := make([]product.Product, len(dashboard.LatestProducts))
	for i, _product := range dashboard.LatestProducts {
		allProducts[i] = product.Product{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      product.Status(_product.Status),
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
		LatestArticles:       latestArticles,
		LatestProducts:       latestProducts,
		FarmMonitor: farmMonitor.FarmMonitor{
			ID:     dashboard.FarmMonitor.ID,
			FarmID: dashboard.FarmMonitor.FarmID,
			Farm: farm.Farm{
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
		AllProducts: allProducts,
	}
}

func (u *Dashboard) ToUseCase() *entities.Dashboard {
	latestArticles := make([]entities.Article, len(u.LatestArticles))
	for i, _article := range u.LatestArticles {
		latestArticles[i] = entities.Article{
			ID:      _article.ID,
			Title:   _article.Title,
			Content: _article.Content,
			Picture: _article.Picture,
			Admin: entities.Admin{
				ID:       _article.Admin.ID,
				FullName: _article.Admin.FullName,
			},
			CreatedAt: _article.CreatedAt,
		}
	}

	latestProducts := make([]entities.Product, len(u.LatestProducts))
	for i, _product := range u.LatestProducts {
		latestProducts[i] = entities.Product{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      entities.ProductStatus(_product.Status),
			Price:       _product.Price,
			CountSold:   _product.CountSold,
			CreatedAt:   _product.CreatedAt,
		}
	}

	allProducts := make([]entities.Product, len(u.LatestProducts))
	for i, _product := range u.LatestProducts {
		allProducts[i] = entities.Product{
			ID:          _product.ID,
			Name:        _product.Name,
			Description: _product.Description,
			Thumbnail:   _product.Thumbnail,
			Status:      entities.ProductStatus(_product.Status),
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
		LatestArticles:            latestArticles,
		LatestProducts:            latestProducts,
		FarmMonitor: entities.FarmMonitor{
			ID:     u.FarmMonitor.ID,
			FarmID: u.FarmMonitor.FarmID,
			Farm: entities.Farm{
				ID:          u.FarmMonitor.Farm.ID,
				Title:       u.FarmMonitor.Farm.Title,
				Description: u.FarmMonitor.Farm.Description,
				Picture:     u.FarmMonitor.Farm.Picture,
			},
			Temperature:     u.FarmMonitor.Temperature,
			PH:              u.FarmMonitor.PH,
			DissolvedOxygen: u.FarmMonitor.DissolvedOxygen,
			CreatedAt:       u.FarmMonitor.CreatedAt,
		},
		AllProducts: allProducts,
	}
}
