package transaction

import (
	"errors"
	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/address"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/courier"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/customer"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/payment"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/product"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/promo"
	"github.com/blueharvest-alterra/go-back-end/drivers/postgresql/transactionDetail"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const TaxFee = float64(10000)

type Transaction struct {
	ID                 uuid.UUID `gorm:"type:varchar(100)"`
	CustomerID         uuid.UUID `gorm:"type:varchar(100)"`
	Customer           customer.Customer
	SubTotal           float64   `gorm:"type:decimal"`
	Tax                float64   `gorm:"type:decimal"`
	Discount           float64   `gorm:"type:decimal"`
	Total              float64   `gorm:"type:decimal"`
	PaymentID          uuid.UUID `gorm:"type:varchar(100)"`
	Payment            payment.Payment
	CourierID          uuid.UUID   `gorm:"type:varchar(100)"`
	PromoID            uuid.UUID   `gorm:"type:varchar(100)"`
	Promo              promo.Promo `gorm:"-"`
	Courier            courier.Courier
	TransactionDetails []transactionDetail.TransactionDetail
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func FromUseCase(transaction *entities.Transaction) *Transaction {
	allTransactionDetails := make([]transactionDetail.TransactionDetail, len(transaction.TransactionDetails))
	for i, _transactionDetail := range transaction.TransactionDetails {
		allTransactionDetails[i] = transactionDetail.TransactionDetail{
			ID:            _transactionDetail.ID,
			TransactionID: _transactionDetail.TransactionID,
			Product: product.Product{
				ID:          _transactionDetail.Product.ID,
				Name:        _transactionDetail.Product.Name,
				Description: _transactionDetail.Product.Description,
				Price:       _transactionDetail.Product.Price,
				Status:      product.Status(_transactionDetail.Product.Status),
				Thumbnail:   _transactionDetail.Product.Thumbnail,
			},
			Quantity:   _transactionDetail.Quantity,
			TotalPrice: _transactionDetail.TotalPrice,
		}
	}

	return &Transaction{
		ID:         transaction.ID,
		CustomerID: transaction.CustomerID,
		Customer: customer.Customer{
			ID:          transaction.Customer.ID,
			FullName:    transaction.Customer.FullName,
			PhoneNumber: transaction.Customer.PhoneNumber,
			BirthDate:   transaction.Customer.BirthDate,
		},
		SubTotal:  transaction.SubTotal,
		Tax:       transaction.Tax,
		Discount:  transaction.Discount,
		Total:     transaction.Total,
		PaymentID: transaction.PaymentID,
		Payment: payment.Payment{
			ID:         transaction.Payment.ID,
			ExternalID: transaction.Payment.ExternalID,
			InvoiceURL: transaction.Payment.InvoiceURL,
			Status:     transaction.Payment.Status,
		},
		CourierID: transaction.CourierID,
		PromoID:   transaction.PromoID,
		Promo: promo.Promo{
			ID:     transaction.Promo.ID,
			Name:   transaction.Promo.Name,
			Code:   transaction.Promo.Code,
			Status: promo.PromoStatus(transaction.Promo.Status),
			Amount: transaction.Promo.Amount,
		},
		Courier: courier.Courier{
			ID:                   transaction.Courier.ID,
			DestinationAddressID: transaction.Courier.DestinationAddressID,
			DestinationAddress: address.Address{
				ID:        transaction.Courier.DestinationAddress.ID,
				Address:   transaction.Courier.DestinationAddress.Address,
				CityID:    transaction.Courier.DestinationAddress.CityID,
				City:      transaction.Courier.DestinationAddress.City,
				StateID:   transaction.Courier.DestinationAddress.StateID,
				State:     transaction.Courier.DestinationAddress.State,
				ZipCode:   transaction.Courier.DestinationAddress.ZipCode,
				Country:   transaction.Courier.DestinationAddress.Country,
				Longitude: transaction.Courier.DestinationAddress.Longitude,
				Latitude:  transaction.Courier.DestinationAddress.Latitude,
			},
			Name: transaction.Courier.Name,
			Fee:  transaction.Courier.Fee,
			Type: transaction.Courier.Type,
		},
		TransactionDetails: allTransactionDetails,
		CreatedAt:          transaction.CreatedAt,
	}
}

func (u *Transaction) ToUseCase() *entities.Transaction {
	allTransactionDetails := make([]entities.TransactionDetail, len(u.TransactionDetails))
	for i, _transactionDetail := range u.TransactionDetails {
		allTransactionDetails[i] = entities.TransactionDetail{
			ID:            _transactionDetail.ID,
			TransactionID: _transactionDetail.TransactionID,
			ProductID:     _transactionDetail.ProductID,
			Product: entities.Product{
				ID:          _transactionDetail.Product.ID,
				Name:        _transactionDetail.Product.Name,
				Description: _transactionDetail.Product.Description,
				Price:       _transactionDetail.Product.Price,
				Status:      entities.ProductStatus(_transactionDetail.Product.Status),
				Thumbnail:   _transactionDetail.Product.Thumbnail,
			},
			Quantity:   _transactionDetail.Quantity,
			TotalPrice: _transactionDetail.TotalPrice,
		}
	}

	return &entities.Transaction{
		ID:         u.ID,
		CustomerID: u.CustomerID,
		Customer: entities.Customer{
			ID:          u.Customer.ID,
			FullName:    u.Customer.FullName,
			PhoneNumber: u.Customer.PhoneNumber,
			BirthDate:   u.Customer.BirthDate,
		},
		SubTotal: u.SubTotal,
		Tax:      u.Tax,
		Discount: u.Discount,
		Total:    u.Total,
		Payment: entities.Payment{
			ID:         u.Payment.ID,
			ExternalID: u.Payment.ExternalID,
			InvoiceURL: u.Payment.InvoiceURL,
			Status:     u.Payment.Status,
		},
		CourierID: u.CourierID,
		PromoID:   u.PromoID,
		Promo: entities.Promo{
			ID:     u.Promo.ID,
			Name:   u.Promo.Name,
			Code:   u.Promo.Code,
			Status: entities.PromoStatus(u.Promo.Status),
			Amount: u.Promo.Amount,
		},
		Courier: entities.Courier{
			ID:                   u.Courier.ID,
			DestinationAddressID: u.Courier.DestinationAddressID,
			DestinationAddress: entities.Address{
				ID:        u.Courier.DestinationAddress.ID,
				Address:   u.Courier.DestinationAddress.Address,
				CityID:    u.Courier.DestinationAddress.CityID,
				City:      u.Courier.DestinationAddress.City,
				StateID:   u.Courier.DestinationAddress.StateID,
				State:     u.Courier.DestinationAddress.State,
				ZipCode:   u.Courier.DestinationAddress.ZipCode,
				Country:   u.Courier.DestinationAddress.Country,
				Longitude: u.Courier.DestinationAddress.Longitude,
				Latitude:  u.Courier.DestinationAddress.Latitude,
			},
			Name: u.Courier.Name,
			Fee:  u.Courier.Fee,
			Type: u.Courier.Type,
		},
		TransactionDetails: allTransactionDetails,
		CreatedAt:          u.CreatedAt,
	}
}

func (u *Transaction) SetCustomerData(db *gorm.DB) error {
	if err := db.Preload("Addresses").Where("id = ?", u.Customer.ID).First(&u.Customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrCustomerNotFound
		}
		return err
	}
	return nil
}

func (u *Transaction) SetPromoData(db *gorm.DB) error {
	if err := db.Where("id = ?", u.PromoID).First(&u.Promo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrPromoNotFound
		}
		return err
	}

	if u.Promo.Status == promo.Unavailable {
		return constant.ErrPromoUnavailable
	}

	u.Discount = u.Promo.Amount
	return nil
}

func (u *Transaction) SetAddressData() error {
	for _, _address := range u.Customer.Addresses {
		if _address.ID == u.Courier.DestinationAddressID {
			u.Courier.DestinationAddress = _address
		}
	}

	if u.Courier.DestinationAddress.ID == uuid.Nil {
		return constant.ErrCustomerAddressNotFound
	}

	return nil
}

func (u *Transaction) SetTransactionDetail(db *gorm.DB, transactionItem *transactionDetail.TransactionDetail) error {
	if err := db.Where(&product.Product{ID: transactionItem.ProductID}).First(&transactionItem.Product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return constant.ErrProductNotFound
		}
		return err
	}

	transactionItem.TotalPrice = transactionItem.Product.Price * float64(transactionItem.Quantity)
	u.SubTotal += transactionItem.TotalPrice

	if transactionItem.Product.Status == "unavailable" {
		return constant.ErrProductUnavailable
	}

	return nil
}
