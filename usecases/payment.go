package usecases

import "github.com/blueharvest-alterra/go-back-end/entities"

type PaymentUseCase struct {
	repository entities.PaymentRepositoryInterface
}

func (p PaymentUseCase) UpdateStatus(payment *entities.Payment) (entities.Payment, error) {
	if err := p.repository.UpdateStatus(payment); err != nil {
		return entities.Payment{}, err
	}

	return *payment, nil
}

func NewPaymentUseCase(repository entities.PaymentRepositoryInterface) *PaymentUseCase {
	return &PaymentUseCase{
		repository: repository,
	}
}
