package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"go-learning/backend-demo/util"
)

func (m *mysqlTestSuite) TestCreateAccount() {
	t := m.T()

	params := CreateAccountParams{
		AccountOwner: util.RandomString(6),
		Balance:      util.RandomInt64(100, 10000),
		Currency:     "RMB",
	}

	err := m.queries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
}
