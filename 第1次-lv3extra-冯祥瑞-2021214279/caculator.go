package main
import (
	"fmt"
)
func main() {
	var c byte
	var a, b,fl int
	fl=1
	//格式化输入
	for{
	fmt.Scanf("%d%c%d\n", &a, &c, &b)
	//switch进行运算符判断
	switch c {
	case '+':
		fmt.Printf("%d\n", a+b)
	case '-':
		fmt.Printf("%d\n", a-b)

	case '*':
		fmt.Printf("%d\n", a*b)
	case '/':
		if (b != 0) {
			fmt.Printf("%d\n", a/b)
		} else {
			fmt.Println("Error")
			break
		}
	}

	ff := "继续请按1 结束按0"
		fmt.Println(ff)

		fmt.Scanf("%d",&fl)
		if fl==0{
			break
		}else{
			continue
		}
}
}
