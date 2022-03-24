package main

import "fmt"

//向频道发布消息
func post(channame string, msg string) {
	ctx := "UBLISH " + channame + " " + msg
	res, err := rdb.Do(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res, "人收到信息")
}

//订阅频道
func addsub(subname string) {
	ctx := "PSUBSCRIBE " + subname
	res, err := rdb.Do(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("message:", res)
}

//退订
func unsub(subname string) {
	ctx := "UNSUBSCRIBE " + subname
	_, err := rdb.Do(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("退订成功")
}
