//package main
//
//import (
//	"context"
//	"fmt"
//	"time"
//
//	"nhooyr.io/websocket"
//	"nhooyr.io/websocket/wsjson"
//)
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
//	defer cancel()
//
//	c, _, err := websocket.Dial(ctx, "ws://127.0.0.1:2022/chat?uid=1&partner_id=2&company_id=2&name=李四&follow=1,2&type=orient", nil)
//	if err != nil {
//		panic(err)
//	}
//	defer c.Close(websocket.StatusInternalError, "内部错误！")
//
//	err = wsjson.Write(ctx, c, "Hello WebSocket Server， 我是UID1")
//	if err != nil {
//		panic(err)
//	}
//
//	var v interface{}
//	err = wsjson.Read(ctx, c, &v)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("接收到服务端响应：%v\n", v)
//
//	c.Close(websocket.StatusNormalClosure, "")
//}

package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"

	//"golang.org/x/net/websocket"
	"log"
)

func main() {
	url := "ws://127.0.0.1:2022/chat?uid=1&partner_id=2&company_id=2&name=李四&follow=2&type=orient" //服务器地址
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {

		err := ws.WriteMessage(websocket.BinaryMessage, []byte("我是李四，你好，王二！"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 2)

	}()

	//reader := bufio.NewReader(os.Stdin)
	//go func() {
	//	lineBytes, _, _ := reader.ReadLine()
	//	//lineStr := string(lineBytes)
	//	ws.WriteMessage(websocket.BinaryMessage, lineBytes)
	//}()

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receive: ", string(data))
	}

}
