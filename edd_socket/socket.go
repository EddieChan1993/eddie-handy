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
	Name      string `json:"name"`
	Content   interface{} `json:"content"`
	Type      string `json:"type"`
	TimeStamp int64  `json:"time_stamp"`
}

//用户体
type User struct {
	Uid  string
	conn *websocket.Conn
	Name string
}

var (
	member         = make(map[string]*User)
	uidMapWs       = make(map[string]*websocket.Conn)
	groupMapMember = make(map[string][]*User)
)

//实例化ws
func NewWS(wss *websocket.Conn) *ws {
	return &ws{coon: wss}
}

//绑定uid
func (this *ws) BindUid(uid, name string) {
	client := User{Uid: uid, Name: name, conn: this.coon}

	member[uid] = &client
	uidMapWs[uid] = this.coon
}

//是否在线
func (this *ws) IsOnline(uid string) bool {
	_, exits := member[uid]
	return exits
}

//断开连接
func (this *ws) CloseUid(uid string) {
	msg := Message{
		Name:member[uid].Name,
		Content: "离开房间",
		Type:    "del_user",
	}
	this.SendToAll(msg)

	delete(member, uid)
	delete(uidMapWs, uid)
	this.coon.Close()
}

//群发消息
func (this *ws) SendToAll(msg Message) {
	msg.TimeStamp = time.Now().Unix()
	sendMess, _ := json.Marshal(msg)

	for k, v := range member {
		if v.conn != this.coon {
			if err := websocket.Message.Send(v.conn, string(sendMess)); err != nil {
				//如果发送断裂，则该socket掉线
				//删除相关map
				delete(member, k)
				delete(uidMapWs, k)
				continue
			}
		}
	}
}

//获取当前组人数
func (this *ws) GetClientCountByGroup(groupName string) int {
	return len(groupMapMember[groupName])
}

func (this *ws) GetClientByGroup(groupName string) []*User {
	return groupMapMember[groupName]
}

//加入某个群
func (this *ws) JoinGroup(groupName, uid string) {
	groupMapMember[groupName] = append(groupMapMember[groupName], member[uid])
}

//给指定群发消息
func (this *ws) SendToGroup(groupName string, msg Message) {
	msg.TimeStamp = time.Now().Unix()
	sendMess, _ := json.Marshal(msg)

	for k, v := range groupMapMember[groupName] {
		if v.conn != this.coon {
			if err := websocket.Message.Send(v.conn, string(sendMess)); err != nil {
				//如果发送断裂，则该socket掉线
				//删除当前组下面的切面中的元素即成员
				kk := k + 1
				groupMapMember[groupName] = append(groupMapMember[groupName][:k], groupMapMember[groupName][kk:]...)
				continue
			}
		}
	}
}

//离开某个群
func (this *ws) LeaveGroup(groupName, uid string) {
	for k, v := range groupMapMember[groupName] {
		if v.Uid == uid {
			kk := k + 1
			groupMapMember[groupName] = append(groupMapMember[groupName][:k], groupMapMember[groupName][kk:] ...)
			break
		}
	}
}

//发送给指定uid
func (this *ws) SendToUid(uid string, msg Message) {
	toWsCoon := uidMapWs[uid]
	msg.TimeStamp = time.Now().Unix()
	sendMess, _ := json.Marshal(msg)

	if err := websocket.Message.Send(toWsCoon, string(sendMess)); err != nil {
		delete(member, uid)
		log.Println(err)
	}
}

//解析客户端消息
func (this *ws) GetMsg(msg *Message) error {
	var reply string
	var err error
	if err = websocket.Message.Receive(this.coon, &reply); err == nil {
		json.Unmarshal([]byte(reply), msg)
	}
	return err
}
