/**
	常用函数
 */
package com_func

import (
	"log"
	"strconv"
)

//string->float64
func StringToFloat64(str string) float64 {
	floatVal,err:=strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return floatVal
}

//string->int64
func StringToInt64(str string) int64 {
	intVal, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

//float64->int64
func Float64ToInt64(floatParam float64) int64 {
	str:=strconv.FormatFloat(floatParam, 'f', -1, 64)
	return StringToInt64(str)
}