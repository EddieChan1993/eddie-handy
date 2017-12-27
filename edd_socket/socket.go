/**
	websocket封装
 */
package edd_socket

import (
	"time"
	"log"
	"encoding/json"
	"golang.org/x/net/websocket"
)

type ws struct {
	coon *websocket.Conn
}

//消息体
type Message struct {
	Data      string `json:"content"`
	Type      string `json:"type"`
	TimeStamp int64  `json:"time_stamp"`
}

//用户体
type user struct {
	uid string
	conn *websocket.Conn
	name string
}

var (
	member   =make(map[string]*user)
	uidMapWs =make(map[string]*websocket.Conn)
)

//实例化ws
func NewWS(wss *websocket.Conn) *ws{
	return &ws{coon:wss}
}

//绑定uid
func (this *ws) BindUid(uid,name string) {
	client:=user{uid:uid,name:name,conn:this.coon}

	member[uid]=&client
	uidMapWs[uid]=this.coon
}

//断开连接
func (this *ws) CloseUid(uid string) {
	delete(member,uid)
	delete(uidMapWs,uid)
	this.coon.Close()
}

//群发消息
func (this *ws) SendToAll(msg Message) {
	msg.TimeStamp=time.Now().Unix()
	sendMess,_:=json.Marshal(msg)

	for k,v:=range member {
		if v.conn!=this.coon{
			if err:=websocket.Message.Send(v.conn,string(sendMess));err!= nil {
				//如果发送断裂，则该socket掉线
				//删除相关map
				delete(member, k)
				delete(uidMapWs, k)
				continue
			}
		}
	}
}

//发送给指定uid
func (this *ws) SendToUid(uid string,msg Message) {
	toWsCoon:= uidMapWs[uid]
	msg.TimeStamp=time.Now().Unix()
	sendMess,_:=json.Marshal(msg)

	if err := websocket.Message.Send(toWsCoon, string(sendMess)); err != nil {
		delete(member, uid)
		log.Println(err)
	}
}

//解析客户端消息
func (this *ws) GetMsg(msg *Message) error{
	var reply string
	var err error
	if err=websocket.Message.Receive(this.coon,&reply);err==nil{
		json.Unmarshal([]byte(reply),msg)
	}
	return err
}