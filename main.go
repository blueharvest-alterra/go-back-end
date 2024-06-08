package main

import (
	"github.com/blueharvest-alterra/go-back-end/config"
	addressController "github.com/blueharvest-alterra/go-back-end/controllers/address"
	adminController "github.com/blueharvest-alterra/go-back-end/controllers/admin"
	articleController "github.com/blueharvest-alterra/go-back-end/controllers/article"
	courierController "github.com/blueharvest-alterra/go-back-end/controllers/courier"
	customerController "github.com/blueharvest-alterra/go-back-end/controllers/customer"
	farmController "github.com/blueharvest-alterra/go-back-end/controllers/farm"
	farmInvestController "github.com/blueharvest-alterra/go-back-end/controllers/farminvest"
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
	farmInvestRP "github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmInvest"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/promo"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transaction"
	"github.com/blueharvest-alterra/go-back-end/routes"
	"github.com/blueharvest-alterra/go-back-end/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitConfigPostgresql()
	db := postgresql.ConnectDB(config.InitConfigPostgresql())

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

	farmInvestRepo := farmInvestRP.NewFarmInvestRepo(db)
	farmInvestUseCase := usecases.NewFarmInvestUseCase(farmInvestRepo)
	newFarmInvestController := farmInvestController.NewFarmInvestController(farmInvestUseCase)

	farmRouteController := routes.FarmRouteController{
		FarmController: newFarmController,
	}
	promoRouteController := routes.PromoRouteController{
		PromoController: newPromoController,
	}
	articleRouteController := routes.ArticleRouteController{
		ArticleController: newArticleController,
	}
	adminRouteController := routes.AdminRouteController{
		AdminController: newAdminController,
	}
	customerRouteController := routes.CustomerRouteController{
		CustomerController: newCustomerController,
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

	farmInvestRouteController := routes.FarmInvestRouteController{
		FarmInvestController: newFarmInvestController,
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
	farmInvestRouteController.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
