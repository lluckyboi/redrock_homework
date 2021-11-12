package main

import (
	"fmt"
	"time"
)



func main(){
	var wt int=0
	var hour,min int
	fmt.Println("想添加计时器吗？O for onetimer ,R for reuse timer ,N for no")
	var ss string="nil"
	fmt.Scanf("%s ",&ss)
	var hoo string
	if ss=="O"{
		fmt.Println("倒计时(D)还是闹钟(C)？")
		fmt.Scanln(&hoo)
		if hoo == "D"{
			fmt.Println("每小时会提醒你一次")
		}else if hoo == "C"{
			fmt.Println("按格式输入时间")
			fmt.Scanf("%d:%d",&hour,&min)
		}else{
			fmt.Println("no such option!")
		}

	}else if ss=="R"{
		fmt.Println("按格式输入时间")
		fmt.Scanf("%d:%d",&hour,&min)
	}else if ss=="N"{
		//无事发生
	}else{
		fmt.Println("no such option!")
	}

	fmt.Println("想要删除计时器吗？1 for Timer1, 2 for Timer2 , 3 for Timer , 4 for addtimerOD ,5 for addtimerOC" +
		"6 for addtimerR 999 for no")

	fmt.Scanln(&wt)

	ch := make(chan int)
	ch4 := make(chan int)
	//timer1
	go Timer1(wt)

	//timer2
	go Timer2(wt)

	//timer3
	go  Timer3(wt,ch)

	//如果添加过
	if ss=="O"&&hoo=="D"{
		go Timer4(wt,ch4)
	}else if ss=="O"&&hoo=="C" {
		go Timer5(wt,hour,min)
	}else if ss=="R"{
		go Timer6(wt,hour,min)
	}
	<-ch
	<-ch4
}


func Timer1(t int) {
	for ; ; {
		now := time.Now()
		now.Format("15:04")
		ho := now.Hour()
		mi := now.Minute()
		b := ho == 02 && mi == 00
		if b {
			fmt.Println("谁能比我卷！")
		}
		if t==3 {
			break
		}
	}
}

func Timer2(t int) {
	for ; ; {
		now := time.Now()
		now.Format("15:04")
		ho := now.Hour()
		mi := now.Minute()
		b := ho == 06 && mi == 00
		if b {
			fmt.Println("早八算什么，早六才是吾辈应起之时！")
		}
		if t==2 {
			break
		}
	}
}

func Timer3(t int,ch chan<-int){
	for ;;{
		ticker := time.Tick(30*time.Second)
		for range ticker {
			fmt.Println("芜湖！起飞！")
		}

		if t==1 {
			break
		}
	}
	ch<- 1
}

func Timer4(t int,ch4 chan<-int){
	for ;;{
		ticker := time.Tick(time.Hour)
		for range ticker {
			fmt.Println("嘀嘀嘀！")
		}

		if t==4 {
			break
		}
	}
	ch4<- 1
}

func Timer5(t int,hour int,min int) {
	for ; ; {
		now := time.Now()
		now.Format("15:04")
		ho := now.Hour()
		mi := now.Minute()
		b := ho == hour && mi == min
		if b {
			fmt.Println("到点儿啦！")
			break
		}
		if t==5 {
			break
		}
	}
}

func Timer6(t int,hour int,min int) {
	for ; ; {
		now := time.Now()
		now.Format("15:04")
		ho := now.Hour()
		mi := now.Minute()
		b := ho == hour && mi == min
		if b {
			fmt.Println("到点儿啦！")
			break
		}
		if t==6 {
			break
		}
	}
}