package service

import (
	"IM/global"
	"IM/logic"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Uid  uint64
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}
		userInfo := global.GlobalUsers[c.Uid]
		var msg = logic.Message{
			From:    c.Uid,
			To:      userInfo.Follow,
			Content: string(message),
			MsgType: userInfo.Type,
		}
		global.MessageChan <- msg

		switch userInfo.Type {
		case global.ORIENT:
			for _, v := range userInfo.Follow {
				if _, ok := global.GlobalUsers[v]; ok {
					msgTmp := <-global.MessageChan
					data, _ := json.Marshal(msgTmp)
					toClient := GlobalClient[v]
					toClient.Send <- data
				}
			}

		case global.RADIO:
			for _, toClient := range GlobalClient {
				msgTmp := <-global.MessageChan
				msg, _ := json.Marshal(msgTmp)
				toClient.Send <- msg
			}

		case global.CHATROOM:
			for _, v := range global.GlobalUsers {
				if userInfo.CompanyID == v.CompanyID {
					toClient := GlobalClient[v.UID]
					msgTmp := <-global.MessageChan
					msg, _ := json.Marshal(msgTmp)
					toClient.Send <- msg
				}
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		fmt.Println("client.go---line 81", string(message))
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
