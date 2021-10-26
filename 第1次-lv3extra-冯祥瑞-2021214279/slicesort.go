package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var s[]int
	var m[]int
	//task1
	for i:=0; i<100; i++ {
		s=append(s,rand.Intn(100))
	}
	//task2
	sort.Ints(s)
	fmt.Println(s)
	//task3
	for i:=0; i<100; i++ {
		m=append(m,rand.Intn(100))
	}
	BubbleSort(&m)
	fmt.Println(m)
}
//冒泡排序
func BubbleSort (ss *[] int){
	var i,j int
	for i = 0; i < 100;i++ {
		for j = 0; j+1 < 100-i; j++{
			if (*ss)[j] < (*ss)[j+1] {
				(*ss)[j], (*ss)[j+1] = (*ss)[j+1], (*ss)[j]
			}
		}
	}

}