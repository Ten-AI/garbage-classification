package main

import (
	_ "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

type ModelInit struct {
	Id   int
}

func init(){
	orm.RegisterDataBase("default", "mysql", "root:116118@tcp(www.whatdoyoudo.club:3306)/dataset?charset=utf8", 30)
	// create table
	orm.RegisterModel(new(ModelInit))
	orm.RunSyncdb("default", false, true)
}

func main()  {
	GroupUrlGet("/m/014j1m")
}

