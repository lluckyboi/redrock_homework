//爷的排序
package main

import (
	"fmt"
	"sort"
)
//排序接口
type ssort interface {
	mysort()
}
//排序方法（只实现了int切片）
func (arr *slicee)mysort(){
	  sort.Ints(arr.val)
}
//对象
type slicee struct {
	val []int
}
//进行一个坑的挖
type mapp struct {}

func main(){
	 var n,i int
	 var mys slicee
	 //格式化输入
	 fmt.Scanf("%d",&n)
	 for i=0;i<n;i++{
		 var s int
		 fmt.Scanf("%d",&s)
		 mys.val=append(mys.val,s)
	 }
	 //进行接口的调用
	mys.mysort()
	fmt.Println(mys)
}