package handler

import (
	"fmt"

	"../itemspy"
	"../model"
)

//ProcessTask is to process task
func ProcessCity(taskch <-chan interface{}) {
	for task := range taskch {
		SearchCity(task.(*Task))
	}
}

func SearchCity(task *Task) {
	citych := make(chan interface{}, 4)
	i := 0
	go itemspy.GetCityUrls(task.URL, citych)
	func() {
		i++
		for elem := range citych {
			if err := elem.(*model.CityInfo).Create(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(i)
		}
	}()
}
