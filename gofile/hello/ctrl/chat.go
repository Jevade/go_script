package ctrl

import (
	"encoding/json"
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

type Message struct {
	ID      int64  `json:"id,omitempty" form:"id"`
	UserID  int64  `json:"userid,omitempty" form:"userid"`
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`
	Media   int    `json:"medid,omitempty" form:"medid"`
	Content string `json:"content,omitempty" form:"content"`
	Pic     string `json:"pic,omitempty" form:"pic"`
	URL     string `json:"url,omitempty" form:"url"`
	Memo    string `json:"memo,omitempty" form:"memo"`
	Amount  int    `json:"amount,omitempty" form:"amount"`
}

/**
消息发送结构体
1、MEDIA_TYPE_TEXT 1
{id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
2、MEDIA_TYPE_NEWS 2
{id:1,userid:2,dstid:3,cmd:10,media:2,content:"标题",pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/dsturl","memo":"这是描述"}
3、MEDIA_TYPE_VOICE 3，amount单位秒
{id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}
4、MEDIA_TYPE_IMG 4
{id:1,userid:2,dstid:3,cmd:10,media:4,url:"http://www.baidu.com/a/log,jpg"}
5、MEDIA_TYPE_REDPACKAGR 5//红包amount 单位分
{id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}
6、MEDIA_TYPE_EMOJ 6
{id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}
7、MEDIA_TYPE_LINK 7
{id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}

8、MEDIA_TYPE_VIDEO 8
{id:1,userid:2,dstid:3,cmd:10,media:8,pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/a.mp4"}

9、MEDIA_TYPE_CONTACT 9
{id:1,userid:2,dstid:3,cmd:10,media:9,"content":"10086","pic":"http://www.baidu.com/a/avatar,jpg","memo":"胡大力"}

*/
var (
	MEDIA_TYPE_TEXT       = 1
	MEDIA_TYPE_NEWS       = 2
	MEDIA_TYPE_VOICE      = 3
	MEDIA_TYPE_IMG        = 4
	MEDIA_TYPE_REDPACKAGR = 5
	MEDIA_TYPE_EMOJ       = 6
	MEDIA_TYPE_LINK       = 7
	MEDIA_TYPE_VIDEO      = 8
	MEDIA_TYPE_CONTACT    = 9
)
var (
	CMD_SINGLE_MSG = 10
	CMD_ROOM_MSG   = 11
	CMD_HEART      = 0
)

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

		dispatch(data)
		fmt.Printf("recv<-%s", data)
	}

}
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		sendMsg(msg.Dstid, data)
	case CMD_ROOM_MSG:
		//群发逻辑
	case CMD_HEART:
		//心跳信息接受

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
