package edd_socket

import (
	"fmt"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"io"
	"crypto/rand"
	"encoding/base64"
	"log"
	"golang.org/x/net/websocket"
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

func Echo(ws *websocket.Conn) {
	wss := NewWS(ws)
	uid := getUid()

	defer func() {
		wss.CloseUid(uid)
	}()

	for {
		var mess Message
		if err := wss.GetMsg(&mess); err != nil {
			//若当前socket接收不到消息，则已掉线，主动退出for
			fmt.Println(err)
			break
		}

		switch mess.Type {
		case "connect":
			//连接准备,客户端标识符name和服务端标识符绑定
			wss.BindUid(uid)
			msg := Message{
				Data: mess.Data,
				Type: "join_room",
			}

			wss.SendToAll(msg)
		case "all":
			//群发
			msg := Message{
				Data: mess.Data,
				Type: "send_all",
			}
			wss.SendToAll(msg)
		case "join_group":
			//wss.JoinGroup("one", uid)
		case "who":
		}
	}
}

func main() {
	var port = "8022"
	fmt.Println("listening port:" + port)

	go func() {
		http.Handle("/chat", websocket.Handler(Echo))
	}()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
