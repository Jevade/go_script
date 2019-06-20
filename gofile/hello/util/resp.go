package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Code int         `json:"code"` //定义json，序列化时候的格式
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := Message{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}
	w.Write(ret)
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, msg, nil)
}

func RespOK(w http.ResponseWriter, data interface{}) {
	Resp(w, 0, "成功", data)
}
