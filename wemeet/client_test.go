package wemeet_test

import (
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/gogoods/wemeet-sdk-golang/wemeet"
)

var meeting = wemeet.Meeting{

	Registered: wemeet.EnableRegistered,
}

func TestNewRequest(t *testing.T) {

	req, err := wemeet.NewRequest("GET", "https://api.meeting.qq.com/v1/meetings", "", meeting)
	if err != nil {
		t.Error(err)
	} else {
		client := wemeet.GetHttpClient()
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		} else {
			content, _ := ioutil.ReadAll(resp.Body)
			t.Log(string(content))
		}
	}

}

func TestMeeting_Do_CreateMeeting(t *testing.T) {

	resp, err := meeting.Do(&wemeet.MeetingCreateRequest{
		Settings: &wemeet.Settings{
			MuteEnableJoin:  true,
			MuteAll:         true,
			AllowUnmuteSelf: true,
		},
		InstanceID: wemeet.InstancePC,
		Type:       wemeet.MeetingTypeBookingMeeting,
		UserID:     "13800138000",
		Password:   "123456",
		StartTime:  strconv.Itoa(int(time.Now().Unix() + 30)),
		EndTime:    strconv.Itoa(int(time.Now().Unix() + 360)),
		Subject:    "测试会议",
	})

	if err != nil {
		t.Error(err)
	} else {

		r := resp.(wemeet.MeetingCreateResponse)
		t.Log(r.MeetingCreationInfo[0].Subject)
		t.Log(r.MeetingCreationInfo[0].StartTime)
		t.Log(r.MeetingCreationInfo[0].EndTime)
		t.Log("meeting code:", r.MeetingCreationInfo[0].MeetingCode)
		t.Log("meeting id:", r.MeetingCreationInfo[0].MeetingID)
		t.Log("meeting join url:", r.MeetingCreationInfo[0].JoinUrl)
		wemeet.PrintResponseJsonString(resp)

	}
}

func TestMeeting_Do_CreateUser(t *testing.T) {

	_, err := meeting.Do(&wemeet.UserCreateRequest{
		UserInfo: wemeet.UserInfo{
			UserID:   "13800138000",
			Username: "刘大胆",
			Phone:    "13800138000",
			Email:    "admin@163.com",
		},
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户创建成功")
	}
}

func TestMeeting_Do_QueryUserDetail(t *testing.T) {

	detail, err := meeting.Do(&wemeet.UserDetailQueryRequest{
		UserID: "13800138000",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户信息获取成功")
		d, ok := detail.(wemeet.UserDetailQueryResponse)
		if ok {
			t.Log(d.UserID)
			t.Log(d.Email)
			t.Log(d.Status)
			t.Log(d.Username)
			t.Log(d.AreaCode)
			t.Log(d.UpdateTime)
		} else {
			t.Error("判断错误")
		}
	}
}

func TestMeeting_Do_UpdateUser(t *testing.T) {

	_, err := meeting.Do(&wemeet.UserDetailUpdateRequest{
		Username: "刘华强",
		Email:    "huaqiang@test.com",
		UserID:   "13800138000",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户修改成功")
	}
}

func TestMeeting_Do_DeleteUser(t *testing.T) {

	_, err := meeting.Do(&wemeet.UserDeleteRequest{
		UserID: "13800138000",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("用户修改成功")
	}
}

func TestMeeting_Do_ListUser(t *testing.T) {

	list, err := meeting.Do(&wemeet.UserListRequest{
		Page:     1,
		PageSize: 20,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("数据获取成功")

		for _, d := range list.(wemeet.UserListResponse).Users {
			t.Log("-------------------------")
			t.Log(d.UserID)
			t.Log(d.Email)
			t.Log(d.Status)
			t.Log(d.Username)
			t.Log(d.AreaCode)
			t.Log(d.UpdateTime)
			t.Log("-------------------------")
		}
	}
}

func TestMeeting_Do_QueryByID(t *testing.T) {

	resp, err := meeting.Do(&wemeet.MeetingQueryByIDRequest{
		UserID:     "13800138000",
		MeetingID:  "3618694739060173625",
		InstanceID: wemeet.InstancePC,
	})
	if err != nil {
		t.Error(err)
	} else {
		r := resp.(wemeet.MeetingQueryResponse)
		t.Log(r.MeetingInfoList[0].Subject)
		t.Log(r.MeetingInfoList[0].StartTime)
		t.Log(r.MeetingInfoList[0].EndTime)
		t.Log(r.MeetingInfoList[0].Hosts[0].UserID)
		t.Log("meeting code:", r.MeetingInfoList[0].MeetingCode)
		t.Log("meeting id:", r.MeetingInfoList[0].MeetingID)
		t.Log("meeting join url:", r.MeetingInfoList[0].JoinUrl)
		t.Log("status:", r.MeetingInfoList[0].Status)
	}
}

func TestMeeting_Do_QueryByCode(t *testing.T) {

	resp, err := meeting.Do(&wemeet.MeetingQueryByCodeRequest{
		UserID:      "13800138000",
		MeetingCode: "290113240",
		InstanceID:  wemeet.InstancePC,
	})
	if err != nil {
		t.Error(err)
	} else {
		r := resp.(wemeet.MeetingQueryResponse)
		t.Log(r.MeetingInfoList[0].Subject)
		t.Log(r.MeetingInfoList[0].StartTime)
		t.Log(r.MeetingInfoList[0].EndTime)
		t.Log("meeting code:", r.MeetingInfoList[0].MeetingCode)
		t.Log("meeting id:", r.MeetingInfoList[0].MeetingID)
		t.Log("meeting join url:", r.MeetingInfoList[0].JoinUrl)
		t.Log("status:", r.MeetingInfoList[0].Status)
	}
}

func TestMeeting_Do_CancelMeeting(t *testing.T) {

	_, err := meeting.Do(&wemeet.MeetingCancelRequest{
		UserID:     "13800138000",
		MeetingID:  "3618694739060173625",
		InstanceID: wemeet.InstancePC,
		ReasonCode: 1080,
	})
	if err != nil {
		t.Error(err)
	} else {

	}
}

func TestMeeting_Do_UpdateMeeting(t *testing.T) {

	result, err := meeting.Do(&wemeet.MeetingUpdateRequest{
		UserID:     "13800138000",
		MeetingID:  "3710955756145476773",
		InstanceID: wemeet.InstancePC,
		Subject:    "我要修改会议信息",
	})
	if err != nil {
		t.Error(err)
	} else {
		log.Println(result.(wemeet.MeetingUpdateResponse).MeetingInfoList[0].MeetingID)
		log.Println(result.(wemeet.MeetingUpdateResponse).MeetingInfoList[0].MeetingCode)
	}
}

func TestMeeting_Do_QueryParticipants(t *testing.T) {

	r, err := meeting.Do(&wemeet.MeetingQueryParticipantsRequest{
		UserID:    "13800138000",
		MeetingID: "7814232111865567045",
	})
	if err != nil {
		t.Error(err)
	} else {
		wemeet.PrintResponseJsonString(r)
	}
}

func TestMeeting_Do_QueryUserMeetingList(t *testing.T) {

	r, err := meeting.Do(&wemeet.MeetingQueryUserMeetingListRequest{
		UserID:     "13800138000",
		InstanceID: 1,
	})
	if err != nil {
		t.Error(err)
	} else {
		wemeet.PrintResponseJsonString(r)
	}
}
