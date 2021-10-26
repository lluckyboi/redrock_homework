package main
import "fmt"
func main(){
	//数据输入
	Database := map[string]string{
		"admin":"admin",
		"zhangsan":"lisi",
		"xiaoming":"xiangwang",
		"npy":"nil",
	}
	pf := "请输入账号:"
	fmt.Println(pf)
	var pff string
	fmt.Scanf("%s",&pff)
	pw := "请输入密码:"
	fmt.Println(pw)
	var pww string
	fmt.Scanf("%s",&pww)


	///数据处理
	var un int = 0
	var ps int = 0
	for Uname,Pword := range Database{
		if ( pff == Uname && pww == Pword ){
			fmt.Println("登录成功!")
			break
		}
		if (pff != Uname ) {
			//用户名不存在！对应un
			un =un+1
		}
		if ( pww != Pword ){
			//密码不存在！对应ps
			ps =ps+1
		}
	}
	if(un==4){
		fmt.Println("用户不存在！")
	}else if(un!=4){
		if(ps==4){
			fmt.Println("密码错误！")
		}
	}
}
