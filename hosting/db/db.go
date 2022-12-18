package db

import (
	"fmt"
)

func init() {
	fmt.Println("init")
	// DB, err := sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer DB.Close()
}

func Test() {
	fmt.Println("test")
}
