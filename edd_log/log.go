package edd_log

import (
	"log"
	"os"
	"time"
	"fmt"
	"io"
)

type Logger struct {
	info *log.Logger //消息信息
	waring *log.Logger //警告信息
	error *log.Logger //严重错误

	folderName string //文件夹名称
}

//实例化
func NewLogger(folderName string) *Logger {
	return &Logger{
		folderName:folderName,
	}
}

//重置文件夹路径
func (this *Logger)SetFolderName(folderName string)  {
	this.folderName=folderName
}

//初始化文件路径
func (this *Logger) initLogFile(prefix string)  {
	if this==nil {
		log.Fatalln("当前对象尚未实例化")
	}

	folder:=fmt.Sprintf("./%s/%s",this.folderName,time.Now().Format("20060102"))
	file:=fmt.Sprintf("%s/%s.log",folder,time.Now().Format("20060102150405"))

	err:=os.MkdirAll(folder,os.ModePerm)
	if err!= nil {
		log.Fatalln(err)
	}

	f,err:=os.OpenFile(file,os.O_CREATE|os.O_APPEND,os.ModePerm)
	if err!= nil {
		log.Fatalln(err)
	}

	log.SetPrefix(prefix)
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(io.MultiWriter(f,os.Stderr))
}

//Info打印输出
func (this *Logger) Info(format string, args ...interface{}) {
	this.initLogFile("[INFO]")
	log.Output(2,fmt.Sprintf(format,args...))
}

//Warn打印输出
func (this *Logger)Warn(format string,args ...interface{})  {
	this.initLogFile("[WARNING]")
	log.Output(2,fmt.Sprintf(format,args...))
}

//Error打印输出
func (this *Logger) Error(format string, args ...interface{}) {
	this.initLogFile("[ERROR]")
	log.Output(2,fmt.Sprintf(format,args...))
}