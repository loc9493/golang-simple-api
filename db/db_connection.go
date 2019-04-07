package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// func CreateCon() *sql.DB {
// 	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/quote")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("db is connected")
// 	}
// 	//defer db.Close()
// 	// make sure connection is available
// 	err = db.Ping()
// 	fmt.Println(err)
// 	if err != nil {
// 		fmt.Println("db is not connected")
// 		fmt.Println(err.Error())
// 	}
// 	return db
// }

func CreateCon() *gorm.DB {
	db, err := gorm.Open("mysql", "root@tcp(localhost:3306)/quote?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB connected")
	}
	return db
}
