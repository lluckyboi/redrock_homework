package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	UserData := map[string]string{}
	//数据输入
	bytes2, _ := ioutil.ReadFile("./userdata")
	json.Unmarshal(bytes2, &UserData)

for;;{
	fmt.Println("输入指令 L登陆 N新建账号 R重设密码 退出按Q")
	var index string
	fmt.Scanln(&index)
	if index == "L" {
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
	} else if index == "N" {
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