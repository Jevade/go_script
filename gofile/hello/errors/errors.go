package errors

type Error struct {
	Err  error
	Code int
	Msg  string
}

const (
// ErrorUserNotLogin = ERROR{Err: errors.New("not login"), Code: 10401, Msg: "未登录"}
)
