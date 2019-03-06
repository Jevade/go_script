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
	go itemspy.GetChannelUrls(task.URL, typech)
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
	log.Info("Begin search type")
	task := &Task{URL: "https://bj.58.com/sale.shtml"}
	SearchType(task)
}