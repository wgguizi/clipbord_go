package clipboard

import (
	"fmt"
	"strings"
	"time"
	"viry_sun/lib/db"
	"viry_sun/lib/log"

	"go.uber.org/zap"
)

const (
	DATA_TABLE = "c_data"
)

type CData struct {
	Id        int64  `db:"id"`
	Code      string `db:"code"`
	Content   string `db:"content"`
	Ip        string `db:"ip"`
	CreatedAt int64  `db:"created_at"`
}

// 根据code查询记录
func GetByCode(code string) *CData {
	if code == "" {
		return nil
	}
	bingData := []interface{}{}
	bingData = append(bingData, code)

	cData := CData{}
	sql := `SELECT %s FROM %s
	WHERE code=?`
	sql = fmt.Sprintf(sql, "*", DATA_TABLE)
	db.Get(&cData, sql, bingData...)

	return &cData
}

// 保存数据
func Save(cData *CData) string {
	if cData.Code == "" {
		return "数据编码必须"
	}

	fcData := GetByCode(cData.Code)
	if fcData != nil && fcData.Id > 0 {
		//更新
		sql := `UPDATE %s SET content=?, ip=?, created_at=UNIX_TIMESTAMP() WHERE id=?`
		sql = fmt.Sprintf(sql, DATA_TABLE)
		_, err := db.DB.Exec(sql, cData.Content, cData.Ip, fcData.Id)
		if err != nil {
			return "更新失败"
		}
		return ""
	} else {
		//插入
		sql := `INSERT IGNORE INTO %s (code, content, ip, created_at) VALUES(?, ?, ?, UNIX_TIMESTAMP())`
		sql = fmt.Sprintf(sql, DATA_TABLE)
		_, err := db.DB.Exec(sql, cData.Code, cData.Content, cData.Ip)
		if err != nil {
			log.L.Error(fmt.Sprintf("Insert Error: %v", err), zap.Stack("Default Stack:"))
			return "插入失败"
		}
		return ""
	}
}

func removeExpireOnce(expireSeconds int) (numRt int64, errRt string) {
	sql := `SELECT id FROM %s WHERE created_at<UNIX_TIMESTAMP()-%d LIMIT 1000`
	sql = fmt.Sprintf(sql, DATA_TABLE, expireSeconds)
	cDatas := []CData{}
	db.DB.Select(&cDatas, sql)

	//log.L.Error(fmt.Sprintf("SQL: %s", sql))

	idArr := []string{}
	for _, v := range cDatas {
		idArr = append(idArr, fmt.Sprintf("%d", v.Id))
	}

	if len(idArr) > 0 {
		sql := `DELETE FROM %s WHERE id IN (%s)`
		sql = fmt.Sprintf(sql, DATA_TABLE, strings.Join(idArr, ","))

		sqlRet, err := db.DB.Exec(sql)
		if err != nil {
			log.L.Error(fmt.Sprintf("SELECT Error: %v", err), zap.Stack("Default Stack:"))
			errRt = "查找失败"
		}
		numRt, _ = sqlRet.RowsAffected()
	}
	return numRt, errRt
}

func RemoveExpire(expireSeconds int) (numRt int64, errRt string) {
	num, err := removeExpireOnce(expireSeconds)
	numRt = num
	t1 := time.NewTicker(50 * time.Millisecond)
	for num > 0 && err == "" {
		<-t1.C //延迟后再执行下次
		num, err = removeExpireOnce(expireSeconds)
		numRt += num
	}

	errRt = err
	return numRt, errRt
}
