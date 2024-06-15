package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/blueharvest-alterra/go-back-end/constant"
	"github.com/blueharvest-alterra/go-back-end/entities"
	"github.com/google/uuid"
)

type Payment struct {
	ID         uuid.UUID `gorm:"type:varchar(100);"`
	ExternalID string    `gorm:"type:text;not null"`
	InvoiceURL string    `gorm:"type:text;not null"`
	Status     string    `gorm:"type:varchar(50);not null"`
	Amount     float64   `gorm:"type:decimal;not null"`
}

func FromUseCase(payment *entities.Payment) *Payment {
	return &Payment{
		ID:         payment.ID,
		ExternalID: payment.ExternalID,
		InvoiceURL: payment.InvoiceURL,
		Status:     payment.Status,
		Amount:     payment.Amount,
	}
}

func (p *Payment) ToUseCase() *entities.Payment {
	return &entities.Payment{
		ID:         p.ID,
		ExternalID: p.ExternalID,
		InvoiceURL: p.InvoiceURL,
		Status:     p.Status,
		Amount:     p.Amount,
	}
}

type xdtInvoiceResponse struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Amount     int64  `json:"amount"`
	InvoiceURL string `json:"invoice_url"`
}

type xdtInvoicePayload struct {
	ExternalID  string `json:"external_id"`
	Amount      int64  `json:"amount"`
	PayerEmail  string `json:"payer_email"`
	Description string `json:"description"`
}

func (p *Payment) Create() error {
	url := "https://api.xendit.co/v2/invoices"
	method := "POST"

	var payload xdtInvoicePayload
	payload.ExternalID = p.ID.String()
	payload.Amount = int64(p.Amount)
	payload.PayerEmail = "blueharvest@gmail.com"
	payload.Description = "Blueharvest App Invoice"
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	xenditAPIKey := os.Getenv("XDT_SECRET_API_KEY")
	if xenditAPIKey == "" {
		xenditAPIKey = utils.GetConfig("XDT_SECRET_API_KEY")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+xenditAPIKey)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.Status != "200 OK" {
		return constant.ErrPaymentGateway
	}

	var response xdtInvoiceResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	p.ExternalID = response.ID
	p.InvoiceURL = response.InvoiceURL
	fmt.Println("hit", p)

	return nil
}
