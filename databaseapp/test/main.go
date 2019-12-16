package main

import (
	"github.com/xormplus/xorm"
//	"guest_private/config"
//	"guest_private/vmp-service"
//	"strings"
//	"time"
	"strconv"
	//"code.raying.com/cloud///utils"
//	"code.raying.com/cloud///utils/builder"
//	"code.raying.com/cloud///utils/logs"
	"sync"
	"flag"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Num int64

var Engine *xorm.Engine

type Person struct {
	ID        int64  `xorm:"bigint pk autoincr 'ID'"`
	Name      string  `xorm:"varchar(20) 'Name'"`
	Sex       string `xorm:"char(10) 'Sex'"`
	GradesID     int64 `xorm:"int 'GradesID'"`
	Phone      string `xorm:"char(15) 'Phone'"`
	Role       string `xorm:"char(10) 'role'"`
	InviteCode  string `xorm:"char(32) unique 'InviteCode'"`
}

type Grades struct{
	ID        int64  `xorm:"bigint pk autoincr 'ID'"`
	Years      int64  `xorm:"int 'Years'"`	
	Classes       int64  `xorm:"int 'Classes'"`		
}

type Teacher2Grades struct{
	ID        int64  `xorm:"bigint pk autoincr 'ID'"`
	GradesID      int64 `xorm:"int 'GradesID'"`	
	TeacherID 	  int64 `xorm:"int 'TeacherID'"`
}
//func (u *Person) Init() {
//}

func (u *Person) TableName() string {
	return "person" +  strconv.FormatInt(Num,10)
}

func (u *Grades) TableName() string {
	return "grades" +  strconv.FormatInt(Num,10)
}

func (u *Teacher2Grades) TableName() string {
	return "teacher2Grades" +  strconv.FormatInt(Num,10)
}

var lock_Person = sync.Mutex{}

func (u *Person) CreateTable(o xorm.Interface, id int64) (err error) {
	lock_Person.Lock()
	Num = id	
	err = Engine.CreateTables(u)
	lock_Person.Unlock()	
	//utils.CheckError(err)
	return	
}

var lock_Grades = sync.Mutex{}

func (u *Grades) CreateTable(o xorm.Interface, id int64) (err error) {
	//var lock = sync.Mutex{}
	lock_Grades.Lock()
	Num = id	
	err = Engine.CreateTables(u)
	lock_Grades.Unlock()	
	//utils.CheckError(err)
	return	
}

var lock_Teacher2Grades = sync.Mutex{}

func (u *Teacher2Grades) CreateTable(o xorm.Interface, id int64) (err error) {
	//var lock = sync.Mutex{}
	lock_Teacher2Grades.Lock()
	Num = id	
	err = Engine.CreateTables(u)
	lock_Teacher2Grades.Unlock()	
	//utils.CheckError(err)
	return	
}

func (u *Person) Insert(o xorm.Interface, id int64) (err error) {
	if o == nil {
		o = Engine
	}
	session:=o.Table("person" + strconv.FormatInt(id,10))
	_, err =session.Insert(u)
	//utils.CheckError(err)
	return
}

func (u *Person) Update(o xorm.Interface, id int64, fields ...string) (err error) {
	if o == nil {
		o = Engine
	}
	session:=o.Table("person" + strconv.FormatInt(id,10))	
	_, err =session.ID(u.ID).Cols(fields...).Update(u)
	//utils.CheckError(err)
	return
}

func (u *Person) DeleteAll(o xorm.Interface, id int64) (err error) {
	if o == nil {
		o = Engine
	}
//	session:=o.Table("person" + strconv.FormatInt(id,10))	
	err = Engine.DropTables("person" + strconv.FormatInt(id,10))
	//utils.CheckError(err)
	return
}


func (u *Person) Query(o xorm.Interface, id int64) (count int64){
	if o == nil {
		o = Engine
	}	
	session:=o.Table("person" + strconv.FormatInt(id,10))	
	count,_ = session.FindAndCount(u)
	return
}

func (u *Grades) Insert(o xorm.Interface, id int64) (err error) {
	if o == nil {
		o = Engine
	}
	session:=o.Table("grades" + strconv.FormatInt(id,10))	
	_, err =session.Insert(u)
	//utils.CheckError(err)
	return
}

func (u *Grades) Update(o xorm.Interface, id int64, fields ...string) (err error) {
	if o == nil {
		o = Engine
	}
	session:=o.Table("grades" + strconv.FormatInt(id,10))	
	_, err =session.ID(u.ID).Cols(fields...).Update(u)
	//utils.CheckError(err)
	return
}

func (u *Grades) DeleteAll(o xorm.Interface, id int64) (err error) {
	if o == nil {
		o = Engine
	}
//	session:=o.Table("grades" + strconv.FormatInt(id,10))	
	err = Engine.DropTables("grades" + strconv.FormatInt(id,10))
	return
}


func (u *Teacher2Grades) Insert(o xorm.Interface, id int64) (err error) {
	if o == nil {
		o = Engine
	}
	session:=o.Table("teacher2Grades" + strconv.FormatInt(id,10))	
	_, err =session.Insert(u)
	//utils.CheckError(err)
	return
}

func (u *Teacher2Grades) Update(o xorm.Interface, id int64, fields ...string) (err error) {
	if o == nil {
		o = Engine
	}
	session:=o.Table("teacher2Grades" + strconv.FormatInt(id,10))	
	_, err =session.ID(u.ID).Cols(fields...).Update(u)
	//utils.CheckError(err)
	return
}

func (u *Teacher2Grades) DeleteAll(o xorm.Interface, id int64) (err error) {
	if o == nil {
		o = Engine
	}
//	session:=o.Table("teacher2Grades" + strconv.FormatInt(id,10))	
	err = Engine.DropTables("teacher2Grades" + strconv.FormatInt(id,10))
	//utils.CheckError(err)
	return
}

func (u *Teacher2Grades) Query(o xorm.Interface, id int64) (count int64){
	if o == nil {
		o = Engine
	}	
	session:=o.Table("teacher2Grades" + strconv.FormatInt(id,10))	
	count,_ = session.FindAndCount(u)
	return
}


func (u *Grades) Query(o xorm.Interface, id int64) (count int64){
	if o == nil {
		o = Engine
	}	
	session:=o.Table("grades" + strconv.FormatInt(id,10))	
	_,_ = session.Get(u)
	log.Println(u.ID)
	return
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
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/webcron?charset=utf8")
	if err != nil {
		log.Printf("new engine failed: %s\n", err)
		return
	}

	Engine.ShowSQL(true)

	grades := &Grades{Years: 3, Classes: 4}
	num := grades.Query(nil,1); 
	log.Println(num)

	return 
}
