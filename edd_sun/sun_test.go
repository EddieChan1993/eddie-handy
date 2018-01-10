package edd_sun

import (
	"testing"
	"fmt"
)

func TestSun(t *testing.T) {
	sunObj :=NewObject(13.2,23.1,(+8))
	sunT,_ :=sunObj.GetSunTime(SUN_SET)
	fmt.Println(sunT)

}
