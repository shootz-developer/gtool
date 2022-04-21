package client

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MysqlProxy mysql的proxy
type MysqlProxy struct {
	Account  string
	Password string
	IP       string
	Port     string
	DBName   string
}

// NewMysqlClient 创建一个mysql的client
func NewMysqlClient(proxy MysqlProxy) (*sql.DB, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", proxy.Account, proxy.Password,
		proxy.IP, proxy.Port, proxy.DBName)
	db, err := sql.Open("mysql", sourceName)
	if err != nil {
		log.Fatalf("New Mysql Client Error %+v", err)
		return nil, err
	}
	return db, nil
}
