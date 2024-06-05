package main

import (
	"github.com/blueharvest-alterra/go-back-end/config"
	addressController "github.com/blueharvest-alterra/go-back-end/controllers/address"
	adminController "github.com/blueharvest-alterra/go-back-end/controllers/admin"
	customerController "github.com/blueharvest-alterra/go-back-end/controllers/customer"
	farmController "github.com/blueharvest-alterra/go-back-end/controllers/farm"
	productController "github.com/blueharvest-alterra/go-back-end/controllers/product"
	promoController "github.com/blueharvest-alterra/go-back-end/controllers/promo"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/address"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/admin"
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

	addressRepo := address.NewAddressRepo(db)
	addressUseCase := usecases.NewAddressUseCase(addressRepo)
	newAddressController := addressController.NewAddressController(addressUseCase)

	promoRouteController := routes.PromoRouteController{
		PromoController: newPromoController,
	}
	adminRouteController := routes.AdminRouteController{
		AdminController: newAdminController,
	}
	customerRouteController := routes.CustomerRouteController{
		CustomerController: newCustomerController,
	}
	farmRouteController := routes.FarmRouteController{
		FarmController: newFarmController,
	}
	productRouteController := routes.ProductRouteController{
		ProductController: newProductController,
	}
	addressRouteController := routes.AddressRouteController{
		AddressController: newAddressController,
	}

	adminRouteController.InitRoute(e)
	customerRouteController.InitRoute(e)
	farmRouteController.InitRoute(e)
	promoRouteController.InitRoute(e)
	productRouteController.InitRoute(e)
	addressRouteController.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
