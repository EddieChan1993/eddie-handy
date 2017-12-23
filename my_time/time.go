package my_time

import (
	"log"
	"time"
)

const (
	YMD_HIS="2006-01-02 15:04:05"
	YMD="2006-01-02"
)

//标准格式转为时间戳
func FormatToStamp(timeStamp int64,timeVal string) string {
	if timeVal!=YMD_HIS&&timeVal!=YMD {
		log.Fatalln("timeVal参数不在指定范围")
	}
		return time.Unix(timeStamp,0).Format(timeVal)
}

//时间戳转为标准格式
func StampToFormat(format string,timeVal string) int64 {
	if timeVal!=YMD_HIS&&timeVal!=YMD {
		log.Fatalln("timeVal参数不在指定范围")
	}

	loc,_:=time.LoadLocation("Local")//获取当地时区
	tm2,err :=time.ParseInLocation(timeVal,format,loc)
	if err!=nil {
		log.Fatalln(err)
	}
	return tm2.Unix()
}