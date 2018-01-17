package edd_socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"io"
	"encoding/base64"
	"crypto/rand"
	"fmt"
	"log"
)

//获取MD5字符串
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//用于将当前用户信息和当前websocket.conn关联
//正式上线，可用数据库用户Id代替
func getUid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}
	wss := NewWs(conn)
	uid := getUid()

	defer func() {
		msg := Message{
			Content: "离开",
			Type:    "close",
		}
		wss.SendToAll(msg)
		wss.CloseUid(uid, msg)
	}()

	var mess Message
	for {
		if err:=wss.GetMsg(&mess);err!= nil {
			break
		}
		switch mess.Type {
		case "connect":
			fmt.Println(uid)
			wss.BindUid(uid)
			mess = Message{
				Content: "有人进入房间啦",
				Type:    "join_room",
			}
			wss.SendToAll(mess)
		case "all":
			wss.SendToAll(mess)
		case "who":
			msg := Message{
				Content: "你好",
				Type:    "xxxx",
			}
			if err:=wss.SendToUid(fmt.Sprintf("%s",mess.Content),msg);err!= nil {
				log.Println(err)
			}
		case "join_group":
			aa :=wss.JoinGroup(fmt.Sprintf("%s",mess.Content),uid)
			fmt.Println(aa)
		case "send_group":
			msg := Message{
				Content: "同志们好",
				Type:    "xxxx",
			}
			fmt.Println(mess)
			wss.SendToGroup(fmt.Sprintf("%s",mess.Content),msg)
		}

	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
