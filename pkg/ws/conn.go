package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
	"sync"
)

var Upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Conn struct {
	id string
	*websocket.Conn
	context     context.Context
	writeChan   chan []byte // 写chan
	readChan    chan []byte
	errHandle   func(msg any, err error, conn *websocket.Conn) // 写err处理
	closeHandle func()                                         // close 处理
	cancel      context.CancelFunc
	once        sync.Once
	resizeEvent chan remotecommand.TerminalSize
}

type Options func(conn *Conn)

func CloseHandle(f func()) Options {
	return func(conn *Conn) {
		conn.closeHandle = f
	}
}

func ErrHandle(f func(msg any, err error, conn *websocket.Conn)) Options {
	return func(conn *Conn) {
		conn.errHandle = f
	}
}

func NewConn(ctx context.Context, id string, echoC echo.Context, opts ...Options) (*Conn, error) {
	ctx, cancel := context.WithCancel(ctx)
	conn, err := Upgrade.Upgrade(echoC.Response(), echoC.Request(), nil)
	if err != nil {
		return nil, err
	}
	c := &Conn{id: id, Conn: conn, writeChan: make(chan []byte, 1), readChan: make(chan []byte, 1),
		context: ctx, cancel: cancel}
	for _, opt := range opts {
		opt(c)
	}
	go c.Writing()
	go c.Reading()
	return c, nil
}

func (item *Conn) Writing() {
	for {
		select {
		case msg := <-item.writeChan:
			err := item.WriteMessage(1, msg)
			if err != nil {
				if item.errHandle != nil {
					item.errHandle(msg, err, item.Conn)
				} else {
					log.Fatalln("write json err ", err)
				}
			}
		case <-item.context.Done():
			return
		}
	}
}

func (item *Conn) Reading() {
	for {
		select {
		case <-item.context.Done():
			return
		default:
			t, data, err := item.ReadMessage()
			if err != nil {
				log.Println("reading err:", err)
				item.Close()
			}
			log.Println("type = ", t)
			log.Println("data = ", string(data))
			item.readChan <- data
		}
	}
}

func (item *Conn) Write(p []byte) (n int, err error) {
	item.writeChan <- p
	return len(p), nil
}

type xtermMessage struct {
	MsgType string `json:"type"`  // 类型:resize客户端调整终端, input客户端输入
	Input   string `json:"input"` // msgtype=input情况下使用
	Rows    uint16 `json:"rows"`  // msgtype=resize情况下使用
	Cols    uint16 `json:"cols"`  // msgtype=resize情况下使用
}

func (item *Conn) Read(p []byte) (n int, err error) {
	var (
		msg []byte
		// xtermMsg xtermMessage
	)

	// 读web发来的输入
	msg = <-item.readChan
	copy(p, msg)
	// // 解析客户端请求
	// if err = json.Unmarshal(msg, &xtermMsg); err != nil {
	// 	return 0, nil
	// }
	//
	// // web ssh调整了终端大小
	// if xtermMsg.MsgType == "resize" {
	// 	// 放到channel里，等remotecommand executor调用我们的Next取走
	// 	item.resizeEvent <- remotecommand.TerminalSize{Width: xtermMsg.Cols, Height: xtermMsg.Rows}
	// } else if xtermMsg.MsgType == "input" { // web ssh终端输入了字符
	// 	// copy到p数组中
	// 	n = len(xtermMsg.Input)
	// 	copy(p, xtermMsg.Input)
	// }
	return len(p), nil
}

func (item *Conn) Close() {
	item.once.Do(func() {
		if item.closeHandle != nil {
			item.closeHandle()
		}
		item.cancel()
		item.Conn.Close()
	})
}

func (item *Conn) Next() (size *remotecommand.TerminalSize) {
	ret := <-item.resizeEvent
	size = &ret
	return
}
