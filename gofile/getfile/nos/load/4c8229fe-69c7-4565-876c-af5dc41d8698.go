package handler

import (
	"fmt"

	"../itemspy"
	"../model"
	"github.com/lexkong/log"
)

func SearchType(task *Task) {
	typech := make(chan interface{}, 4)
	i := 0
	go itemspy.GetChannelUrls(task.Shortcut, typech)
	func() {
		i++
		for elem := range typech {
			if err := elem.(*model.TypeInfo).Create(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(i)
		}
	}()
}

func ProcessType() {
	cityInfos, err := model.GetAllCity()
	for _, cityInfo := range *cityInfos {
		if err != nil {
			return
		}
		log.Info("Begin search type")
		task := &Task{Shortcut: cityInfo.Shortcut}
		SearchType(task)
	}
}
