package wemeet

// MeetingRecordListRequest 查询会议录制列表
type MeetingRecordListRequest struct {
	MeetingID   string `query:"meeting_id"`
	MeetingCode string `query:"meeting_code"`
	UserID      string `query:"userid"`
	StartTime   int    `query:"start_time"`
	EndTime     int    `query:"end_time"`
	Page        int    `json:"page" query:"page"`           // 当前页
	PageSize    int    `json:"page_size" query:"page_size"` // 分页大小
}

func (req *MeetingRecordListRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorRecordList
}

func (req *MeetingRecordListRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *MeetingRecordListRequest) fillDefaultValue() {}

type MeetingRecordListResponse struct {
	TotalCount     int              `json:"total_count"`
	CurrentSize    int              `json:"current_size"`
	CurrentPage    int              `json:"current_page"`
	TotalPage      int              `json:"total_page"`
	RecordMeetings []*MeetingRecord `json:"record_meetings"`
}

// MeetingRecordAddressRequest 查询会议录制地址
type MeetingRecordAddressRequest struct {
	MeetingRecordID string `query:"meeting_record_id"`
	UserID          string `query:"userid"`
}

func (req *MeetingRecordAddressRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorRecordAddress
}

func (req *MeetingRecordAddressRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *MeetingRecordAddressRequest) fillDefaultValue() {}

type MeetingRecordAddressResponse struct {
	MeetingRecordID string            `json:"meeting_record_id"`
	MeetingID       string            `json:"meeting_id"`
	MeetingCode     string            `json:"meeting_code"`
	Subject         string            `json:"subject"`
	RecordFiles     []*LiveRecordFile `json:"record_files"`
}

// MeetingRecordDetailRequest 查询单个录制详情
type MeetingRecordDetailRequest struct {
	RecordFileID string `param:"meeting_id"`
	UserID       string `query:"userid"`
}

func (req *MeetingRecordDetailRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorRecordDetail
}

func (req *MeetingRecordDetailRequest) fillPlaceholder(args ...interface{}) string {
	return defaultPlaceholderFiller(req, args...)
}

func (req *MeetingRecordDetailRequest) fillDefaultValue() {}

type MeetingRecordDetailResponse struct {
	MeetingID               string            `json:"meeting_id"`
	MeetingCode             string            `json:"meeting_code"`
	RecordFileID            string            `json:"record_file_id"`
	ViewAddress             string            `json:"view_address"`
	DownloadAddress         string            `json:"download_address"`
	DownloadAddressFileType string            `json:"download_address_file_type"`
	AudioAddress            string            `json:"audio_address"`
	AudioAddressFileType    string            `json:"audio_address_file_type"`
	MeetingSummary          []*MeetingSummary `json:"meeting_summary"`
}
