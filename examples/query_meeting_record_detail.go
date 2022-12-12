package main

import (
	"fmt"

	"github.com/gogoods/wemeet-sdk-golang/wemeet"
)

func QueryMeetingRecordDetail() {
	response, err := meeting.Do(&wemeet.MeetingRecordDetailRequest{
		RecordFileID: "1584710376757571585",
		UserID:       "d9cc233205294fd8ae0bde019da11567",
	})

	if err != nil {
		fmt.Println("err:", err)
		if e, ok := err.(wemeet.MeetingError); ok {
			fmt.Println("CODE:", e.Code)
			fmt.Println("MSG:", e.Message)
		}
	} else {
		result := response.(wemeet.MeetingRecordDetailResponse)
		wemeet.PrintJson(&result)
	}
}
