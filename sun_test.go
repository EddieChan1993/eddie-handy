package main

import (
	"testing"
	."eddie-handy/edd_log"
	"eddie-handy/sun"
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
	aa.Write([]byte("what fox"))
}

