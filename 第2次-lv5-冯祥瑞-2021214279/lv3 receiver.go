package main

import (
	"fmt"
)

//函数在此

func Receiver(v interface{})  {
	switch v.(type){
	case int: fmt.Println("这个是int")
	case bool:fmt.Println("这个是bool")
	case string:fmt.Println("这个是string")
	case CAT:fmt.Println("这个是猫猫")
}
}

//空接口来啦
type reee interface{}
type CAT struct {}
func main(){
	var rec reee

	//fmt.Scan(&rec)	读不进
	rec=CAT{}
	Receiver(rec)
}
