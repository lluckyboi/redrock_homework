package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("input:数字个数")//

	var num,i int
	var a[999]int

	fmt.Scanf("%d",&num)
	for i=0;i<num;i=i+1{
			fmt.Scanf("%d",&a[i])
	}
	b := a[0:num]
	sort.Ints(b)
	fmt.Println("output")
	for i=0;i<num;i=i+1{
		fmt.Printf("%d ",b[i])
	}
}


