package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() {
	//insert()
	query()
	update()
	delete()
	transaction()

	defer func(database *sqlx.DB) {
		err := database.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(Db) // 注意这行代码要写在上面err判断的下面
}

func insert() {
	begin, _ := Db.Begin()

	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
	roll := begin.Rollback()
	fmt.Println(roll)
}

func query() {
	var person []Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Println("select succ:", person)
}

func update() {
	res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("update succ:", row)
}

func delete() {
	res, err := Db.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed,", err)
	}

	fmt.Println("delete succ:", row)
}

func transaction() {
	conn, _ := Db.Begin()
	_, _ = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	conn.Commit()
}
