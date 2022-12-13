package main

import (
	"fmt"

	"github.com/gogoods/wemeet-sdk-golang/wemeet"
)

func QueryMeetingApp() {

	response, err := meeting.Do(&wemeet.QueryMeetingAppRequest{
		InstanceID: wemeet.InstancePC,
		Userid:     "d9cc233205294fd8ae0bde019da11567",
		MeetingID:  "9759507870638593493",
	})

	if err != nil {
		fmt.Println(err)
		if e, ok := err.(wemeet.MeetingError); ok {
			fmt.Println("CODE:", e.Code)
			fmt.Println("MSG:", e.Message)
		}
	} else {
		result := response.(wemeet.QueryMeetingAppResponse)
		wemeet.PrintJson(&result)
	}

}
