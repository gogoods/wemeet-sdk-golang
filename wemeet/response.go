package wemeet

import (
	"encoding/json"
	"errors"
	"log"
)

func (meeting Meeting) handleResponse(respBody []byte, descriptor *MeetingRequestDescriptor) (MeetingResponse, error) {
	var response MeetingResponse
	var err error
	switch *descriptor {
	case RequestDescriptorMeetingCreate:
		var resp MeetingCreateResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorUserCreate,
		RequestDescriptorUserUpdate,
		RequestDescriptorUserDelete,
		RequestDescriptorMeetingCancelByID:
		// 返回的body为空
		return nil, nil
	case RequestDescriptorUserDetailQuery:
		var resp UserDetailQueryResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorUserList:
		var resp UserListResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorMeetingQueryByCode, RequestDescriptorMeetingQueryByID:
		var resp MeetingQueryResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorMeetingUpdateByID:
		var resp MeetingUpdateResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorMeetingQueryParticipantsByID:
		var resp MeetingQueryParticipantsResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorMeetingQueryUserMeetingList:
		var resp MeetingQueryUserMeetingListResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorRecordList:
		var resp MeetingRecordListResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorRecordAddress:
		var resp MeetingRecordAddressResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	case RequestDescriptorRecordDetail:
		var resp MeetingRecordDetailResponse
		err = json.Unmarshal(respBody, &resp)
		response = resp
	default:
		log.Println("undefined response")
		log.Println(string(respBody))
		return string(respBody), errors.New("undefined response")
	}

	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
