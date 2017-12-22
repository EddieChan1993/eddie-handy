/**
	根据经纬度获取太阳落山的时间
 */
package sun

import (
	"time"
	"math"
	"fmt"
	"eddie-handy/com_func"
	"errors"
)

var All float64

const (
	SUN_RISE=1//日出
	SUN_SET=2//日落
)

type object struct {
	longitude float64//经度
	latitude float64//纬度
	zone int //时区

	h_m_s float64
	err error
	sunTime
}


type sunTime struct {
	SunFormatTime string
	SunStampTime  int64
}

func NewObject(lit,lat float64,zone int)*object{
	return &object{
		longitude:lit,
		latitude:lat,
		zone:zone,
	}
}

//获取hour min sec
func (this *object)getMHS()  {
	timeNow:=time.Now()
	yearDay:=timeNow.YearDay()

	y1:=10547*math.Pi/81000*math.Cos(float64(2)*math.Pi*float64(yearDay+9)/float64(365))
	y2:=this.latitude*math.Pi/180
	x1:=math.Tan(y1)*math.Tan(y2)
	x:=math.Acos(x1)*180/math.Pi

	this.h_m_s=x
}

func (this *object)thinkTime(all float64)  {
	//截取小数点后两位
	AllTime :=fmt.Sprintf("%.2f",All)

	Time :=com_func.StringToFloat64(AllTime)
	//获取时分秒
	hour:=math.Floor(Time)
	min:=math.Floor((Time - hour)*60)
	sec:=math.Floor(((Time-hour)*60-min)*60)
	upTimeFormat:=time.Now().Format("2006-01-02")

	hInt:=com_func.Float64ToInt64(hour)
	mInt:=com_func.Float64ToInt64(min)
	sInt:=com_func.Float64ToInt64(sec)

	//拼接完整时间格式
	upTimeFormat =fmt.Sprintf("%s %02d:%02d:%02d",upTimeFormat,hInt,mInt,sInt)

	//返回Time结构体
	tm2,err :=time.Parse("2006-01-02 15:04:05",upTimeFormat)
	if err != nil{
		this.err=err
	}


	sumT:= sunTime{
		SunFormatTime: upTimeFormat,
		SunStampTime:  tm2.Unix(),
	}

	this.sunTime=sumT
}

func (this *object) GetSunTime(sun_flag int) (sunT sunTime,err error) {
	this.getMHS()

	if sun_flag==SUN_RISE {
		All =(180+(+8)*15- this.longitude-this.h_m_s)/ 15
		this.thinkTime(All)
	}else if sun_flag==SUN_SET {
		All=(180+(+8)*15-this.longitude+this.h_m_s)/ 15
		this.thinkTime(All)
	}else {
		this.err=errors.New("sun_flag参数传入错误")
	}

	return this.sunTime,this.err
}



