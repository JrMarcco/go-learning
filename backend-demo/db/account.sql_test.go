package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"go-learning/backend-demo/util"
	"testing"
)

func (m *mysqlTestSuite) createAccount(t *testing.T, account CreateAccountParams) int64 {
	res, err := m.queries.CreateAccount(context.Background(), account)
	require.NoError(t, err)
	id, err := res.LastInsertId()

	require.NoError(t, err)
	require.NotZero(t, id)

	return id
}

func (m *mysqlTestSuite) TestCreateAccount() {
	t := m.T()

	_ = m.createAccount(t, CreateAccountParams{
		AccountOwner: util.RandomString(6),
		Balance:      util.RandomInt64(100, 10000),
		Currency:     "RMB",
	})
}

func (m *mysqlTestSuite) TestGetAccount() {
	t := m.T()

	args := CreateAccountParams{
		AccountOwner: util.RandomString(6),
		Balance:      util.RandomInt64(100, 10000),
		Currency:     "RMB",
	}

	id := m.createAccount(t, args)

	account, err := m.queries.GetAccount(context.Background(), sql.NullInt64{
		Int64: id,
		Valid: true,
	})

	require.NoError(t, err)
	require.Equal(t, args.AccountOwner, account.AccountOwner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
}

func (m *mysqlTestSuite) TestDeleteAccount() {
	t := m.T()

	args := CreateAccountParams{
		AccountOwner: util.RandomString(6),
		Balance:      util.RandomInt64(100, 10000),
		Currency:     "RMB",
	}

	id := m.createAccount(t, args)

	err := m.queries.DeleteAccount(context.Background(), sql.NullInt64{
		Int64: id,
		Valid: true,
	})

	require.NoError(t, err)

	account, err := m.queries.GetAccount(context.Background(), sql.NullInt64{
		Int64: id,
		Valid: true,
	})
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func (m *mysqlTestSuite) TestUpdateAccount() {
	t := m.T()

	createArgs := CreateAccountParams{
		AccountOwner: util.RandomString(6),
		Balance:      util.RandomInt64(100, 10000),
		Currency:     "RMB",
	}

	id := m.createAccount(t, createArgs)
	updateArgs := UpdateAccountParams{
		ID: sql.NullInt64{
			Int64: id,
			Valid: true,
		},
		Balance: util.RandomInt64(100, 10000),
	}

	err := m.queries.UpdateAccount(context.Background(), updateArgs)
	require.NoError(t, err)

	account, err := m.queries.GetAccount(context.Background(), sql.NullInt64{
		Int64: id,
		Valid: true,
	})

	require.NoError(t, err)
	require.Equal(t, updateArgs.Balance, account.Balance)
}
