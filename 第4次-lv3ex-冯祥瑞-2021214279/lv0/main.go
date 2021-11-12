package main
//一个client端
//发送数据的同时会格式化后的发送时间一同发送
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

//格式化发送时间的函数

func Nowftime() []byte {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	tt :=now.Format("2006-01-02 15:04:05.000 Mon Jan")
	return []byte(tt)
}


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")//建立连接 dial意思是拨号
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q或Q就退出
			return
		}
		_, err = conn.Write(Nowftime())// 发送时间
		_, err = conn.Write([]byte(inputInfo)) // 发送数据

		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))

	}
}
