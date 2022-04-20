package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	pongWait       = 60 * time.Second
	maxMessageSize = 512
	pingPeriod     = (pongWait * 9) / 10
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var count int
var ctchan = make(chan int)
var MsgCh = make(chan Msg, 50)
var mplock sync.Mutex
var mlock sync.Mutex
var room = make(map[string]*websocket.Conn)

type Msg struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
	Content  string    `json:"content"`
}

func Rec(username string, conn *websocket.Conn) {
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	ticker := time.NewTicker(pingPeriod)
	for {
		typ, content, err := conn.ReadMessage()
		if err != nil {
			log.Println("coon.ReadMessage() err:", err)
			//如果是客户端发送，再close一遍
			conn.Close()
			delete(room, username)
			return
		} else {
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if typ == websocket.TextMessage {
				MsgCh <- Msg{
					Time:     time.Now(),
					Username: username,
					Content:  string(content),
				}
			}
		}
		select {
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		default:
		}
	}
}

func Broadcast() {
	for {
		select {
		case msg := <-MsgCh:
			mplock.Lock()
			for _, j := range room {
				j.WriteJSON(msg)
			}
			mplock.Unlock()
		}
	}
}

func main() {
	r := gin.Default()
	go Broadcast()
	r.GET("/test", func(c *gin.Context) {
		//升级协议
		conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"status": 10022,
				"info":   "failed",
			})
			return
		}
		mlock.Lock()
		count++
		username := strconv.Itoa(count)
		mlock.Unlock()

		fmt.Println(username)
		mplock.Lock()
		room[username] = conn
		mplock.Unlock()
		go Rec(username, conn)
	})
	r.Run()
}
