package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//task 1
	os.Create("plan.txt")
	//task 2
	//file.Write不知道为什么用不了 file飘红
	//只好用ioutil.WriteFile实现了
	mee :=[]byte("I’m not afraid of difficulties and insist on learning programming")
	err := ioutil.WriteFile("plan.txt", mee,0644)
	if err != nil {
		return 
	}
	//task 3
	//完了 file.Read也没有补全
	b,_ :=ioutil.ReadFile("plan.txt")
	_,_ =fmt.Fprintln(os.Stderr,string(b))//把b强转为string

}
