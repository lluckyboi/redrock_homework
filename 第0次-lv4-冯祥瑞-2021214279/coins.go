package main

import (
	"fmt"
	"strings"
)

func main(){
	//初始化 人人为0
	pers :=map[string]int{
		"Matthew" : 0,
		"Sarah"	  : 0,
		"Augustus": 0,
		"Heidi"	  : 0,
		"Emilie"  : 0,
		"Peter"	  : 0,
		"Giana"	  : 0,
		"Elizabeth":0,
	}
	var i,t int = 0,0

	//遍历计算 暴力双循环 切开判断
	for na,_ := range pers{
		l :=len(na)
		n :=strings.Split(na,"")
		for i=0;i<l;i++{
			if(n[i]=="e"||n[i]=="E"){
				pers[na]=pers[na]+1
				t=t+1
			}
			if(n[i]=="i"||n[i]=="I"){
				pers[na] = pers[na] +2
				t=t+2
			}
			if(n[i]=="o"||n[i]=="O"){
				pers[na] = pers[na] +3
				t=t+3
			}
			if(n[i]=="u"||n[i]=="U"){
				pers[na] = pers[na] +4
				t=t+4
			}

		}
	}

	fmt.Println(pers)
	fmt.Println("总共：")
	fmt.Println(t)
}
