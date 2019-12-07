package main


import (
	"github.com/astaxie/beego/httplib"
	"errors"
)
var Url ="http://127.0.0.1:8080/group1"
func init(){}

func FileUploadToDFS(imageID string,group string) (md5 string,err error){
	obj:=make(map[string]string,10)
	req:=httplib.Post(Url+"/upload")
	req.PostFile("file","./TempImage/"+imageID+".jpg")//注意不是全路径
	req.Param("output","json")
	req.Param("scene",group)
	req.Param("path","")
	req.ToJSON(&obj)
	return obj["md5"] ,nil
}
func FileDeleteFromDFS(md5 string,group string)(bool,error){
	var obj interface{}
	req:=httplib.Get(Url+"/delete?md5="+md5)
	req.ToJSON(&obj)
	stat:=obj.(map[string]interface{})["status"].(string)
	if stat=="ok"{
		return true ,nil
	}else if stat=="false"{
		return false ,	errors.New("File Already Delete!")
	}
	return false,nil
}
func FileStat()(num int,size int){
	var obj interface{}
	mp :=make(map[string]int,10)
	mp["a"]=1
	req:=httplib.Get("http://127.0.0.1:8080/group1/stat")
	req.ToJSON(&obj)
	fileNum:=int(obj.(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["fileCount"].(float64))
	fileSize:=int(obj.(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["totalSize"].(float64))
	return fileNum,fileSize
}
