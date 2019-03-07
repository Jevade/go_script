package handler

import (
	"fmt"
	"time"

	"../itemspy"
	"../model"
)

var done chan string

//SearchItem will search all item in a url
func SearchItem(task *Task) {

	itemch := make(chan interface{}, 4)

	for i := 0; i < 10; i++ {
		go SaveItem(itemch)
	}
	for i := task.Pose + 1; i < int(task.Limit); i += task.Step {
		time.Sleep(500 * time.Millisecond)
		url := fmt.Sprintf("%spn%02d/", task.URL, i)
		fmt.Println("The url is:", url)
		itemspy.GetItemInfo(url, itemch)
	}
}

//Task is to define search task
type Task struct {
	URL   string `json:"url"`
	Step  int    `json:"step"`
	Pose  int    `json:"pose"`
	Limit int    `json:"limit"`
}

func SaveItem(itemch chan interface{}) {
	for item := range itemch {
		if nil == item {
			continue
		}
		if _, have, _ := model.GetItem(item.(*model.ItemInfo).ItemID); !have {
			continue
		}
		item.(*model.ItemInfo).Create()
		fmt.Println(item.(*model.ItemInfo).Desc)
	}
}

//Process is to run search item tasks
func Process() {
	typeInfo, err := model.GetAllType()
	if err != nil {
		return
	}
	taskch := make(chan interface{}, 200)
	defer close(taskch)
	go ProcessTask(taskch)
	go ProcessTask(taskch)
	go ProcessTask(taskch)
	go ProcessTask(taskch)
	for _, elem := range *typeInfo {
		fmt.Println("Send task")
		task := &Task{}
		task.URL = elem.URL
		task.Step = 1
		task.Pose = 1
		task.Limit = 50
		taskch <- task
		time.Sleep(60 * time.Second)
	}
	<-done
}

//ProcessTask is to process task
func ProcessTask(taskch <-chan interface{}) {
	for task := range taskch {
		SearchItem(task.(*Task))
	}
}
