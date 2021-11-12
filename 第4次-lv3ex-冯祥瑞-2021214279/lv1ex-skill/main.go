package main

import (
	"bytes"
	"fmt"
	"os"
)
var banneddata =[]string{"ddz","fufu","卷卷","改造计划","开卷！"}
var qq string
func main() {
	for ;; {
		//找不到技能时会报错 接收错误
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("没有这个技能哦")
			}
		}()
		// 定义一个字符串映射到func()的map,然后填充这个map
		var skill = map[string]func(string){
			"蝶火燎原": func(ss string) {
				fmt.Printf("%s", ss)
			},
			"裁雨留虹": func(ss string) {
				fmt.Printf("%s", ss)
			},
			"左右横跳": func(ss string) {
				fmt.Printf("%s", ss)
			},
		}
		fmt.Println("\n本次游戏是否新添加临时技能？Y/N?	")
		var q string
		fmt.Scanln(&q)
		if q == "Y" {
			var newsk string
			fmt.Println("输入你要新添加的技能")
			fmt.Scanln(&newsk)
			for i := range banneddata {
				ss := bytes.Index([]byte(newsk), []byte(banneddata[i]))
				if ss != -1 {
					fmt.Println("有敏感词哦!")
					os.Exit(1)
				}
			}
			skill[newsk] = func(ss string) {
				fmt.Printf("%s", ss)
			}
		}

		var sname string
		fmt.Println("请输入你要释放的技能名")
		fmt.Scanln(&sname)
		for i := range banneddata {
			ss := bytes.Index([]byte(sname), []byte(banneddata[i]))
			if ss != -1 {
				fmt.Println("有敏感词哦!")
				os.Exit(1)
			}
		}

		f := skill[sname]

		fmt.Println("是否要自定义描述？Y/N?")
		var zdy string
		fmt.Scanln(&zdy)

		var descri [999]string
		var sktimes int = 1

		if zdy == "Y" {
			fmt.Println("技能描述出现几次？")
			fmt.Scanln(&sktimes)
			for j := 1; j <= sktimes; j++ {
				fmt.Printf("输入你的第%d个模板", j)
				fmt.Scanf("%s", &descri[j])
			}
			for j := 1; j <= sktimes; j++ {
				ReleaseSkill(sname, func(skillName string) {
					f(descri[j])
					fmt.Printf("%s", sname)
				})
			}
		} else {
			for j := 1; j <= sktimes; j++ {
				ReleaseSkill(sname, func(skillName string) {
					f(descri[j])
					fmt.Printf("%s", sname)
				})
			}
		}
		fmt.Println("\n结束游戏？Y/N?")
		fmt.Scanln(&qq)
		if qq=="Y"{
			os.Exit(0)
		}else{
			continue
		}
	}
}


func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}