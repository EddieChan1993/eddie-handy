/**
	时间的转换
 */
package edd_time

import (
	"log"
	"time"
)

const (
	//timeVal
	YMD_HIS="2006-01-02 15:04:05"
	YMD="2006-01-02"
)

//时间戳转为标准格式
func FormatToStamp(timeStamp int64, timeFormat string) string {
	if timeFormat !=YMD_HIS&& timeFormat !=YMD {
		log.Fatalln("timeVal参数不在指定范围")
	}
	return time.Unix(timeStamp,0).Format(timeFormat)
}

//标准格式转为时间戳
func StampToFormat(format string, timeFormat string) int64 {
	if timeFormat !=YMD_HIS&& timeFormat !=YMD {
		log.Fatalln("timeVal参数不在指定范围")
	}

	loc,_:=time.LoadLocation("Local")//获取当地时区
	tm2,err :=time.ParseInLocation(timeFormat,format,loc)
	if err!=nil {
		log.Fatalln(err)
	}
	return tm2.Unix()
}