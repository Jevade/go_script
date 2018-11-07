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

//UserInfo represent a user info
type UserInfo struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayhello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

//UserList represent a list of userinfo
type UserList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*UserInfo
}

//Token represents a web json token
type Token struct {
	Token string `json:"token"`
}
