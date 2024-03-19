package db

import (
	"fmt"
	"viry_sun/lib/config"
	"viry_sun/lib/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	db, err := sqlx.Connect(config.C.Db.Type, config.C.Db.Dns)
	if err != nil {
		log.L.Panic(fmt.Sprintf("DB Connect Error! error: %v", err))
	} else {
		DB = db
	}
}

// 获取单条数据
// 数据为空，则不报错
func Get(dest interface{}, sql string, args ...interface{}) {
	row := DB.QueryRowx(sql, args...)

	if err := row.Err(); err != nil {
		log.L.Error(fmt.Sprintf("%s, bind:%v, err:%#v", sql, args, err))
	}

	row.StructScan(dest)
}
