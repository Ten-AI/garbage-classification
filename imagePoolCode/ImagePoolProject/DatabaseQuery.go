package main
import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init()  {
}

func GroupUrlGet(Group string){
	o:=orm.NewOrm()
	var obj []interface{}
	_,err:=o.Raw("select imageID from imageLabelTable where labelName = ? and confidence=1 limit 10;",Group).QueryRows(&obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v",obj)
}