package sql

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type sqlTestSuite struct {
	suite.Suite

	driver string
	dsn    string

	db *sql.DB
}

func TestMysql(t *testing.T) {
	suite.Run(t, &sqlTestSuite{
		driver: "mysql",
		dsn:    "root:u2E3WWtgam@tcp(192.168.3.50:31964)/test_db",
	})
}

func (s *sqlTestSuite) TearDownTest() {
	if _, err := s.db.Exec("DELETE FROM test_model;"); err != nil {
		s.T().Fatal(err)
	}
}

func (s *sqlTestSuite) SetupSuite() {
	db, err := sql.Open(s.driver, s.dsn)
	if err != nil {
		s.T().Fatal(err)
	}

	s.db = db

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = s.db.ExecContext(ctx, `
CREATE TABLE IF NOT EXISTS test_model (
    id integer primary key,
    first_name text not null,
	last_name text not null,
	json text not null,
	age integer
)
`)

	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *sqlTestSuite) TestCURD() {
	t := s.T()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 或者 Exec(xxx)
	res, err := s.db.ExecContext(ctx,
		"INSERT INTO `test_model`(`id`, `first_name`, `last_name`, `age`) VALUES (?, ?, ?, ?);",
		1, "Foo", "Bar", 18,
	)

	if err != nil {
		t.Fatal(err)
	}

	affected, err := res.RowsAffected()
	if err != nil || affected != 1 {
		t.Fatal(err)
	}

	rows, err := s.db.QueryContext(ctx,
		"SELECT `id`, `first_name`, `last_name`, `age` FROM `test_model` LIMIT ?",
		1,
	)
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		tm := TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age)
		if err != nil {
			_ = rows.Close()
			t.Fatal(err)
		}
		assert.Equal(t, "Foo", tm.FirstName)
		assert.Equal(t, "Bar", tm.LastName.String)
		assert.Equal(t, 18, tm.Age)

		_ = rows.Close()

		// 或者 Exec(xxx)
		res, err = s.db.ExecContext(ctx,
			"UPDATE `test_model` SET `first_name` = ? WHERE `id` = ?",
			"New Foo", 1,
		)
		if err != nil {
			t.Fatal(err)
		}

		affected, err = res.RowsAffected()
		if err != nil || affected != 1 {
			t.Fatal(err)
		}

		row := s.db.QueryRowContext(ctx,
			"SELECT `id`, `first_name`, `last_name`, `age` FROM `test_model` LIMIT ?",
			1,
		)
		if row.Err() != nil {
			t.Fatal(row.Err())
		}

		tm = TestModel{}

		if err = row.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "New Foo", tm.FirstName)

	}
}

type TestModel struct {
	Id        int64
	FirstName string
	LastName  *sql.NullString
	Age       int
}
