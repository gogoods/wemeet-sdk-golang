package main

import (
	"fmt"

	"github.com/gogoods/wemeet-sdk-golang/wemeet"
)

func QueryByID() {
	response, err := meeting.Do(&wemeet.MeetingQueryByIDRequest{
		MeetingID:  "10628915715042419559",
		InstanceID: wemeet.InstancePC,
		UserID:     "d9cc233205294fd8ae0bde019da11567",
	})

	if err != nil {
		fmt.Println(err)
		if e, ok := err.(wemeet.MeetingError); ok {
			fmt.Println("CODE:", e.Code)
			fmt.Println("MSG:", e.Message)
		}
	} else {
		result := response.(wemeet.MeetingQueryResponse)
		wemeet.PrintJson(&result)
	}
}
