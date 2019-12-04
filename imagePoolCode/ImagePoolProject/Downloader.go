package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/astaxie/beego/httplib"
	"strconv"
	"strings"
	"sync"
)
//图片下载函数，保存图片到TempImage文件夹下
func ImageDownload(imageID string ,url string )(status bool,err error){
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false,errors.New("can't connected")
	}
	if resp.StatusCode==200{
		img,err:=os.OpenFile("./TempImage/"+imageID+".jpg",os.O_WRONLY|os.O_CREATE|os.O_EXCL,0666)
		defer img.Close()
		if err != nil {
			println(err)
		}
		image:=make([]byte,10240)
		for {
			n,err := resp.Body.Read(image)
			if err == io.EOF {
				break
			}
			n,err = img.Write(image[:n])
			if err != nil {
				fmt.Println("img write:",err)
			}
		}
		return true ,nil
	} else{
		err:=errors.New("StatusCode Error")
		return false,err
	}
}
func urlDownload(url string)[]string{

	resp:=httplib.Get(url)
	text,err:=resp.String()
	if err != nil {
		fmt.Println(err)
	}
	urls:=strings.Split(text,"\n")
	return urls
}
func ImageNetDownload(urls []string){
	println(len(urls))
	w:=sync.WaitGroup{}

	for i:=0;i<len(urls);i++{
		w.Add(1)
		println(i)
		go func(i int,url string) {
			url=strings.ReplaceAll(url,"\r","")
			ImageDownload(strconv.Itoa(i),url)
			w.Done()
		}(i,urls[i])
	}
	w.Wait()
}