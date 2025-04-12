package repository

import (
	"database/sql"
	"time"

	"github.com/thiagotnunes08/go-gateway-api/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(
		`INSERT INTO accounts (id,name,email,api_key,balance,created_at,updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		`)
	if err != nil {

		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM accounts
		WHERE api_key = $1
	`, apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&createdAt,
		&updatedAt,
	)

	if err == sql.ErrNoRows {

		return nil, domain.ErrAccountNotFound
	}

	if err != nil {

		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt
	return &account, nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {

	var account domain.Account
	var created_at, updated_at time.Time

	err := r.db.QueryRow(
		` SELECT id,name,email,api_key,balance,created_at,updated_at 
		from accounts 
		where id = $1 `,
		id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err == sql.ErrNoRows {

		return nil, domain.ErrAccountNotFound
	}

	if err != nil {

		return nil, err
	}

	account.CreatedAt = created_at
	account.UpdatedAt = updated_at
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {

	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var currentBalance float64

	err = tx.QueryRow(`SELECT balance from accounts where id = $1 for update`,
		account.ID).Scan(&currentBalance)

	if err == sql.ErrNoRows {

		return domain.ErrAccountNotFound
	}

	if err != nil {

		return err
	}

	_, err = tx.Exec(`update accounts 
	                  set balance = $1, 
					  updated_at = $2 
					  where id = $3
					  `, currentBalance+account.Balance, time.Now(), account.ID)

	if err != nil {

		return err
	}

	return tx.Commit()
}
