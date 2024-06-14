package main

import (
	// "context"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/cart"
	// chatBot "github.com/blueharvest-alterra/go-back-end/drivers/postgresql/chat-bot"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmInvest"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmMonitor"
	// "github.com/blueharvest-alterra/go-back-end/drivers/redis"
	"log"

	"github.com/blueharvest-alterra/go-back-end/config"
	addressController "github.com/blueharvest-alterra/go-back-end/controllers/address"
	adminController "github.com/blueharvest-alterra/go-back-end/controllers/admin"
	articleController "github.com/blueharvest-alterra/go-back-end/controllers/article"
	cartController "github.com/blueharvest-alterra/go-back-end/controllers/cart"
	// chatBotController "github.com/blueharvest-alterra/go-back-end/controllers/chat-bot"
	courierController "github.com/blueharvest-alterra/go-back-end/controllers/courier"
	customerController "github.com/blueharvest-alterra/go-back-end/controllers/customer"
	farmController "github.com/blueharvest-alterra/go-back-end/controllers/farm"
	farmInvestController "github.com/blueharvest-alterra/go-back-end/controllers/farm-invest"
	farmMonitorController "github.com/blueharvest-alterra/go-back-end/controllers/farm-monitor"
	paymentController "github.com/blueharvest-alterra/go-back-end/controllers/payment"
	productController "github.com/blueharvest-alterra/go-back-end/controllers/product"
	promoController "github.com/blueharvest-alterra/go-back-end/controllers/promo"
	transactionController "github.com/blueharvest-alterra/go-back-end/controllers/transaction"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/address"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/admin"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/article"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/courier"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farm"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/payment"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/promo"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transaction"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/blueharvest-alterra/go-back-end/routes"
	"github.com/blueharvest-alterra/go-back-end/usecases"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron"
)

// var ctx = context.Background()

func main() {
	dbConfig := config.InitConfigPostgresql()
	db := postgresql.ConnectDB(dbConfig)
	// redisConfig := config.InitConfigRedis()
	// redisClient := redis.ConnectRedis(redisConfig)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	adminRepo := admin.NewAdminRepo(db)
	adminUseCase := usecases.NewAdminUseCase(adminRepo)
	newAdminController := adminController.NewAdminController(adminUseCase)

	customerRepo := customer.NewCustomerRepo(db)
	customerUseCase := usecases.NewCustomerUseCase(customerRepo)
	newCustomerController := customerController.NewCustomerController(customerUseCase)

	farmRepo := farm.NewFarmRepo(db)
	farmUseCase := usecases.NewFarmUseCase(farmRepo)
	newFarmController := farmController.NewFarmController(farmUseCase)

	promoRepo := promo.NewPromoRepo(db)
	promoUseCase := usecases.NewPromoUseCase(promoRepo)
	newPromoController := promoController.NewPromoController(promoUseCase)

	productRepo := product.NewProductRepo(db)
	productUseCase := usecases.NewProductUseCase(productRepo)
	newProductController := productController.NewProductController(productUseCase)

	articleRepo := article.NewArticleRepo(db)
	articleUseCase := usecases.NewArticleUseCase(articleRepo)
	newArticleController := articleController.NewarticleController(articleUseCase)

	addressRepo := address.NewAddressRepo(db)
	addressUseCase := usecases.NewAddressUseCase(addressRepo)
	newAddressController := addressController.NewAddressController(addressUseCase)

	transactionRepo := transaction.NewTransactionRepo(db)
	transactionUseCase := usecases.NewTransactionUseCase(transactionRepo)
	newTransactionController := transactionController.NewTransactionController(transactionUseCase)

	courierRepo := courier.NewCourierRepo(db)
	courierUseCase := usecases.NewCourierUseCase(courierRepo)
	newCourierController := courierController.NewCourierController(courierUseCase)

	paymentRepo := payment.NewPaymentRepo(db)
	paymentUseCase := usecases.NewPaymentUseCase(paymentRepo)
	newPaymentController := paymentController.NewPaymentController(paymentUseCase)

	farmInvestRepo := farmInvest.NewFarmInvestRepo(db)
	farmInvestUseCase := usecases.NewFarmInvestUseCase(farmInvestRepo)
	newFarmInvestController := farmInvestController.NewFarmInvestController(farmInvestUseCase)

	farmMonitorRepo := farmMonitor.NewFarmMonitorRepo(db)
	farmMonitorUseCase := usecases.NewFarmMonitorUseCase(farmMonitorRepo)
	newFarmMonitorController := farmMonitorController.NewFarmMonitorController(farmMonitorUseCase)

	cartRepo := cart.NewCartRepo(db)
	cartUseCase := usecases.NewCartUseCase(cartRepo)
	newCartController := cartController.NewCartController(cartUseCase)
	
	// chatBotRepo := chatBot.NewChatBotRepo(db, redisClient)
	// chatBotUseCase := usecases.NewChatBotUseCase(chatBotRepo)
	// newChatBotController := chatBotController.NewChatBotController(chatBotUseCase)

	adminRouteController := routes.AdminRouteController{
		AdminController: newAdminController,
	}
	customerRouteController := routes.CustomerRouteController{
		CustomerController: newCustomerController,
	}
	farmRouteController := routes.FarmRouteController{
		FarmController: newFarmController,
	}
	promoRouteController := routes.PromoRouteController{
		PromoController: newPromoController,
	}
	articleRouteController := routes.ArticleRouteController{
		ArticleController: newArticleController,
	}
	productRouteController := routes.ProductRouteController{
		ProductController: newProductController,
	}
	addressRouteController := routes.AddressRouteController{
		AddressController: newAddressController,
	}
	transactionRouteController := routes.TransactionRouteController{
		TransactionController: newTransactionController,
	}
	courierRouteController := routes.CourierRouteController{
		CourierController: newCourierController,
	}
	paymentRouteController := routes.PaymentRouteController{
		PaymentController: newPaymentController,
	}
	farmInvestRouteController := routes.FarmInvestRouteController{
		FarmInvestController: newFarmInvestController,
	}
	farmMonitorRouteController := routes.FarmMonitorRouteController{
		FarmMonitorController: newFarmMonitorController,
	}
	// chatBotRouteController := routes.ChatBotRouteController{
	// 	ChatBotController: newChatBotController,
	// }

	cartRouteController := routes.CartRouteController{
		CartController: newCartController,
	}

	adminRouteController.InitRoute(e)
	customerRouteController.InitRoute(e)
	farmRouteController.InitRoute(e)
	promoRouteController.InitRoute(e)
	productRouteController.InitRoute(e)
	articleRouteController.InitRoute(e)
	addressRouteController.InitRoute(e)
	transactionRouteController.InitRoute(e)
	courierRouteController.InitRoute(e)
	paymentRouteController.InitRoute(e)
	farmInvestRouteController.InitRoute(e)
	farmMonitorRouteController.InitRoute(e)
	cartRouteController.InitRoute(e)
	// chatBotRouteController.InitRoute(e)

	//init cron
	c := cron.New()
	frp := farm.NewFarmRepo(db)
	fmrp := farmMonitor.NewFarmMonitorRepo(db)

	err := c.AddFunc("@daily", func() { processDailyFarmMonitor(frp, fmrp) })
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	c.Start()
	defer c.Stop()

	e.Logger.Fatal(e.Start(":8080"))
}

func processDailyFarmMonitor(frp *farm.Repo, fmrp *farmMonitor.Repo) {
	var farms []entities.Farm
	if err := frp.GetAll(&farms); err != nil {
		log.Fatalf("Error getting farms: %v", err)
	}

	for _, _farm := range farms {
		farmMonitor := entities.FarmMonitor{
			ID:              uuid.New(),
			FarmID:          _farm.ID,
			Temperature:     float64(0),
			PH:              float64(0),
			DissolvedOxygen: float64(0),
		}

		if err := fmrp.Create(&farmMonitor); err != nil {
			log.Fatalf("Error creating farm monitor: %v", err)
		}
	}
	log.Printf("Successfully running cron")
}
