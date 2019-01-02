package message

const (
	LoginMsgType = "LoginMsg"
	LoginRspType = "LoginRsp"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginRsp struct {
	Code int `json:"code"`
	Error string `json:"error"`
}