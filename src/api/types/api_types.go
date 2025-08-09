package api_types

// 直播和录播的的列表的接口类型

type LiveListUserInfo struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	TeamLogo string `json:"teamLogo"`
	UserId   string `json:"userId"`
}

type LiveListContentInfo struct {
	CoverPath              string           `json:"coverPath"`
	Ctime                  string           `json:"ctime"`
	LiveId                 string           `json:"liveId"`
	RoomId                 string           `json:"roomId"`
	LiveType               int              `json:"liveType"` // 1：直播，2：电台，5：游戏
	LiveMode               int              `json:"liveMode"` // 0：正常，1：录屏
	Title                  string           `json:"title"`
	InMicrophoneConnection bool             `json:"inMicrophoneConnection"`
	Status                 int              `json:"status"`
	UserInfo               LiveListUserInfo `json:"userInfo"`
}

type LiveListContent struct {
	Next           string                `json:"next"`
	SlideUpAndDown bool                  `json:"slideUpAndDown"`
	LiveList       []LiveListContentInfo `json:"liveList"`
}

type LiveListResponse struct {
	Message string          `json:"message"`
	Status  int             `json:"status"`
	Success bool            `json:"success"`
	Content LiveListContent `json:"content"`
}

// 直播和录播的单个信息的类型

type LiveOneUser struct {
	UserAvatar string `json:"userAvatar"`
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
}

type LiveOneContent struct {
	LiveId         string      `json:"liveId"`
	Title          string      `json:"title"`
	RoomId         string      `json:"roomId"`
	PlayStreamPath string      `json:"playStreamPath"`
	Ctime          string      `json:"ctime"`
	User           LiveOneUser `json:"user"`
}

type LiveOneResponse struct {
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Success bool           `json:"success"`
	Content LiveOneContent `json:"content"`
}
