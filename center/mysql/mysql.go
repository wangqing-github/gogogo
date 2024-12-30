package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitConnection() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gogogo")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()
	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}
	fmt.Println("Successfully connected to MySQL!")
}
