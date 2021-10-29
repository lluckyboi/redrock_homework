package main

import "fmt"

func main() {
	over := make(chan bool,1)
	oee := make(chan bool,0)
	for i := 0; i < 10; i++ {
		/*go func() {
			fmt.Println(i)
		}()
		if i == 9 {
			over <- true
		}*/  //主协程和协程同时进行，却没有阻塞（只有当i=9时有）

		go func(){
			fmt.Println(i)
			<-oee
		}()
		oee<-true
		//保持步调一致
		if i == 9 {
			over <- true
		}

	}
	<-over
	fmt.Println("over!!!")
}