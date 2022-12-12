package wemeet

// 创建用户
type UserCreateRequest struct {
	UserInfo
}

func (req *UserCreateRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserCreate
}

func (req *UserCreateRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

func (req *UserCreateRequest) fillDefaultValue() {}

// 获取用户信息
type UserDetailQueryRequest struct {
	UserID string `json:"userid" param:"userid"`
}

func (req *UserDetailQueryRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserDetailQuery
}

func (req *UserDetailQueryRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *UserDetailQueryRequest) fillDefaultValue() {}

type UserDetailQueryResponse struct {
	UserDetail
}

// 更新用户信息
type UserDetailUpdateRequest struct {
	UserID   string `json:"-" param:"userid"` // 调用方用于标示用户的唯一 ID
	Email    string `json:"email"`            // 新的邮箱地址
	Username string `json:"username"`         // 新的用户昵称
}

func (req *UserDetailUpdateRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserUpdate
}

func (req *UserDetailUpdateRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *UserDetailUpdateRequest) fillDefaultValue() {}

// 删除用户
type UserDeleteRequest struct {
	UserID string `json:"-" param:"userid"` // 调用方用于标示用户的唯一 ID
}

func (req *UserDeleteRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserDelete
}

func (req *UserDeleteRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *UserDeleteRequest) fillDefaultValue() {}

// 获取用户列表
type UserListRequest struct {
	Page     int `json:"page" query:"page"`           // 当前页
	PageSize int `json:"page_size" query:"page_size"` // 分页大小
}

func (req *UserListRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorUserList
}

func (req *UserListRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

func (req *UserListRequest) fillDefaultValue() {}

type UserListResponse struct {
	Pager
	Users []*UserDetail `json:"users"`
}
