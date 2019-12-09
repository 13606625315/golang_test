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
	Name string `xorm:"varchar(64) notnull 'user_name'"`
	Age  int    `xorm:"default(18)"`
}

type oper_i interface{
	handle()
}

type oper_i_1 interface{
	init(engine *xorm.Engine,items ...interface{})
}

type oper_s struct{
	user [2]*User
	engine *xorm.Engine
}

func (_ *oper_s)handle(){

}

func (this *oper_s)init(engine *xorm.Engine,items ...interface{}){
//	this.user = user
	this.engine = engine
	for i, x := range items {
		switch a:=x.(type) {
		case *User:
			this.user[i] = a
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}	
}

type insert_s struct{
	Oper *oper_s
}
func (this *insert_s)handle(){
	this.Oper.engine.Sync2(this.Oper.user[0])
	affected, err := this.Oper.engine.Insert(this.Oper.user[0])
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(affected)
}

type update_s struct{
	Oper *oper_s
}
func (this *update_s)handle(){
	//this.Oper.engine.Sync2(this.Oper.user)
	tests := make([]User, 0)
	this.Oper.engine.Where("user_name = ?", this.Oper.user[0].Name).Find(&tests)	

	for _,v := range tests{
		v.Name = this.Oper.user[1].Name
		affected, err := this.Oper.engine.ID(v.Id).Update(v)
		if err != nil {
			log.Println(err)
			return
		}		
		log.Println(affected)
	}

	//this.Oper.user[0].Name = this.Oper.user[1].Name;
	//fmt.Println("name = ",this.Oper.user[0].Name)
//	affected, err := this.Oper.engine.Update(tests)
//	affected, err := this.Oper.engine.Update(this.Oper.user[0])
//	if err != nil {
//		log.Println(err)
	//	return
//	}
//	log.Println(affected)
}

type delete_s struct{
	Oper *oper_s
}
func (this *delete_s)handle(){

	affected, err := this.Oper.engine.Delete(this.Oper.user[0])
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(affected)
}

type query_s struct{
	Oper *oper_s
}
func (this *query_s)handle(){
	//this.Oper.engine.Sync2(this.Oper.user)
	if(this.Oper.user[0].Name == ""){
		tests := make([]User, 0)
		errr := this.Oper.engine.Distinct("id", "user_name", "age").Find(&tests)
		if errr != nil {
			panic(errr)
		}
		fmt.Printf("总共查询出 %d 条数据\n", len(tests))
		for _, v := range tests {
			fmt.Printf("Id: %d, 姓名: %s, 年纪: %d\n", v.Id, v.Name, v.Age)
		}
	}else{
		user := new(User)
		user.Name = this.Oper.user[0].Name
		has, _ := this.Oper.engine.Get(user)
		if has {
			log.Println(user)
		}
	}
}

func factory_oper_init(oper string) (oper_i,oper_i_1){
	fmt.Println("1111111");
	switch oper {
	case "insert":
		m:=&insert_s{}
		m.Oper = &oper_s{}
		return m, m.Oper
	case "update":
		m:=&update_s{}
		m.Oper = &oper_s{}		
		return m, m.Oper
	case "delete":
		m:=&delete_s{}
		m.Oper = &oper_s{}		
		return m, m.Oper
	case "query":
		m:=&query_s{}
		m.Oper = &oper_s{}		
		return m, m.Oper		
	default:
		panic("123");
		
	}
}

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
	op_init.init(engine, user,user1)
	op.handle()

	return 
/*
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
*/
}
