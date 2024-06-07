package main

import (
	"github.com/blueharvest-alterra/go-back-end/config"
	addressController "github.com/blueharvest-alterra/go-back-end/controllers/address"
	adminController "github.com/blueharvest-alterra/go-back-end/controllers/admin"
	courierController "github.com/blueharvest-alterra/go-back-end/controllers/courier"
	articleController "github.com/blueharvest-alterra/go-back-end/controllers/article"
	customerController "github.com/blueharvest-alterra/go-back-end/controllers/customer"
	farmController "github.com/blueharvest-alterra/go-back-end/controllers/farm"
	productController "github.com/blueharvest-alterra/go-back-end/controllers/product"
	promoController "github.com/blueharvest-alterra/go-back-end/controllers/promo"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/address"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/admin"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/courier"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/article"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farm"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/promo"
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

	courierRepo := courier.NewCourierRepo(db)
	courierUseCase := usecases.NewCourierUseCase(courierRepo)
	newCourierController := courierController.NewCourierController(courierUseCase)

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
	courierRouteController := routes.CourierRouteController{
		CourierController: newCourierController,
	}

	adminRouteController.InitRoute(e)
	customerRouteController.InitRoute(e)
	farmRouteController.InitRoute(e)
	promoRouteController.InitRoute(e)
	productRouteController.InitRoute(e)
	articleRouteController.InitRoute(e)
	addressRouteController.InitRoute(e)
	courierRouteController.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
