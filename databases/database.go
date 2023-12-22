package databases

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Client *sql.DB

func Init() error {
	dsn := fmt.Sprintf("root:548926@tcp(127.0.0.1:3306)/gowhere?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(fmt.Errorf("open gorm err: %v", err))
		return err
	}
	Client = db
	return nil
}
