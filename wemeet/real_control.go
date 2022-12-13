package wemeet

// 会中管理: https://cloud.tencent.com/document/product/1095/57946

// 设置联席主持人：https://cloud.tencent.com/document/product/1095/57946
type RealControlCohostsRequest struct {
	MeetingID      string          `json:"-" param:"meeting_id"` // 会议的唯一 ID。
	OperatorID     string          `json:"operator_id"`          // 调用方用于标示用户的唯一 ID
	OperatorIDType int             `json:"operator_id_type"`     // 操作者 ID 的类型： 2：openid 4：ms_open_id
	UUID           string          `json:"uuid"`
	InstanceID     int             `json:"instanceid"` // 用户的终端设备类型
	Action         string          `json:"action"`     // 具体设置动作：true：设置联席主持人 false：撤销联席主持人
	User           *OperateUserObj `json:"user"`       //被操作用户对象信息。
}

func (req *RealControlCohostsRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorRealControlCohosts
}

func (req *RealControlCohostsRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *RealControlCohostsRequest) fillDefaultValue() {}

// 关闭会议中用户屏幕共享权限:https://cloud.tencent.com/document/product/1095/57949
type RealControlScreenSharedRequest struct {
	MeetingID      string          `json:"-" param:"meeting_id"` // 会议的唯一 ID。
	OperatorID     string          `json:"operator_id"`          // 调用方用于标示用户的唯一 ID
	OperatorIDType int             `json:"operator_id_type"`     // 操作者 ID 的类型： 2：openid 4：ms_open_id
	UUID           string          `json:"uuid"`
	InstanceID     int             `json:"instanceid"` // 用户的终端设备类型
	User           *OperateUserObj `json:"user"`       //被操作用户对象信息。
}

func (req *RealControlScreenSharedRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorRealControlScreenShared
}

func (req *RealControlScreenSharedRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *RealControlScreenSharedRequest) fillDefaultValue() {}

// 被操作用户对象
type OperateUserObj struct {
	ToOperatorID     string `json:"to_operator_id"`
	ToOperatorIDType int    `json:"to_operator_id_type"`
	UUID             string `json:"uuid"`
	InstanceID       int    `json:"instanceid"` // 用户的终端设备类型
}
