package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sqlx.DB //是一个连接池对象
)

func main() {

	err := initDB()
	if err != nil {
		fmt.Println("connect failed")
		return
	}
	fmt.Println("数据库连接池初始化成功")
	getById(1)
	batchGet(1)
	updateById(2, 9002)
	batchGet(1)
}

func initDB() (err error) {

	dsn := "root:rootroot@tcp(127.0.0.1:3306)/goday10"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	return
}

func getById(id int)  {

	queryStr:= "select id,name,age from user where id = ?"
	var u user
	err := db.Get(&u, queryStr, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)
}

func batchGet(id int) {

	queryStr := "select id,name,age from user where id >= ?"
	var users []user
	//slice本来就是个引用，因此这里用users就好，不用&users
	err := db.Select(&users, queryStr, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}

func updateById(id int, newAge int){

	updateStr := "update user set age = ? where id = ?"
	ret, _ := db.Exec(updateStr, newAge, id)
	fmt.Println(ret.RowsAffected())
}

func deleteById(id int)  {

}

//这里用到了反射，因此要传指针，并且首字母大写，保证别的包可以看到
type user struct {
	Id int
	Name string
	Age int
}


