package main

import (
	"github.com/blueharvest-alterra/go-back-end/config"
	adminController "github.com/blueharvest-alterra/go-back-end/controllers/admin"
	customerController "github.com/blueharvest-alterra/go-back-end/controllers/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/admin"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
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

	adminRouteController := routes.AdminRouteController{
		AdminController: newAdminController,
	}
	customerRouteController := routes.CustomerRouteController{
		CustomerController: newCustomerController,
	}

	adminRouteController.InitRoute(e)
	customerRouteController.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
