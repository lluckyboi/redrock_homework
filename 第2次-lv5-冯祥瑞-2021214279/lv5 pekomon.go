//时间有限 坑挖着
package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"


)

type sqbable interface {
	Beseized()
	Appear()
	Battle()
	sing()
	//...
}
//皮卡丘 pika~
type pika struct {}
//可达鸭
type psyduck struct {}
//杰尼龟
type zeni struct {}

//pika实现接口
func (ba pika)Beseized(){
	 fmt.Println("你抓到了pikachu!")
}
//直接摆烂
func (ba pika)Battle(){
	 fmt.Println("你打败了pikachu!")
}
func (ba pika)sing(){
	fmt.Println("pika~pika~")
}
func (ba pika)Appera(){
	fmt.Println("一只pika出现了！")
}
//duck
func (ba psyduck)Beseized(){
	fmt.Println("你抓到了psyduck!")
}
func (ba psyduck)Battle(){
	fmt.Println("你打败了psyduck!")
}
func (ba psyduck)sing(){
	fmt.Println("ga~ga~")
}
func (ba psyduck)Appear(){
	fmt.Println("一只psyduck出现了！")
}

//zeni
func (ba zeni)Beseized(){
	fmt.Println("你抓到了zenigame!")
}
func (ba zeni)Battle(){
	fmt.Println("你打败了zenigame!")
}
func (ba zeni)sing(){
	fmt.Println("zeni~zeni~")
}
func (ba zeni)Appear(){
	fmt.Println("一只zeni出现了！")
}


func main() {
	//宝贝id表
	var form map[int]string = make(map[int]string)
	form[1] = "pikachu"
	form[2] = "psyduck"
	form[3] = "zenigame"

	var pikaa pika
	var zenii zeni
	var psyduckk psyduck
	var fs string = "/////////////////////////////////////////////\n" +
		"            欢迎来到pokem game!\n" +
		"			  双击Enter键进入游戏\n" +
		"//////////////////////////////////////////////\n"
	fmt.Println(fs)
	//一个时停

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	//开始啦
	//"真"随机
	// 生成 1 个 [0, 810) 范围的真随机数。
	//到目前 一共出现了809只Pokémon
	result, _ := rand.Int(rand.Reader, big.NewInt(810))
	idd := result.Int64() //类型由bigInt转化为int64
	idd = idd % 4 //取模 因为只有3个宝贝

	//int64转int
	//先转为golang string 再转为int
	//指针也可以实现
	strInt64 := strconv.FormatInt(idd, 10)
	id, _ := strconv.Atoi(strInt64)
	// 遇见pokemon  会是谁捏
	//直接爆搜好吧

	switch id {
	case 0:{fmt.Println("运气太差 没有pokemon 洗洗睡吧");os.Exit(1)}//原来是抽卡游戏
	case 1:
		pikaa.Appera()
	case 2:
		psyduckk.Appear()
	case 3:
		zenii.Appear()
	}
	fmt.Println("快输入B和它战斗！")
	bufio.NewReader(os.Stdin).ReadBytes('B')
	switch id {
	case 1:
		pikaa.Battle()
	case 2:
		psyduckk.Battle()
	case 3:
		zenii.Battle()
	} //屎山是什么
	fmt.Println("要不要抓它？按S试试？")
	var a byte
	fmt.Scanln(&a)
	//按啥都一样 我摆烂了
	switch id {
	case 1:
		pikaa.Beseized()
	case 2:
		psyduckk.Beseized()
	case 3:
		zenii.Beseized()
	}
	switch id {
	case 1:
		pikaa.sing()
	case 2:
		psyduckk.sing()
	case 3:
		zenii.sing()
	}
	fmt.Println("你的pokemon之旅结束啦！！！！\n" +
						"按Q退出游戏")
	bufio.NewReader(os.Stdin).ReadBytes('Q')

}