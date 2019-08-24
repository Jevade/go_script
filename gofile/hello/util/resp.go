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
type H struct {
	Code  int         `json:"code"` //定义json，序列化时候的格式
	Rows  interface{} `json:"rows,omitempty"`
	Total int         `json:"total"`
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

func RespFailList(w http.ResponseWriter, msg string) {
	RespList(w, -1, H{Code: -1}, 0)
}

func RespOKList(w http.ResponseWriter, list interface{}, total int) {
	RespList(w, 0, list, total)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}
	w.Write(ret)
}
