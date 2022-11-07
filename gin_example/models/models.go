package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go_learning/gin_example/pkg/logging"
	"go_learning/gin_example/pkg/setting"
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

var db *gorm.DB

func SetUp() {
	var err error
	db, err = gorm.Open(
		setting.DbSetting.Type,
		fmt.Sprintf(""+
			"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DbSetting.User,
			setting.DbSetting.Password,
			setting.DbSetting.Host,
			setting.DbSetting.Name,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DbSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxOpenConns(100)

	db.Callback().Create().Replace("gorm:update_time_stamp", createCb)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateCb)
	db.Callback().Delete().Replace("gorm:delete", deleteCb)
}

func CloseDB() {
	defer func(db *gorm.DB) {
		if err := db.Close(); err != nil {
		}
	}(db)
}

func createCb(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now().Unix()
		if createdOn, ok := scope.FieldByName("CreatedOn"); ok && createdOn.IsBlank {
			if err := createdOn.Set(now); err != nil {
				logging.Error("Fail to set CreatedOn", err)
			}
		}

		if modifiedOn, ok := scope.FieldByName("ModifiedOn"); ok && modifiedOn.IsBlank {
			if err := modifiedOn.Set(now); err != nil {
				logging.Error("Fail to set ModifiedOn", err)
			}
		}
	}
}

func updateCb(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		if err := scope.SetColumn("ModifiedOn", time.Now().Unix()); err != nil {
			logging.Error("Fail to set ModifiedOn", err)
		}
	}
}

func deleteCb(scope *gorm.Scope) {
	if !scope.HasError() {
		deletedOn, ok := scope.FieldByName("DeletedOn")
		if !scope.Search.Unscoped && ok {
			var extraOp string
			if str, ok := scope.Get("gorm:delete_option"); ok {
				extraOp = fmt.Sprint(str)
			}

			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOn.DBName),
				scope.AddToVars(time.Now().Unix()),
				AddExtraSpaceIfExist(scope.CombinedConditionSql()),
				AddExtraSpaceIfExist(extraOp),
			)).Exec()
		}
	}
}

func AddExtraSpaceIfExist(str string) string {
	if str != "" {
		return "" + str
	}
	return str
}
