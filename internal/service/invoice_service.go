package service

import (
	"github.com/thiagotnunes08/go-gateway-api/internal/domain"
	"github.com/thiagotnunes08/go-gateway-api/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {

	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

func (s *InvoiceService) Create(input dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {

	account, err := s.accountService.FindByAPIKey(input.APIKey)

	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, account.ID)

	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(input.APIKey, invoice.Amount)

		if err != nil {
			return nil, err
		}

	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err

	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) GetByID(id, apiKey string) (*dto.InvoiceOutput, error) {

	invoice, err := s.invoiceRepository.FindByID(id)

	if err != nil {
		return nil, err
	}

	accountOutPut, err := s.accountService.FindByAPIKey(apiKey)

	if err != nil {
		return nil, err
	}

	if invoice.AccountId != accountOutPut.ID {
		return nil, domain.ErrUnauthorizedAcces
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*dto.InvoiceOutput, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accountID)

	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutput, len(invoices))

	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}

	return output, nil
}

func (s *InvoiceService) ListByAccountAPIKey(apiKey string) ([]*dto.InvoiceOutput, error) {

	accounts, err := s.accountService.FindByAPIKey(apiKey)

	if err != nil {
		return nil, err
	}

	return s.ListByAccount(accounts.ID)
}
