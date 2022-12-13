package wemeet

// 扩展应用接口
// https://doc.weixin.qq.com/doc/w3_m_KLdLmlChxpvg?scode=AJEAIQdfAAokxWulT9ACsAhga1ADI&version=4.0.19.6020&platform=win

// 会议绑定扩展应用
type MeetingBindAppRequest struct {
	Userid        string        `json:"userid"`     // 	string          `json:"uuid"`
	InstanceID    int           `json:"instanceid"` // 用户的终端设备类型
	MeetingID     string        `json:"meeting_id"`
	AutoOpenSdkid string        `json:"auto_open_sdkid"`
	ToolList      []*MeetingApp `json:"tool_list"`
}

func (req *MeetingBindAppRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetingBindApp
}

func (req *MeetingBindAppRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *MeetingBindAppRequest) fillDefaultValue() {
	if req.InstanceID == 0 {
		req.InstanceID = 1
	}
}

type MeetingApp struct {
	ToolAppid          string         `json:"tool_appid"`
	ToolSdkid          string         `json:"tool_sdkid"`
	VisibleType        uint32         `json:"visible_type"`
	VisibleList        []*VisibleItem `json:"visible_list"`
	UniqueCode         string         `json:"unique_code"`
	EnableCustomerData uint32         `json:"enable_customer_data"`
	EnableAddRobot     uint32         `json:"enable_add_robot"`
}

type VisibleItem struct {
	VisibleAppid  string `json:"visible_appid"`
	VisibleUserid string `json:"visible_userid"`
	VisibleOpenid string `json:"visible_openid"`
}

// 查询会议绑定的应用
type QueryMeetingAppRequest struct {
	Userid     string `query:"userid"`     // 	string          `json:"uuid"`
	InstanceID int    `query:"instanceid"` // 用户的终端设备类型
	MeetingID  string `query:"meeting_id"`
}

func (req *QueryMeetingAppRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorQueryMeetingApp
}

func (req *QueryMeetingAppRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *QueryMeetingAppRequest) fillDefaultValue() {
	if req.InstanceID == 0 {
		req.InstanceID = 1
	}
}

type QueryMeetingAppResponse struct {
	ToolList      []*MeetingApp `json:"tool_list"`
	AutoOpenSdkid string        `json:"auto_open_sdkid"`
}
