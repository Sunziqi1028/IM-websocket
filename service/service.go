package service

import (
	"IM/global"
	"IM/logic"
	"IM/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var GlobalClient = make(map[uint64]*Client)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	var conn *websocket.Conn
	conn, err := websocketUpgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	//defer conn.Close()

	var uid uint64 = 0           // 用户ID
	var partner_id uint64 = 0    // 用户Partner_ID
	var company_id uint64 = 0    // 用户组织ID
	var partner_name string = "" // 用户名称
	var follow []uint64          // 关注者
	var messageType = ""         //  消息类型 room ：聊天室 | radio：广播  | orient：定向
	if r.URL.RawQuery != "" {
		values, _ := url.ParseQuery(r.URL.RawQuery)
		intUid, _ := strconv.Atoi(values["uid"][0])
		uid = uint64(intUid)
		ok := utils.CheckUidUnique(uid) // 校验UID 是否唯一
		if !ok {
			global.GlobalUsers[uid] = logic.User{}
		} else {
			log.Println("改用户UID已经存在:", r)
			conn.Close()
			delete(MyServer.Clients, uid)
		}
		intPartnerId, _ := strconv.Atoi(values["partner_id"][0])
		partner_id = uint64(intPartnerId)
		ok = utils.CheckPartnerIDUnique(partner_id) // 校验partner_id是否唯一
		if !ok {
			global.GlobalUsers[uid] = logic.User{}
		} else {
			log.Println("该业务ID已经存在:", r)
			conn.Close()
			delete(MyServer.Clients, uid)
		}

		intCompanyId, _ := strconv.Atoi(values["company_id"][0])
		company_id = uint64(intCompanyId)
		partner_name = values["name"][0]
		follow, _ = utils.ConvertString2IntSlice(values["follow"][0])
		messageType = values["type"][0]
		var UserInfo = logic.User{
			UID:       uid,
			PartnerID: partner_id,
			CompanyID: company_id,
			Name:      partner_name,
			Type:      messageType,
			Follow:    follow,
		}
		fmt.Println(UserInfo, "service.go--65line")
		global.GlobalUsers[uid] = UserInfo // 存储全部的用户信息

	}

	client := &Client{
		Uid:  uid,
		Conn: conn,
		Send: make(chan []byte, 1024),
	}
	GlobalClient[uid] = client
	MyServer.Register <- client

	go client.Read()
	go client.Write()
}

// 升级http为websocket协议
func websocketUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// 升级http为websocket
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}
