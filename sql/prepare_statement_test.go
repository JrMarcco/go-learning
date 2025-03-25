package sql

import (
	"context"
	"github.com/stretchr/testify/assert"
)

func (s *sqlTestSuite) TestPrepareStatement() {
	t := s.T()

	stmt, err := s.db.Prepare("select  * from `test_model` where `id` = ?;")
	if err != nil {
		t.Fatal(err)
	}

	_, err = stmt.QueryContext(context.Background(), 1)
	assert.Nil(t, err)

	err = stmt.Close()
	assert.Nil(t, err)
}
