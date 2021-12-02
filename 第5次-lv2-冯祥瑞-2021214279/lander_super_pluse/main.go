package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	//先进行账号密码管理  再启动WEB服务
	//虽然有点怪

	UserData := map[string]string{}
	//数据输入
	bytes2, _ := ioutil.ReadFile("./userdata")
	json.Unmarshal(bytes2, &UserData)

	for;;{
		fmt.Println("输入指令 N新建账号 R重设密码 开启服务按Q")
		var index string
		fmt.Scanln(&index)
		/*if index == "L" {
			pf := "请输入账号:"
			fmt.Println(pf)
			var pff string
			fmt.Scanf("%s", &pff)
			pw := "请输入密码:"
			fmt.Println(pw)
			var pww string
			fmt.Scanf("%s", &pww)

			///数据处理
			var un int = 0

			for Uname, Pword := range UserData {
				if (pff == Uname ) {
					isOk, _ := ValidatePassword(pww,Pword)
					if !isOk {
						fmt.Println("密码错误")
						un=1
						break
					}else {
						fmt.Println("登录成功!")
						un=1
					}
				}

			}
			if un==0 {
				fmt.Println("用户不存在！")
			}
		} else*/ if index == "N" {
			fmt.Println("输入账号")
			var username string
			fmt.Scanln(&username)
			ch :=bytecheck(username)
			if ch==1{
				fmt.Println("包含了特殊字符")
			}else {
				var re int = 1
				for Uname, _ := range UserData {

					if (Uname == username) {
						fmt.Println("请勿重复注册")
						re = 0
					}

				}
				if re == 1 {
					fmt.Println("请输入密码")
					var pswd string
					fmt.Scanln(&pswd)
					ch :=bytecheck(pswd)
					if ch==1{
						fmt.Println("包含了特殊字符")
					}else {
						passwordbyte, _ := GeneratePassword(pswd)
						UserData[username] = string(passwordbyte)
						fmt.Println("注册成功！")
					}
				}
			}
		} else if index == "R" {
			fmt.Println("输入要重设的账号")
			var ruse string
			var exi int = 0
			fmt.Scanln(&ruse)
			for Uname, _ := range UserData {
				if (Uname == ruse) {
					exi = 1
				}
			}
			if exi == 0 {
				fmt.Println("没有此账号")
			} else {
				fmt.Println("请输入密码")
				var rpass string
				fmt.Scanln(&rpass)
				ch :=bytecheck(rpass)
				if ch==1{
					fmt.Println("包含了特殊字符")
				}else{
					if len(rpass) < 6 {
						fmt.Println("请输入六位以上的密码！")
					} else {
						passwordbyte, _ := GeneratePassword(rpass)
						UserData[ruse] = string(passwordbyte)
						fmt.Println("重设成功！")
					}
				}
			}
		} else if index == "Q" {
			break
		}

	}
	//操作结束后进行文件写入
	//先清空文件
	os.OpenFile("./userdata", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	//再写入
	ioutil.ReadFile("./userdata")
	bytes, _ := json.Marshal(&UserData)
	ioutil.WriteFile("./userdata",bytes,02)

	//进行一个拿来主义
	//解序列化先
	bytes3, _ := ioutil.ReadFile("./userdata")
	json.Unmarshal(bytes3, &UserData)

	//开一个默认路由
	r := gin.Default()
	//获取cookie
	auth :=func(c *gin.Context){
		value,err :=c.Cookie("gin_cookie2")
		//错误处理
		if err!=nil{
			c.JSON(403,gin.H{
				"message":"认证失败 无cookie",
			})
			//终止后面该请求的所有进程
			c.Abort()
		}else{
			//获取的cookie写入上下文
			c.Set("cookie",value)
			//挂起来执行剩下进程
			c.Next()
		}
	}

	//收到post请求时 怎么处理
	r.POST("/login",func (c *gin.Context){
		//从表单读数据
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		pd :=0
		//判定
		for uname,pword:=range UserData {
			//密码对比
			isok,_:=ValidatePassword(password,pword)
			if username == uname &&isok{
				//设置cookie 设置之后 前面就可以读了(就有了登录状态信息）
				c.SetCookie("gin_cookie2", username, 3600, "/", "", false, true)
				//输出成功信息
				c.String(200, "%s login successfully!", username)
				pd=1
				}
			}
			if pd==0 {
				c.JSON(403, gin.H{
				"message": "认证失败,账号密码错误",
			})
		}
	})
	//最后还得把HOOK挂上去鉴权
	//收到get请求时 r.Get有不定参数HandlerFunc 所以可以放入auth
	//匿名函数实现了接口 所以直接用？
	r.GET("/hello",auth,func(c* gin.Context){
		//尝试读cookie 看有没有登录
		cookie,_:=c.Get("cookie")
		//转为string
		str:=cookie.(string)
		//以string形式输出
		c.String(200,"hello world"+str)
	})
	r.Run()
	bytes4, _ := json.Marshal(&UserData)
	ioutil.WriteFile("./userdata",bytes4,02)
}

func bytecheck(ss string)int{
	if strings.Index(ss,";")!=-1{
		fmt.Println("含非法字符")
		return 1
	}else if strings.Index(ss," ")!=-1{
		fmt.Println("含非法字符")
		return 1
	}else if strings.Index(ss,"/")!=-1 {
		return 1
	}
	return 0
}

//GeneratePassword 给密码就行加密操作
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//ValidatePassword 密码比对
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}