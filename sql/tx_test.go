package sql

import (
	"context"
	"database/sql"
	"time"
)

func (s *sqlTestSuite) TestTx() {
	t := s.T()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		t.Fatal(err)
	}

	res, err := tx.ExecContext(ctx,
		"insert into `test_model`(`id`, `first_name`, `last_name`, `age`) values (?, ?, ?, ?);",
		2, "Tom", "Cat", 20,
	)
	if err != nil {
		t.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil || affected != 1 {
		t.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		t.Fatal(err)
	}
}
