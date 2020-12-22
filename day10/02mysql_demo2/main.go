package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//初始化数据库
	err:= initDb()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	batchQuery(1)
	insert("二爸", 27)
	batchQuery(1)
	updateRow(9000, 2)
	batchQuery(1)
	deleteById(7)
	batchQuery(1)
}

var db *sql.DB

func initDb() (err error){

	dsn:= "root:rootroot@tcp(127.0.0.1:3306)/goday10"
	db, err = sql.Open("mysql", dsn) //注意 := 和 = 的区别，相当申请局部变量
	if err != nil {
		return //前提是治理的err和函数返回值里面的（err error）参数名字一样
	}
	err = db.Ping()
	if err != nil {
		return
	}
	//数据库连接成功
	fmt.Println("数据库连接成功...")
	db.SetMaxOpenConns(10)
	return
}

type user struct {
	id int
	name string
	age int
}

func queryOne(id int) {
	var u1 user
	sqlStr := "select id,name,age from user where id = ?"
	//scan方法内部会释放连接
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	fmt.Println(u1)
}

func batchQuery(id int) {

	sqlStr := "select id,name,age from user where id >= ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil{
		fmt.Println("批量查询失败")
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		_ = rows.Scan(&u.id, &u.name, &u.age)
		fmt.Println(u)
	}
}

func insert(name string, age int) {

	sqlStr := `insert into user(id, name, age) values(7, ?, ?)`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("id : ", id)
}

func updateRow(newAge int, id int)  {

	updateStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(updateStr, newAge, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(affected)
}

func deleteById(id int)  {

	deleteStr := "delete from user where id = ?"
	ret, err := db.Exec(deleteStr, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(affected)
}
