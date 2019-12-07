package main

import (
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"strings"
)

type ModelInit struct {
	Id   int
}

//func init(){
//	orm.RegisterDataBase("default", "mysql", "root:116118@tcp(www.whatdoyoudo.club:3306)/dataset?charset=utf8", 30)
//	// create table
//	orm.RegisterModel(new(ModelInit))
//	orm.RunSyncdb("default", false, true)
//}

func ImageNetDownloader(){
	arg := os.Args
	label := arg[1]
	pwd,err:= os.Getwd()
	if err != nil {

	}
	println(pwd)
	err=os.MkdirAll(pwd+"TempImage",0777)
	if err != nil {
		println(err)
	}
	err=os.MkdirAll(pwd+"/TempImage/"+label,0777)
	if err != nil {
		println(err)
	}
	urls :=urlDownload("http://image-net.org/api/text/imagenet.synset.geturls?wnid="+label)
	n:=ImageNetDownload(urls,label)
	println("下载",n-10,"条"+label+"图像")
}

func ImageNetDownloader2(){
	arg := os.Args
	label := arg[1]
	pwd,err:= os.Getwd()
	if err != nil {

	}
	println(pwd)
	err=os.MkdirAll(pwd+"TempImage",0777)
	if err != nil {
		println(err)
	}
	err=os.MkdirAll(pwd+"/TempImage/"+label,0777)
	if err != nil {
		println(err)
	}
	urls :=urlDownload("http://image-net.org/api/text/imagenet.synset.geturls?wnid="+label)
	for i,k:=range(urls){
		println(i,k)
		k=strings.ReplaceAll(k,"\r","")
		ImageDownload(strconv.Itoa(i),k,label)
	}
}

func main(){
	ImageNetDownloader()
}