package main

import (
	"fmt"
	"os"
	"log"
	"archive/zip"
	"io"
)

const (
	LOGFILEPATH = "d:\\zip.log"
)

func main(){
	logfile,err := os.OpenFile(LOGFILEPATH,os.O_CREATE|os.O_RDWR,0);
	if err!=nil {
		fmt.Println(err.Error());
		return;
	}
	defer logfile.Close();
	logger := log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile);
	if logger==nil {
		fmt.Println("logger init error");
	}
	r,err := zip.OpenReader("E:\\新建文本文档.zip");
	if err!=nil {
		logger.Fatal(err);
	}
	defer r.Close();
	for _,f := range r.File {
		fmt.Println("FileName : ",f.Name);
		rc,err := f.Open();
		if err!=nil {
			logger.Fatal(err);
		}
		_,err = io.CopyN(os.Stdout,rc,68); //打印文件内容
		if err!=nil {
			if err!=io.EOF {
				logger.Fatal(err);
			}
		}
	}
}
