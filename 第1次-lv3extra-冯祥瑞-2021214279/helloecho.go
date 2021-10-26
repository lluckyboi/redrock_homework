package main

import (
	"fmt"
	"strings"
)

func main(){
	var b string
	var n[]string
	fmt.Scanf("%s",&b)
	n=strings.Split(b,"")//切开
	for a:=len(n)-1;a>=0;a--{
		fmt.Printf("%s",n[a])//逆序打印
	}

}

