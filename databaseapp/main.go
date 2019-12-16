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

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

    var db string
    var name string
    var oper string
    flag.StringVar(&db,"db","test","数据库")
    flag.StringVar(&name,"name","","姓名")
	flag.StringVar(&oper,"oper","query","增删改查")
	
    var name_update string
    flag.StringVar(&name_update,"name_update","","更新的姓名")	

    flag.Parse()

	fmt.Println("Hello ab = ",db,"name = ",name,"oper = ",oper,"name_update = ",name_update)

	//	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/?charset=utf8")
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

	engine, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/"+ db +"?charset=utf8")
	if err != nil {
		log.Printf("new engine failed: %s\n", err)
		return
	}

	if err := engine.Ping();err != nil{
		fmt.Println(err)
	}

	engine.ShowSQL(true)

	op,op_init:= factory_oper_init(oper)
	user:=&User{Name:name }
	user1:=&User{Name:name_update}
	op_init.init(engine, user, user1)
	op.handle()
/*
	engine.DropTables("user")
	log.Println(111)
	user.nn = "user1"
	engine.CreateTables(user)
	log.Println(111)
	has,_ := engine.IsTableEmpty("Persons");
	fmt.Println(has);
	has,_ = engine.IsTableEmpty("user");
	fmt.Println(has);
	has,_ = engine.IsTableEmpty("123");
	fmt.Println(has);	

	has,_ = engine.IsTableExist("Persons");
	fmt.Println(has);
	has,_ = engine.IsTableExist("user");
	fmt.Println(has);
	has,_ = engine.IsTableExist("123");
	fmt.Println(has);	
*/
	return 
}
