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

const (
	ChanSize int = 5
)

//图片下载函数，保存图片到TempImage文件夹下
func ImageDownload(imageID string ,url string,label string)(status bool,err error){
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false,errors.New("can't connected")
	}
	if resp.StatusCode==200{
		img,err:=os.OpenFile("./TempImage/"+label+"/"+imageID+".jpg",os.O_WRONLY|os.O_CREATE,0666)
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
			_,err = img.Write(image[:n])
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
func ImageDownload2(count chan int,url string,label string)(status bool,err error){
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false,errors.New("can't connected")
	}
	if resp.StatusCode==200{
		a:=<-count
		println(a,url)
		img,err:=os.OpenFile("./TempImage/"+label+"/"+strconv.Itoa(a)+".jpg",os.O_WRONLY|os.O_CREATE,0777)
		defer img.Close()
		if err != nil {
			println(err)
		}
		image:=make([]byte,256)
		for {
			n,err := resp.Body.Read(image)
			if err == io.EOF {
				break
			}
			_,err = img.Write(image[:n])
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
func ImageNetDownload(urls []string,label string)(int){
	println(len(urls))
	w:=sync.WaitGroup{}
	lock:=sync.Mutex{}
	sum :=9
	count:=make(chan int,ChanSize)
	for i:=0;i<ChanSize;i++{
		count<-i
	}
	for i:=0;i<len(urls);i++{
		w.Add(1)
		go func(i int,url string,imgnum *int) {
			url=strings.ReplaceAll(url,"\r","")
			flag,err:=ImageDownload2(count,url,label)
			if err != nil {
				fmt.Printf("%v\n",err)
			}
			if flag==true{
				lock.Lock()
				*imgnum++
				count<-*imgnum
				lock.Unlock()
			}
			w.Done()
		}(i,urls[i],&sum)
	}
	w.Wait()
	return sum
}