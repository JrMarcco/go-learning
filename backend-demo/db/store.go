package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) withTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	if err = fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v\n", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromID int64 `json:"fromID"`
	ToID   int64 `json:"toID"`
	Amount int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"fromAccount"`
	ToAccount   Account  `json:"toAccount"`
	FromEntry   Entry    `json:"fromEntry"`
	ToEntry     Entry    `json:"toEntry"`
}

func (s *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := s.withTx(ctx, func(queries *Queries) error {

		var err error

		// create transfer record
		transfer, err := s.createTransfer(ctx, CreateTransferParams{
			FromID: args.FromID,
			ToID:   args.ToID,
			Amount: args.Amount,
		})
		if err != nil {
			return err
		}
		result.Transfer = transfer

		// create from entry
		fromEntry, err := s.createEntry(ctx, CreateEntryParams{
			AccountID: args.FromID,
			Amount:    -args.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry = fromEntry

		// create to entry
		toEntry, err := s.createEntry(ctx, CreateEntryParams{
			AccountID: args.ToID,
			Amount:    args.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry = toEntry

		// TODO: update account balance

		return err
	})

	return result, err
}

func (s *Store) createTransfer(ctx context.Context, args CreateTransferParams) (Transfer, error) {
	res, err := s.CreateTransfer(ctx, args)
	if err != nil {
		return Transfer{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Transfer{}, err
	}

	return Transfer{
		ID: sql.NullInt64{
			Int64: id,
			Valid: true,
		},
		FromID: args.FromID,
		ToID:   args.ToID,
		Amount: args.Amount,
	}, nil
}

func (s *Store) createEntry(ctx context.Context, args CreateEntryParams) (Entry, error) {
	res, err := s.CreateEntry(ctx, CreateEntryParams{
		AccountID: args.AccountID,
		Amount:    args.Amount,
	})
	if err != nil {
		return Entry{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Entry{}, err
	}

	return Entry{
		ID: sql.NullInt64{
			Int64: id,
			Valid: true,
		},
		AccountID: args.AccountID,
		Amount:    args.Amount,
	}, nil
}