package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID             string
	AccountId      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Status string

const (
	StatusPeding   Status = "pending"
	StatusApproved Status = "paid"
	StatusRejected Status = "rejected"
)

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardHolderName string
}

func NewInvoice(accountID string, amout float64, decription string, paymentType string, card CreditCard) (*Invoice, error) {

	if amout < 0 {
		return nil, ErrInvalidAmout
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.NewString(),
		AccountId:      accountID,
		Amount:         amout,
		Status:         StatusPeding,
		Description:    decription,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {

	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))

	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus

	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {

	if i.Status != StatusPeding {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
