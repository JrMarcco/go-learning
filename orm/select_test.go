package orm

import (
	"entgo.io/ent/dialect/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelector_Build(t *testing.T) {

	tcs := []struct {
		name     string
		builder  StatBuilder
		wantStat *Statement
		wantErr  error
	}{
		{
			name:    "basic * select without from",
			builder: &Selector[TestModel]{},
			wantStat: &Statement{
				SQL: "SELECT * FROM `TestModel`;",
			},
			wantErr: nil,
		}, {
			name:    "basic * select with from",
			builder: (&Selector[TestModel]{}).From("test_model"),
			wantStat: &Statement{
				SQL: "SELECT * FROM `test_model`;",
			},
			wantErr: nil,
		}, {
			name:    "basic * select with empty from",
			builder: (&Selector[TestModel]{}).From(""),
			wantStat: &Statement{
				SQL: "SELECT * FROM `TestModel`;",
			},
			wantErr: nil,
		}, {
			name:    "basic * select with from db name",
			builder: (&Selector[TestModel]{}).From("test_db.test_model"),
			wantStat: &Statement{
				SQL: "SELECT * FROM `test_db`.`test_model`;",
			},
			wantErr: nil,
		}, {
			name:    "basic where",
			builder: (&Selector[TestModel]{}).Where(Col("Age").Eq(18)),
			wantStat: &Statement{
				SQL:  "SELECT * FROM `TestModel` WHERE `Age`=?;",
				Args: []any{18},
			},
			wantErr: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			stat, err := tc.builder.Build()
			assert.Equal(t, tc.wantErr, err)

			if err == nil {
				assert.Equal(t, tc.wantStat, stat)
			}
		})
	}

}

type TestModel struct {
	Id        int64
	Age       int8
	FirstName string
	LastName  *sql.NullString
}
