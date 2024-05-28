package postgresql

import (
	"fmt"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/admin"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/auth"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
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
	err := db.AutoMigrate(auth.Auth{}, customer.Customer{}, admin.Admin{})
	if err != nil {
		return
	}
}
