package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func initDb() (err error) {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/goday10"
	db, _ = sql.Open("mysql", dsn)
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(20)
	return
}

func prepareInsert() {

	prepareSqlStr := `insert into user(name, age) values(?,?)`
	stmt, err := db.Prepare(prepareSqlStr)
	if err != nil {
		fmt.Println("prepare failed, err : ", err)
		return
	}
	defer stmt.Close()
	//后续那stmt去处理一些请求任务
	var m = map[string]int{
		"九九":5,
		"风扇👄":88,
		"吃好吃吃嘴":99,
	}
	for k,v := range m{
		ret, _ := stmt.Exec(k, v)
		fmt.Println(ret.LastInsertId())
	}
}

func prepareSelect() {

}

func main() {
	err := initDb()
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	prepareInsert()
}


