package common

type RequestForm struct {
	UserId    int    `json:"user_id"`
	RequestId string `json:"request_id"`
	GameId    int    `json:"game_id" binding:"required,gt=0"`
	CpCode    string `json:"cp_code"`
	PkgName   string `json:"pkg_name"`
	Version   string `json:"version" binding:"required"`
	Os        string `json:"os" binding:"oneof=android ios"`
	DeviceId  string `json:"device_id" binding:"required"`
	Ip        string `json:"ip"`
	Ts        int64  `json:"ts"`
	Lang      string `json:"lang"`
	Debug     string `json:"debug" `
}

type TopicExtra struct {
	Id    string `json:"id"`
	ReTry int    `json:"re_try"`
	Ts    int64  `json:"ts"`
}
