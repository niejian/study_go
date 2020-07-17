package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
var DB *sql.DB

type User struct {
	Id int64 `json:"id"`
	UserName string `json:"user_name"`
	Age int8 `json:"age"`
	CreateTime time.Time `json:"create_time"`

}

func initDB() (err error) {
	dsn := "apollo:qq123123@tcp(192.168.242.103)/test?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
		return err
	}
	//defer db.Close()
	DB = db
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping err %v \n", err)
		return err
	}
	return nil
}

// 查询
func queryDataById(id string) []*User  {
	var users []*User = make([]*User, 10)
	sqlStr := "select id, user_name, age, create_time from `user` where id > ?"
	fmt.Printf("DB %v \n", DB)
	rows, err := DB.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("查询失败：%v \n", err)

	}
	defer rows.Close()
	for rows.Next()  {
		var u *User = &User{
			Id:         0,
			UserName:   "",
			Age:        0,
			CreateTime: time.Time{},
		}
		err := rows.Scan(&u.Id, &u.UserName, &u.Age, &u.CreateTime)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)

		}
		users = append(users, u)
	}

	return users
}

// mysql 操作
func main() {
	err := initDB()
	if nil != initDB(){
		fmt.Printf("init db failed,err:%v\n", err)
		return
	} else {
		fmt.Printf("init db success \n")
	}


	// 设置最大连接数
	//db.SetMaxOpenConns(500)
	//// 最大空闲连接数
	//db.SetMaxIdleConns(100)

	users := queryDataById("1")
	length := len(users)
	for i := 0; i < length; i++ {
		u := users[i]
		if nil == u {
			continue
		}
		fmt.Printf("id:%v userName:%v age:%v, createTime:%v\n", u.Id, u.UserName, u.Age, u.CreateTime)
	}

}
