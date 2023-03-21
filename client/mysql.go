package client

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shootz-developer/gtool/constant"

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
	db, err := sql.Open(constant.MySQLDrive, sourceName)
	if err != nil {
		log.Fatalf("New Mysql Client Error %+v", err)
		return nil, err
	}
	return db, nil
}

// StringHandler 处理从数据库里查出来的字符串的信息
func StringHandler(results *sql.Rows) []string {
	characters := make([]string, 0)
	for results.Next() {
		character := ""
		err := results.Scan(&character)
		if err != nil {
			log.Printf("Scan string err: [%+v]", err)
			continue
		}

		characters = append(characters, character)
	}

	return characters
}

// NullStringHandler 处理从数据库里查出来的字符串的信息
func NullStringHandler(results *sql.Rows) []sql.NullString {
	characters := make([]sql.NullString, 0)
	for results.Next() {
		character := sql.NullString{}
		err := results.Scan(&character.String)
		if err != nil {
			log.Printf("Scan string err: [%+v]", err)
			continue
		}

		characters = append(characters, character)
	}

	return characters
}

// NumberHandler 处理从数据库里查出来的数字的信息
func NumberHandler(results *sql.Rows) []int {
	numbers := make([]int, 0)
	for results.Next() {
		number := 0
		err := results.Scan(&number)
		if err != nil {
			log.Printf("Scan number err: [%+v]", err)
			continue
		}

		numbers = append(numbers, number)
	}
	return numbers
}
