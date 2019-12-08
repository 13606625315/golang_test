package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"flag"
)

type User struct {
	Id   int64
	Name string `xorm:"varchar(64) notnull unique 'user_name'"`
	Age  int    `xorm:"default(18)"`
}

func (u *User) String() string {
	return fmt.Sprintf("Id:%d Name:%q Age:%d", u.Id, u.Name, u.Age)
}
 


func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

    var db string
    var name string
    var oper string
    flag.StringVar(&db,"db","test","数据库")
    flag.StringVar(&name,"name","","姓名")
    flag.StringVar(&oper,"oper","query","增删改查")
    flag.Parse()

	fmt.Println("Hello ab = ",db,"name = ",name,"oper = ",oper)

	//	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(192.168.4.119:3306)/?charset=utf8")
	if err != nil {
		log.Printf("new engine failed: %s\n", err)
		return
	}

	if oper == "drop"{
		engine.Exec("DROP DATABASE " + db)
		fmt.Println("drop database = ",db)
		return;
	}

	engine.Exec("CREATE DATABASE " + db)

	engine, err = xorm.NewEngine("mysql", "root:root@tcp(192.168.4.119:3306)/"+ db +"?charset=utf8")
	if err != nil {
		log.Printf("new engine failed: %s\n", err)
		return
	}

	if err := engine.Ping();err != nil{
		fmt.Println(err)
	}

	engine.ShowSQL(true)

	if oper == "insert" {
		engine.Sync2(new(User))

		user1 := &User{Name: name}
		affected, err := engine.Insert(user1)
		if err != nil {
			log.Println(err)
			//return
		}
		log.Println(affected)
	}else if oper == "query"{
		if(name == ""){
			tests := make([]User, 0)
			errr := engine.Distinct("id", "user_name", "age").Limit(10, 0).Find(&tests)
			if errr != nil {
				panic(errr)
			}
			fmt.Printf("总共查询出 %d 条数据\n", len(tests))
			for _, v := range tests {
				fmt.Printf("信息Id: %d, 姓名: %s, 邮箱: %s\n", v.Id, v.Name, v.Age)
			}
		}else{
			user := new(User)
			user.Name = name
			has, _ := engine.Get(user)
			if has {
				log.Println(user)
			}
		}

	}else if oper == "delete"{

	}else if oper == "update"{

	}
}
