package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/xormplus/core"
	"github.com/xormplus/xorm"
//	"flag"
)

type User struct {
	Id   int64
	Name string `xorm:"varchar(64) notnull 'user_name'"`
	Age  int    `xorm:"default(18)"`
	nn string `xorm:"varchar(64) notnull 'user_nn'"`
}

func (u *User) TableName() string {
	return u.nn
}

func (u *User) setTableName(str string)  {
	u.nn = str
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
		count ,errr := this.Oper.engine.Distinct("id", "user_name", "age").FindAndCount(&tests)

		if errr != nil {
			log.Println(errr)
			panic(errr)
		}
		fmt.Printf("总共查询出 %d 条数据\n", count)
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