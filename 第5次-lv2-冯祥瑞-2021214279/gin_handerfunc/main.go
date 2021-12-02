package main

import (
	"github.com/gin-gonic/gin"
)
func main() {

	r := gin.Default()
	//获取cookie
	auth :=func(c* gin.Context){
		value,err :=c.Cookie("gin_cookie")
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
		//判定
		if username == "ddz" && password == "zzd"{
			//设置cookie 设置之后 前面就可以读了(就有了登录状态信息）
			c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
			//输出成功信息
			c.String(200,"%s login successfully!",username)
		}else{
			c.JSON(403,gin.H{
				"message":"认证失败,账号密码错误",
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
}
