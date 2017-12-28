package main

import (
	"testing"
	."eddie-handy/edd_log"
	."eddie-handy/edd_curl"
	sun "eddie-handy/edd_sun"
	"fmt"
)

func TestSun(t *testing.T) {
	sunObj :=sun.NewObject(13.2,23.1,(+8))
	sunT,_ :=sunObj.GetSunTime(sun.SUN_SET)
	fmt.Println(sunT)

}

func TestLog(t *testing.T) {
	Config("one.log")
	Debugf("Debugf")
	Infof("Infof")
	Warnf("Warnf")
	Errorf("Errorf")
	Fatalf("Fatalf")

	aa:=LogFile{}
	aa.Write([]byte("what fox\n"))
}

func TestCurl(t *testing.T) {
	url:= "http://sms-api.luosimao.com/v1/send.json"

	headers:=map[string]string{
		"Content-Type":"application/x-www-form-urlencoded",
		"Authorization":BasicAuth("api","78aac6166f23182bd2eaceae0fba6aa84"),
	}
	postData:=map[string]string{
		"mobile":"18380591566",
		"message":"go-lang test【环球娃娃】",
	}

	req:=NewRequst(url)
	result:=req.
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	fmt.Println(result)
}