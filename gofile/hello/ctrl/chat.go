package ctrl

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"gopkg.in/fatih/set.v0"

	"github.com/gorilla/websocket"
)

var rwlock sync.RWMutex

//Node to store websocket conn and data
type Node struct {
	Conn      *websocket.Conn
	DataQuery chan []byte //并行转串行
	GroupSets set.Interface
}

var clientMap = make(map[int64]*Node)

//Chat 处理用户好友列表
func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userID, _ := strconv.ParseInt(id, 10, 64)
	isValid := checkToken(userID, token)
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return isValid
	}}).Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	node := &Node{
		Conn:      conn,
		DataQuery: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	rwlock.Lock()
	clientMap[userID] = node
	rwlock.Unlock()

	go sendproc(node)
	go recvproc(node)
	sendMsg(userID, []byte("Hello,I am Go boy!"))
}

func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQuery:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Printf("recv<-%s", data)
	}

}

func sendMsg(userID int64, msg []byte) {
	fmt.Printf("send->%s", msg)
	fmt.Println("")
	rwlock.RLock()
	node, ok := clientMap[userID]
	rwlock.RUnlock()
	if ok {
		node.DataQuery <- msg
	}
}

func checkToken(userid int64, token string) bool {
	user := userService.Find(userid)
	return user.Token == token
}
