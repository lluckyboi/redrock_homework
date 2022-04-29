
package main

import (
	"RpcDemo/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := proto.NewUserClient(conn)
	for {
		fmt.Println("input chose,1 for login ,2 for register")
		cs:=0
		_, _ = fmt.Scanln(&cs)
		if cs==1{
			iptName := ""
			iptPassword := ""
			fmt.Scanln(&iptName,&iptPassword)
			fmt.Println(iptName,iptPassword)

			loginResp, _ := cl.Login(context.Background(), &proto.LoginReq{
				UserName: iptName,
				PassWord: iptPassword,
			})
			fmt.Println(loginResp)
			if loginResp.OK {
				fmt.Println("success")
				break
			}
			fmt.Println("retry")
		} else if cs==2{
			iptName := ""
			iptPassword := ""
			q1:=""
			q2:=""
			a1:=""
			a2:=""
			fmt.Scanln(&iptName,&iptPassword,&q1,&q2,&a1,&a2)
			loginResp, _ := cl.Newuser(context.Background(), &proto.LoginReq{
				UserName: iptName,
				PassWord: iptPassword,
				Ques1: q1,
				Ques2: q2,
				Ans1: a1,
				Ans2: a2,
			})
			if loginResp.OK {
				fmt.Println("success")
				break
			}
			fmt.Println("retry")
		}
	}
}

