package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"blog/common"
	"time"
)

var orm *xorm.Engine

func SetEngine() *xorm.Engine {
	dbhost := common.Cfg.MustValue("db", "dbhost", "127.0.0.1")
	dbport := common.Cfg.MustValue("db", "dbport", "root")
	dbuser := common.Cfg.MustValue("db", "dbuser", "user")
	dbpass := common.Cfg.MustValue("db", "dbpass", "pass")
	dbname := common.Cfg.MustValue("db", "dbname", "test")
	orm, _ = xorm.NewEngine("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8")
	orm.TZLocation = time.Local
	return orm
}