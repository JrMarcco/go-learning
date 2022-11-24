package service

import (
	"entgo.io/ent/dialect/sql"
	"go-learning/grpc/pcbook/config"
	"go-learning/grpc/pcbook/db"
	"log"
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
