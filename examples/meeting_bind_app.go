package main

import (
	"fmt"

	"github.com/gogoods/wemeet-sdk-golang/wemeet"
)

func MeetingBindApp() {

	response, err := meeting.Do(&wemeet.MeetingBindAppRequest{
		InstanceID: wemeet.InstancePC,
		Userid:     "d9cc233205294fd8ae0bde019da11567",
		MeetingID:  "9759507870638593493",
		ToolList: []*wemeet.MeetingApp{
			&wemeet.MeetingApp{
				ToolAppid: "219284983",
				ToolSdkid: "21370167653",
			},
		},
	})

	if err != nil {
		fmt.Println(err)
		if e, ok := err.(wemeet.MeetingError); ok {
			fmt.Println("CODE:", e.Code)
			fmt.Println("MSG:", e.Message)
		}
	} else {
		wemeet.PrintJson(&response)
	}

}
