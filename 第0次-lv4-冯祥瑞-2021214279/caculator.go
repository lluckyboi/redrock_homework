package main
import (
	"fmt"
)
func main() {
	var c byte
	var a, b int
	//格式化输入
	fmt.Scanf("%d\n%c\n%d", &a, &c, &b)
	//switch进行运算符判断
	switch c {
	case '+':
		fmt.Printf("%d", a+b)
	case '-':
		fmt.Printf("%d", a-b)

	case '*':
		fmt.Printf("%d",a*b)
	case '/':
		if(b!=0){
			fmt.Printf("%d",a/b)
		}else{
			fmt.Println("Error")
		}

}
}
