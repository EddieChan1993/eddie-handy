package edd_func

//删除指定元素
func ValRemove(slice []interface{},value interface{})[]interface{}  {
	for i:=0;i<len(slice);i++ {
		if value == slice[i] {
			slice=KeyRemove(slice,i)
			//slice的某元素在被删除后，后面的元素向前移，然后而索引也应该移向前一位
			//避免错过元素
			i=i-1
		}
	}
	return slice
}

//根据索引删除
func KeyRemove(slice []interface{}, i int) []interface{} {
	//    copy(slice[i:], slice[i+1:])
	//    return slice[:len(slice)-1]
	return append(slice[:i], slice[i+1:]...)
}

//插入到指定索引位置
func Insert(slice *[]interface{}, index int, value interface{}) {
	rear := append([]interface{}{}, (*slice)[index:]...)
	*slice = append(append((*slice)[:index], value), rear...)
}
