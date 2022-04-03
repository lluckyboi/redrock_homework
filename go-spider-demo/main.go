package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func savToFile(index int, data [][]string) {
	f, err := os.Create("第" + strconv.Itoa(index) + "页.txt")
	if err != nil {
		fmt.Println("os create err", err)
		return
	}
	defer f.Close()
	// 查出有多少条
	n := len(data)
	for i := 1; i < n; i++ {
		m := len(data[i])
		for j := 1; j < m; j++ {
			f.WriteString(data[i][j])
		}
		f.WriteString("\n\n")
	}
}

func main() {
	var start, end int
	fmt.Print("请输入要爬取的起始页")
	fmt.Scan(&start)
	fmt.Print("请输入要爬取的终止页")
	fmt.Scan(&end)
	working(start, end)
}

func working(start int, end int) {
	fmt.Printf("正在爬取%d到%d页", start, end)
	for i := start; i <= end; i++ {
		SpiderPage(i)
	}
}

// 爬取一个页面数据信息保存到文档
func SpiderPage(index int) {
	// 获取url
	url := "http://xiaodiaodaya.cn/article/view.aspx?id=" + strconv.Itoa(index)

	// 爬取url对应页面
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("httpget err", err)
		return
	}
	fmt.Println("result=", result)
	ret2 := regexp.MustCompile("([1-9][0-9]{0,1}、)?[^(\"开心一刻,笑话大全,爆笑笑话,冷笑话,笑掉大牙\")(，笑话大全，爆笑冷笑话精选，经典笑话尽在笑掉大牙！海量笑话内容每日定时更新，给您带来更多快乐。\" )]([\\d]{0,3}[\u4e00-\u9fa5][^a-zA-Z<>/0-9]+)+")
	data := ret2.FindAllStringSubmatch(result, -1)

	savToFile(index, data)

}

// 爬取指定url页面，返回result
func HttpGet(url string) (result string, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	// 设置头部信息
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 OPR/66.0.3515.115")
	resp, err1 := (&http.Client{}).Do(req)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	//循环爬取整页数据
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
