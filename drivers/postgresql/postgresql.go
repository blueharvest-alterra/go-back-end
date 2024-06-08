package postgresql

import (
	"fmt"

	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/address"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/admin"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/article"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/auth"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/cart"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/courier"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farm"
	farminvest "github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmInvest"
	farmmonitor "github.com/blueharvest-alterra/go-back-end/drivers/postgresql/farmMonitor"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/promo"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transaction"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transactionDetail"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Name string
	User string
	Pass string
	Host string
	Port string
}

func ConnectDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		config.Host,
		config.User,
		config.Pass,
		config.Name,
		config.Port,
	)
	fmt.Println("Connecting to", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic(err)
	}

	MigrationUser(db)
	return db
}

func MigrationUser(db *gorm.DB) {
	db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'promo_status') THEN CREATE TYPE promo_status AS ENUM ('available', 'unavailable'); END IF; END $$;")

	err := db.AutoMigrate(
		auth.Auth{},
		customer.Customer{},
		address.Address{},
		admin.Admin{},
		farm.Farm{},
		promo.Promo{},
		article.Article{},
		product.Product{},
		transaction.Transaction{},
		transactionDetail.TransactionDetail{},
		courier.Courier{},
		farmmonitor.FarmMonitor{},
		farminvest.FarmInvest{},
		cart.Cart{},
	)
	if err != nil {
		return
	}
}
