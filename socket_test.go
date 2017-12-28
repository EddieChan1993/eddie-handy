package main

import (
	"fmt"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"io"
	"crypto/rand"
	"encoding/base64"
	"log"
	"goweb/edd_socket"
	"golang.org/x/net/websocket"
	"eddie-handy/edd_log"
)

func init() {
	edd_log.Config("logs/one.log")
}

//获取MD5字符串
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//获取连接唯一标识uid
func getUid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

func Echo(ws *websocket.Conn) {
	wss := edd_socket.NewWS(ws)
	uid := getUid()

	defer func() {
		wss.CloseUid(uid)
	}()

	for {
		var mess edd_socket.Message
		if err := wss.GetMsg(&mess); err != nil {
			//若当前socket接收不到消息，则已掉线，主动退出for
			fmt.Println(err)
			break
		}

		switch mess.Type {
		case "connect":
			//连接准备
			username := mess.Name
			wss.BindUid(uid, username)
			msg := edd_socket.Message{
				Name:    username,
				Content: mess.Content,
				Type:    "join_room",
			}

			wss.SendToAll(msg)
		case "all":
			//群发
			msg := edd_socket.Message{
				Name:    mess.Name,
				Content: mess.Content,
				Type:    "send_all",
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
