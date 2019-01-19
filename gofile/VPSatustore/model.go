package main

import "time"

//VPS define a virtual host
type VPS struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Name      string     `gorm:"column:name" json:"-"`
	IP        string     `gorm:"column:ip" json:"-"`
	Hostname  string     `gorm:"column:hostname" json:"-"`
	NetworkID uint64     `gorm:"column:networkID" json:"-"`
	ImageID   string     `gorm:"column:imageID" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
	CPU       float32    `gorm:"column:cpu" json:"-"`
	RAM       float32    `gorm:"column:ram" json:"-"`
	DISK      float32    `gorm:"column:disk" json:"-"`
	MEM       float32    `gorm:"column:memory" json:"-"`
	TotalCPU  float32    `gorm:"column:totalcpu" json:"-"`
	TotalRAM  float32    `gorm:"column:totalram" json:"-"`
	TotalDISK float32    `gorm:"column:totaldisk" json:"-"`
	TotalMEM  float32    `gorm:"column:totalmem" json:"-"`
	VPSType   string     `gorm:"column:vpstype" json:"-"`
}

//User define a vps user
type User struct {
	ID         uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	UserName   string     `gorm:"column:username" json:"-"`
	PassWord   string     `gorm:"column:password" json:"-"`
	VPSLimit   string     `gorm:"column:vpsLimit" json:"-"`
	Level      uint64     `gorm:"column:level" json:"-"`
	Department string     `gorm:"column:department" json:"-"`
	Office     string     `gorm:"column:office" json:"-"`
	Email      string     `gorm:"column:email" json:"-"`
	Cellphone  string     `gorm:"column:cellphone" json:"-"`
	IsOn       bool       `gorm:"column:isOn" json:"-"`
	IsAdmin    bool       `gorm:"column:isAdmin" json:"-"`
	CreatedAt  time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt  time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt  *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

//VPSOwn define user's ownship for a vps
type VPSOwn struct {
	ID          uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	VPSID       uint64     `gorm:"column:vpsID" json:"-"`
	UserID      uint64     `gorm:"column:userID" json:"-"`
	VPSUsername string     `gorm:"column:vpsUsername" json:"-"`
	VPSPassword string     `gorm:"column:vpsPassword" json:"-"`
	IsOn        bool       `gorm:"column:isOn" json:"-"`
	IsPrivate   bool       `gorm:"column:isPrivate" json:"-"`
	CreatedAt   time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt   time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt   *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

//Network define a subnetwork of VPSs
type Network struct {
	IP string `gorm:"column:ip" json:"-"`
}

//Image define the image of a container
type Image struct {
	ID          uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	BaseImageID uint64    `gorm:"column:baseImageID" json:"-"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"-"`
	Size        float32   `gorm:"column:size" json:"-"`
}

//UsageInfo define vps usage info in different time
type UsageInfo struct {
	VPSID     uint64    `gorm:"column:vpsID" json:"-"`
	L1        float64   `gorm:"column:l1" json:"-"`
	L5        float64   `gorm:"column:l5" json:"-"`
	L15       float64   `gorm:"column:l15" json:"-"`
	MEM       float64   `gorm:"column:mem" json:"-"`
	CPU       float64   `gorm:"column:cpu" json:"-"`
	Disk      float64   `gorm:"column:disk" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
}
