package main

import (
	"testing"
	"eddie-handy/sun"
	"fmt"
	"log"
	"eddie-handy/edd_log"
)

func TestGetTime(t *testing.T) {
	sunObject :=sun.NewObject(123.12,32.2,1.8)
	val, err :=sunObject.GetSunTime(sun.SUN_RISE)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(val.SunFormatTime)
	fmt.Println(val.SunStampTime)
}

func TestEdd(t *testing.T)  {
	edd_log.Edd()
}