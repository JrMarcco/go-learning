package service

import (
	"log"

	"entgo.io/ent/dialect/sql"
	"github.com/JrMarcco/go-learning/grpc/pcbook/config"
	"github.com/JrMarcco/go-learning/grpc/pcbook/db"

	_ "github.com/go-sql-driver/mysql"
)

var entClient *db.Client

func Setup() {
	dbCfg := config.DbCfg()

	drv, err := sql.Open(dbCfg.Driver, dbCfg.Source)
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB := drv.DB()
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdle)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpen)

	entClient = db.NewClient(db.Driver(drv))
}

func TearDown() {
	defer func() {
		if entClient != nil {
			_ = entClient.Close()
		}
	}()
}
