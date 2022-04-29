package main

import (
	"RpcDemo/proto"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)
var Db *sql.DB

type User struct {
	Id       int
	Username 		string
	Password 		string
	Ques1 	 		string//密保1
	Ans1 	 		string//答案1
	Ques2	 		string//密保2
	Ans2	 		string//答案2
	Introduction	string
	PictureUrl		string
}

func main() {
	// 监听端口
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer() //获取新服务示例
		proto.RegisterUserServer(s, &server{})
		fmt.Println(s)
		InitDb()
		// 开始处理
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println(s)
}

type server struct {
	proto.UnimplementedUserServer // 用于实现proto包里BiliServer接口
}

func InitDb(){
	db, err := sql.Open("mysql", "root:WADX750202@/user")
	if err != nil {
		log.Fatal(err)
	}
	Db=db
}

func (s *server)Login(context context.Context,req *proto.LoginReq)(*proto.LoginResp, error){
	resp := &proto.LoginResp{}
	log.Println("recv:", req.UserName, req.PassWord)
	if req.PassWord != GetPassWord(req.UserName) {
		resp.OK = false
		return resp, nil
	}
	resp.OK = true
	fmt.Println("OKK")
	return resp,nil
}

func GetPassWord(usn string)string{
	user,err:=SelectQuesAndPasswordByUsername(usn)
	if err!=nil{
		log.Println(err)
		return ""
	}
	return user.Password
}

func SelectQuesAndPasswordByUsername(username string)(User,error){
	user :=User{}
	sqlstr :="select ans1,ans2,password from user where username=?"
	errs :=Db.QueryRow(sqlstr,username)
	//Err provides a way for wrapping packages to check for query errors without calling Scan.
	if errs.Err() != nil {
		return user, errs.Err()
	}
	//扫进user
	err := errs.Scan(&user.Ans1,&user.Ans2,&user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *server)Newuser(ctx context.Context, req *proto.LoginReq)(*proto.LoginResp, error){
	resp := &proto.LoginResp{}
	log.Println("recv:", req.UserName, req.PassWord,req.Ques1,req.Ques2,req.Ans1,req.Ans2)

	err,fl:=LengthCheck(req.PassWord)
	if(!fl){
		log.Println(err)
		resp.OK=false
		return resp,nil
	}
	err, fl = LengthCheck(req.UserName)
	if(!fl){
		log.Println(err)
		resp.OK=false
		return resp,nil
	}

	user :=User{
		Username :  req.UserName,
		Password :  req.PassWord,
		Ques1    :  req.Ques1,
		Ans1 	 :  req.Ans1,
		Ques2	 :  req.Ques2,
		Ans2	 :  req.Ans2,
	}
	isok,errr:= IsUsernameRepeat(req.UserName)
	if errr!=nil{
		log.Println("judge username repeat err: ", errr)
		resp.OK=false
		return resp,errr
	}
	//重复了
	if isok!=true{
		log.Println("username repeated")
		resp.OK=false
		return resp,nil
	}

	errrr := NewUser(user)
	if errrr!=nil{
		log.Println("register err: ", err)
		resp.OK=false
		return resp,errrr
	}
	resp.OK=true
	return resp,nil
}

func IsUsernameRepeat(username string)(bool,error){
	_,err:=SelectQuesAndPasswordByUsername(username)
	if err!=nil&&err!=sql.ErrNoRows{
		return false,err
	}
	return true,nil
}

func LengthCheck(ss string)(string,bool){
	if len(ss)>20||len(ss)<2{
		err:="长度错误！"
		return err,false
	}
	return "",true
}

func NewUser(user User)(error){
	sqlstr:="insert into user(username,password,ques1,ques2,ans1,ans2)values(?,?,?,?,?,?);"
	_, errs :=Db.Exec(sqlstr,user.Username,user.Password,user.Ques1,user.Ques2,user.Ans1,user.Ans2)
	if errs!=nil{
		return errs
	}
	return nil
}