package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gogoods/wemeet-sdk-golang/wemeet"
)

func CreateMeeting() {

	response, err := meeting.Do(&wemeet.MeetingCreateRequest{
		InstanceID: wemeet.InstancePC,
		UserID:     "d9cc233205294fd8ae0bde019da11567",
		Hosts: []*wemeet.UserObj{
			{
				UserID: "d9cc233205294fd8ae0bde019da11567",
			},
		},
		Guests: []*wemeet.Guest{
			{
				Area:        "86",
				PhoneNumber: "18575515885",
				GuestName:   "",
			},
		},
		Subject:   "测试会议-20221121-3",
		StartTime: strconv.Itoa(int(time.Now().Unix() + 60)),
		EndTime:   strconv.Itoa(int(time.Now().Unix() + 3600)),
		Settings: &wemeet.Settings{
			MuteEnableJoin:  true,
			AllowUnmuteSelf: true,
		},
		EnableHostKey: true,
		HostKey:       "111111",
		Password:      "123456",
		//TimeZone:      "Asia/Shanghai",
	})

	if err != nil {
		fmt.Println(err)
		if e, ok := err.(wemeet.MeetingError); ok {
			fmt.Println("CODE:", e.Code)
			fmt.Println("MSG:", e.Message)
		}
	} else {
		result := response.(wemeet.MeetingCreateResponse)
		fmt.Println("会议主题", result.MeetingCreationInfo[0].Subject)
		fmt.Println("会议ID", result.MeetingCreationInfo[0].MeetingID)
		fmt.Println("会议号", result.MeetingCreationInfo[0].MeetingCode)
		fmt.Println("开始时间", result.MeetingCreationInfo[0].StartTime)
		fmt.Println("结束时间", result.MeetingCreationInfo[0].EndTime)
		fmt.Println("密码", result.MeetingCreationInfo[0].Password)
		fmt.Println("入会连接", result.MeetingCreationInfo[0].JoinUrl)
		fmt.Println("主持人秘钥", result.MeetingCreationInfo[0].HostKey)

		/*
			会议主题 测试会议
			会议ID 964XXXXXXXXX2387579848
			会议号 725XXX060
			开始时间 159XXX9701
			结束时间 1596XXX10001
			密码 <nil>
			入会连接 https://meeting.tencent.com/s/qNl8n01ad9f1
		*/

	}

}
