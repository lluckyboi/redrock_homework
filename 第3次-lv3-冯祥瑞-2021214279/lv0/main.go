package main

import (
	"fmt"
)

var x int64
var ch =make(chan int64)
var cb =make(chan int64)

func main() {
	go func (){
		for i :=0;i<50000;i++{
			x=x+1
		}
		cb<-1
	}()
	go func (){
		<-cb
		for i :=0;i<50000;i++{
			x=x+1
		}
		<-ch
	}()
	//一个一个执行
	ch<-1
	fmt.Println(x)
}