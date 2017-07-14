package main

import (
	_"github.com/go-sql-driver/mysql"
	_"github.com/jmoiron/sqlx"
	_"github.com/astaxie/beedb"
	"database/sql"
	"fmt"
	"encoding/json"
	"github.com/go-martini/martini"
	"strings"
	"strconv"
	_"github.com/devfeel/dotweb"
)

var db *sql.DB

func init(){
	var error error
	db ,error = sql.Open("mysql","saokajun:Xx978271541@tcp(118.89.184.104:3306)/sosoga?charset=utf8")
	if error!=nil{
		panic(error)
	}
}

type user struct {
	UserId int
	OpenId string
	NickName string
	HeadUrl string
	AvailableCredit int
	GlobalCredit int
	CreateTime string
	LastupdateTime string
}

func selectUser(id int) *user{
	var myUser user
	rows, err := db.Query("select * from user where userId = ? ", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(	&myUser.UserId,
					&myUser.OpenId,
					&myUser.NickName,
					&myUser.HeadUrl,
					&myUser.AvailableCredit,
					&myUser.GlobalCredit,
					&myUser.CreateTime,
					&myUser.LastupdateTime)
		if err != nil {
			panic(err.Error())
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	return &myUser
}

func selectAll(tableName string){
	rows, _:= db.Query("SELECT * FROM "+tableName)
	columns, _:= rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]sql.RawBytes, len(columns))

	//result := make([]user,len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		rows.Scan(scanArgs...)
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value)
		}
	}

	defer rows.Close()
}

func deletUserById(userId int) sql.Result {
	result,err := db.Exec("DELETE from user where userId = ? ", userId)
	if err != nil {
		panic(err.Error())
	}
	return result
}

func main(){

	m := martini.Classic()

	m.Get("/selectUserById/:userId",func(params martini.Params) []byte {
		userId,_:=strconv.Atoi(strings.Split(params["userId"],"=")[1])
		byte,_ := json.Marshal(selectUser(userId))
		return byte
	})

	m.Get("/deletUserById/:userId",func(params martini.Params){
		userId,_:=strconv.Atoi(strings.Split(params["userId"],"=")[1])
		deletUserById(userId)
	})

	m.Get("/test/:userId",func(params martini.Params){
		fmt.Println(params["userId"])
	})

	m.RunOnAddr(":8080")
	m.Run()

	defer db.Close()
}
