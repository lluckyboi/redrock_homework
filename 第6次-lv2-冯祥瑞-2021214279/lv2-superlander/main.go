package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func main() {

	r :=gin.Default()
	//抽象数据库
	db, err := sql.Open("mysql", "root:WADX750202@/user")
	if err != nil {
		log.Fatal(err)
	}

	type user struct {
		id       int
		username string
		password string
		ques1 	 string//密保1
		ans1 	 string//答案1
		ques2	 string//密保2
		ans2	 string//答案2
	}
	//首次运行先建表
	//sqlStr := "CREATE TABLE `user` (`id` BIGINT(20) NOT NULL AUTO_INCREMENT,`username` VARCHAR(20) DEFAULT '',`password` VARCHAR(20) DEFAULT '',`ques1` VARCHAR(20) DEFAULT '',`ques2` VARCHAR(20) DEFAULT '',`ans1` VARCHAR(20) DEFAULT '',`ans2` VARCHAR(20) DEFAULT '',PRIMARY KEY(`id`))ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
	//_, err = db.Exec(sqlStr)
	//if err != nil {
	//	fmt.Printf("failed, err:%v\n", err)
	//	return
	//}

		//cookie中间件
		auth :=func(c* gin.Context){
			value,err :=c.Cookie("gin_cookie")
			//错误处理
			if err!=nil{
			}else{
				//获取的cookie写入上下文
				c.Set("cookie",value)
				//挂起来执行剩下进程
				c.Next()
			}
		}



		//登录
		r.POST("/login",auth,func(c *gin.Context){
				pff:=c.PostForm("username")
				pww:=c.PostForm("password")
				sqlstr := "select *from user where username=?"
				var u user
				errL := db.QueryRow(sqlstr, pff).Scan(&u.id,&u.username, &u.password,&u.ques1,&u.ques2,&u.ans1,&u.ans2)
				if errL != nil {
					c.String(200,"没有此账号")
				}
				if (pff == u.username && pww == u.password) {
					c.SetCookie("gin_cookie", pff, 3600, "/", "", false, true)

					cookie,err:=c.Get("cookie")
					//首次登录无cookie
					if err==false{
						c.String(200,"登录成功!"+pff)
					}
					//后续登录时
					//转为string
					str:=cookie.(string)
					//以string形式输出
					c.String(200,"登陆成功!"+str)
				} else {
					c.String(200,"密码错误")
				}
		})


		//注册
		r.POST("/newbuild",func(c *gin.Context){
			var username string
			var password string
			var ques1 string
			var ques2 string
			var ans1 string
			var ans2 string

			sqlstr:="select *from user where username=?"

			var u user

			username=c.PostForm("username")
			password=c.PostForm("password")
			ques1=c.PostForm("ques1")
			ques2=c.PostForm("ques2")
			ans1=c.PostForm("ans1")
			ans2=c.PostForm("ans2")

			errn := db.QueryRow(sqlstr,username).Scan(&u.id,&u.username, &u.password,&u.ques1,&u.ques2,&u.ans1,&u.ans2)
			//如果能找到
			if (errn==nil) {
				c.String(200,"请勿重复注册！")
			}else {
				sqlstr:="insert into user(username,password,ques1,ques2,ans1,ans2)values(?,?,?,?,?,?);"
				ret, errs :=db.Exec(sqlstr,username,password,ques1,ques2,ans1,ans2)
				if errs != nil {
					c.String(200,"insert failed, err:%v\n", err)
					return
				}else {
					theID, _ := ret.LastInsertId() // 新插入数据的id
					c.String(200,"注册成功！ID:%d\n", theID)
				}
			}
		})


		r.POST("/refindpassword",func(c *gin.Context){
			var u user
			var username string
			var ans1 string
			var ans2 string

			username=c.PostForm("username")
			ans1=c.PostForm("ans1")
			ans2=c.PostForm("ans2")

			sqlstr:="select *from user where username=?"

			errr := db.QueryRow(sqlstr,username).Scan(&u.id,&u.username, &u.password,&u.ques1,&u.ques2,&u.ans1,&u.ans2)
			if errr!=nil {
				c.String(200, "没有该账户噢！")
			}else{
				if ans1==u.ans1&&ans2==u.ans2{
					c.String(200,"你的密码是:%s",u.password)
				}else {
					c.String(200,"有误")
				}
			}
		})

	r.Run()
}











	/*	 else if index == "R" {
			fmt.Println("输入要找回的账号")
			var rusename string
			fmt.Scanln(&rusename)
			sqlstr := "select *from user where username=?"
			var u user
			errr := db.QueryRow(sqlstr,rusename).Scan(&u.id,&u.username, &u.password,&u.ques1,&u.ques2,&u.ans1,&u.ans2)
			if errr != nil {
				fmt.Println("没有此账号:%v",errr)
			} else {
						var wnum int
						var answer string
						fmt.Println("请回答密保问题")
						fmt.Println("问题1：",u.ques1)
						fmt.Println("问题2：",u.ques2)
						fmt.Println("你选择回答？")
						fmt.Scanln(&wnum)
						fmt.Scanln(&answer)
						if wnum==1{
							if answer==u.ans1{
								fmt.Println("你的密码是",u.password)
							}else {
								fmt.Println("答案错误！")
							}
						}else{
							if answer==u.ans2{
								fmt.Println("你的密码是",u.password)
							}else {
								fmt.Println("答案错误！")
							}
						}

					}


		} else if index == "Q" {
			break
		}

}*/
//func bytecheck(ss string)int{
//	if strings.Index(ss,";")!=-1{
//		fmt.Println("含非法字符")
//		return 1
//	}else if strings.Index(ss," ")!=-1{
//		fmt.Println("含非法字符")
//		return 1
//	}else if strings.Index(ss,"/")!=-1 {
//		return 1
//	}
//	return 0
//}
