package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-learning/backend-demo/util"
)

func (m *mysqlTestSuite) TestTransferTx() {
	t := m.T()

	store := NewStore(m.db)

	aid1 := m.createAccount(t, CreateAccountParams{
		AccountOwner: util.RandomString(6),
		Balance:      util.RandomInt64(100, 1000),
		Currency:     "RMB",
	})
	aid2 := m.createAccount(t, CreateAccountParams{
		AccountOwner: util.RandomString(6),
		Balance:      util.RandomInt64(100, 1000),
		Currency:     "RMB",
	})

	amount := int64(10)
	n := 5

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			res, err := store.TransferTx(context.Background(), TransferTxParams{
				FromID: aid1,
				ToID:   aid2,
				Amount: amount,
			})

			errs <- err
			results <- res
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		res := <-results

		require.NoError(t, err)

		// check transfer
		require.NotEmpty(t, res.Transfer)
		require.NotZero(t, res.Transfer.ID)
		require.Equal(t, aid1, res.Transfer.FromID)
		require.Equal(t, aid2, res.Transfer.ToID)
		require.Equal(t, amount, res.Transfer.Amount)

		// check from entry
		require.NotEmpty(t, res.FromEntry)
		require.Equal(t, aid1, res.FromEntry.AccountID)
		require.Equal(t, -amount, res.FromEntry.Amount)

		// check to entry
		require.NotEmpty(t, res.ToEntry)
		require.Equal(t, aid2, res.ToEntry.AccountID)
		require.Equal(t, amount, res.ToEntry.Amount)
	}
}
