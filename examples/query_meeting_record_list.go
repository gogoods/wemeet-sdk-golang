package main

import (
	"fmt"

	"github.com/gogoods/wemeet-sdk-golang/wemeet"
)

func QueryMeetingRecordList() {
	response, err := meeting.Do(&wemeet.MeetingRecordListRequest{
		MeetingID: "14147592802167314686",
		//MeetingCode: "672395158",
		UserID:    "d9cc233205294fd8ae0bde019da11567",
		StartTime: 1666598030,
		EndTime:   1666770830,
		PageSize:  10,
		Page:      1,
	})

	if err != nil {
		fmt.Println("err:", err)
		if e, ok := err.(wemeet.MeetingError); ok {
			fmt.Println("CODE:", e.Code)
			fmt.Println("MSG:", e.Message)
		}
	} else {
		result := response.(wemeet.MeetingRecordListResponse)
		wemeet.PrintJson(&result)
	}
}
