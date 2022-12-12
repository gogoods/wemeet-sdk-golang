package main

import "github.com/gogoods/wemeet-sdk-golang/wemeet"

var meeting = wemeet.Meeting{
	Registered: wemeet.EnableRegistered,
	Version:    "1.0.0",
	AppID:      "219284983",
	SdkID:      "20980278515",
	SecretID:   "DPepEnCDqdXiUnAG9Gd5CEnrOGUe8EQR",
	SecretKey:  "qKJBsi5Nl0oUOy0vF4Gtye91XaezWfFCTM7e6g0qNI1f3Gut",
}

func main() {
	CreateMeeting()
	//QueryByID()
	//QueryMeetingRecordList()
	//QueryMeetingRecordAddress()
	//QueryMeetingRecordDetail()
}
