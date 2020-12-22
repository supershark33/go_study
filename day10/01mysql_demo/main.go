package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //这里前面加"_"表示引用包，但是不直接引用包内容，只执行其中的init方法
)

func main()  {

	//链接数据库
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/goday10"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open %s invalid, err : %v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("dsb : %s failed, err : %v\n", dsn, err)
		return
	}
	fmt.Printf("数据库连接成功!!!")
}
