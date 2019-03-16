package model

import (
	"sync"
	"time"
)

// BaseModel is base enrty
type BaseModel struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

//ItemInfo represent a user info
type ItemInfo struct {
	ItemID    string  `json:"itemid"`
	UserID    string  `json:"userid"`
	ItemTitle string  `json:"itemTitle"`
	URL       string  `json:"url"`
	Desc      string  `json:"desc"`
	Province  string  `json:"Province"`
	City      string  `json:"city"`
	District  string  `json:"district"`
	Area      string  `json:"Area"`
	Type      string  `json:"Type"`
	Price     float32 `json:"price"`
	Telphone  string  `json:"telphone"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

//ItemList represent a list of userinfo
type ItemList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*ItemInfo
}

//Token represents a web json token
type Token struct {
	Token string `json:"token"`
}

//TypeInfo present type info and task info
type TypeInfo struct {
	ID       uint64 `json:"id"`
	Typename string `json:"typename"`
	URL      string `json:"url"`
	Tasknum  uint64 `json:"tasknum"`
	Amount   uint64 `json:"amount"`
	IsCoped  bool   `json:"isCoped"`
}

//CityInfo is to return cityinfo
type CityInfo struct {
	Citykey  uint64 `json:"citykey"`
	BaseHost string `json:"basehost"`
	CityName string `json:"cityname"`
	Province string `json:"Province"`
}
