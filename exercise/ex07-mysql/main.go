package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


//func Open(driveName, dataSourceName string) (*DB, error){
//
//}


func main() {
	dsn := "root:root1234@tcp(127.0.0.1:13306)/sql_demo"
	db,err := sql.Open("mysql", dsn)
	if err!= nil{
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connect to db ... ")


	err2 := db.Ping()
	if err2 != nil{
		fmt.Println("Connect to db failed, err2: %v\n ", err2)
		return
	}
	fmt.Println("Connectr to db Sucess !!! ")

	db.SetConnMaxLifetime(time.Second*10)
	db.SetMaxOpenConns(280)
	db.SetMaxIdleConns(10)


}
