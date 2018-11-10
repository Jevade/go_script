package service

import (
	"fmt"
	"sync"

	"../model"
	"../util"
)

//ListUser will complete userdata logical and return info
func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	users, count, err := model.ListUser(username, int(offset), int(limit))
	if err != nil {
		return nil, count, err
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*model.UserInfo, len(users)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	var ids = make([]uint64, len(users))
	for idx, user := range users {
		ids[idx] = user.ID
	}

	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()
			shortID, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IDMap[u.ID] = &model.UserInfo{
				ID:        u.ID,
				Username:  u.Username,
				Password:  u.Password,
				SayHello:  fmt.Sprintf("Hello %s", shortID),
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	var userInfos = make([]*model.UserInfo, len(users))

	for idx, v := range ids {

		userInfos[idx] = userList.IDMap[v]
	}
	return userInfos, count, err
}
