package dashboard

import (
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/article"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmInvest"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmMonitor"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transaction"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transactionDetail"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/middlewares"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sort"
	"time"
)

type Repo struct {
	DB *gorm.DB
}

func (r Repo) AdminDashboard(dashboard *entities.Dashboard) error {
	dashboardDb := FromUseCase(dashboard)

	var transactions []transaction.Transaction
	var articles []article.Article
	var products []product.Product
	paidTransactionIDs := make([]uuid.UUID, 0)
	var transactionDetails []transactionDetail.TransactionDetail
	var farmInvests []farmInvest.FarmInvest

	// Chart section
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30)

	result := r.DB.Joins("JOIN payments ON payments.id = transactions.payment_id").
		Where("payments.status = ?", "PAID").
		Where("transactions.created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&transactions)

	if result.Error != nil {
		return result.Error
	}

	summaries := make(map[string]float64)

	currentDate := startDate
	for !currentDate.After(endDate) {
		day := currentDate.Format("2006-01-02")
		summaries[day] = 0
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	for _, tx := range transactions {
		paidTransactionIDs = append(paidTransactionIDs, tx.ID)
		dashboardDb.EarningThisMonth += tx.Total
		day := tx.CreatedAt.Format("2006-01-02")
		summaries[day] += tx.Total
	}

	var sortedDates []string
	for date := range summaries {
		sortedDates = append(sortedDates, date)
	}
	sort.Strings(sortedDates)

	for _, date := range sortedDates {
		dashboardDb.EarningCharts = append(dashboardDb.EarningCharts, EarningChart{Date: date, Total: summaries[date]})
	}

	// Top Article section
	if err := r.DB.Preload("Admin").Order("created_at DESC").Limit(3).Find(&articles).Error; err != nil {
		return err
	}
	dashboardDb.LatestArticles = articles

	// Top Product section
	if err := r.DB.Order("created_at DESC").Limit(3).Find(&products).Error; err != nil {
		return err
	}
	dashboardDb.LatestProducts = products

	// product sold this month
	if err := r.DB.Where("transaction_id IN ?", paidTransactionIDs).Find(&transactionDetails).Error; err != nil {
		return err
	}
	for _, tx := range transactionDetails {
		dashboardDb.ProductSoldThisMonth += tx.Quantity
	}

	// farm invests total
	if err := r.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&farmInvests).Error; err != nil {
		return err
	}
	for _, tx := range farmInvests {
		dashboardDb.FarmsInvestThisMonth += tx.InvestmentAmount
	}

	*dashboard = *dashboardDb.ToUseCase()
	return nil
}

func (r Repo) CustomerDashboard(dashboard *entities.Dashboard, userData *middlewares.Claims) error {
	dashboardDb := FromUseCase(dashboard)

	var articles []article.Article
	var products []product.Product
	var farmInvestData farmInvest.FarmInvest
	var farmMonitorData farmMonitor.FarmMonitor

	if err := r.DB.Order("created_at DESC").Limit(3).Find(&articles).Error; err != nil {
		return err
	}
	dashboardDb.LatestArticles = articles

	if err := r.DB.Find(&products).Error; err != nil {
		return err
	}
	dashboardDb.AllProducts = products

	lengthAllProducts := len(products)
	if lengthAllProducts < 4 {
		dashboardDb.LatestProducts = products
	} else {
		dashboardDb.LatestProducts = products[lengthAllProducts-4:]
	}

	if err := r.DB.Where("customer_id = ?", userData.ID).Find(&farmInvestData).Error; err != nil {
		return err
	}

	if farmInvestData.ID != uuid.Nil {
		if err := r.DB.Order("created_at desc").Preload("Farm").Where("farm_id = ?", farmInvestData.FarmID).Find(&farmMonitorData).Error; err != nil {
			return err
		}
		dashboardDb.FarmMonitor = farmMonitorData
	}

	*dashboard = *dashboardDb.ToUseCase()
	return nil
}

func NewDashboardRepo(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}
