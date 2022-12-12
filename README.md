# tencent-meeting-sdk-golang
**腾讯会议SDK golang版**

tencent-meeting-sdk-golang 是用golang实现的一套用于调用腾讯会议API的开发工具包。该SDK目前已实现腾讯会议API文档中涉及的所有功能，包括企业会议管理以及企业用户管理。PS:该开发工具包由个人开发，并非官方发布的工具包。

#### 腾讯会议 REST API 简介
腾讯会议（Tencent Meeting，TM）Rest API 是为参与腾讯会议生态系统建设的合作方开发者接入并访问腾讯会议资源提供的一组工具，是访问腾讯会议 SaaS 服务的入口。合作伙伴可以通过腾讯会议 API 进行二次开发，例如创建一个会议，修改会议，查询会议信息等。
来源：https://cloud.tencent.com/document/product/1095/42407
## The Latest Release
#### v2.4.2
该SDK已兼容腾讯会议API 2.4.2。
API文档地址：https://cloud.tencent.com/document/product/1095/42407
API文档PDF版：https://main.qcloudimg.com/raw/document/product/pdf/1095_42406_cn.pdf

## 目录
- Tencent Meeting SDK Golang
  - [快速开始](#快速开始)
  - [示例](#示例)
  - [API列表](#API列表)
    - [企业会议管理](#企业会议管理) 
    - [企业用户管理](#企业用户管理)
    - [录制管理](#录制管理)
    - [注意事项](#注意事项)
      
  
## 快速开始
```go
package main

import (
	"fmt"
	"github.com/gogoods/wemeet-sdk-golang/wemeet"
	"strconv"
	"time"
)

func main() {

	meeting := wemeet.Meeting{
		Registered: wemeet.EnableRegistered, // 开启企业用户管理，建议开启
		Version:    "1.0.0",                    // 自定义版本号
		SecretKey:  "XXXXXXXXXXXXXX", 
		AppID:      "XXXXXXXXXXXXXXXX",
		SdkID:      "XXXXXXXXXXXXXXXXXXXX",
		SecretID:   "XXXXXXXXXXXXXXXXXXXXXX",
	}
    
	response, err := meeting.Do(wemeet.MeetingCreateRequest{
		InstanceID: wemeet.InstancePC,
		UserID:     "13800138000",
		Hosts: []*wemeet.UserObj{
			{
				UserID: "13800138000",
			},
		},
		Subject:   "测试会议",
		StartTime: strconv.Itoa(int(time.Now().Unix() + 60)),
		EndTime:   strconv.Itoa(int(time.Now().Unix() + 360)),
		Settings: &wemeet.Settings{
			MuteEnableJoin:  true,
			AllowUnmuteSelf: true,
		},
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
		
		/*
		会议主题 测试会议
		会议ID 96431717********848
		会议号 725***060
		开始时间 1596809701
		结束时间 1596810001
		密码 <nil>
		入会连接 https://meeting.tencent.com/s/qNl8**1a89f1
		*/

	}

}
```
如果没有AppID、 SdkID、SecretID、SecretKey的话可以通过 https://meeting.tencent.com/open-api.html 申请。
**另外，如果不开启企业用户管理，使用时会受比较多的限制，建议开启**

## 示例
[examples](./gogoods/wemeet-sdk-golang/tree/master/examples)

## API列表
### 企业会议管理
| 序号 | 功能 | 请求（Request） | 响应（Response） |  Doc  | 
| :---: | :---------- | :------------- | :------------- |  :---  | 
| 1 | 创建会议 | MeetingCreateRequest | MeetingCreateResponse |  https://cloud.tencent.com/document/product/1095/42417  | 
| 2 | 通过会议ID查询会议 | MeetingQueryByIDRequest | MeetingQueryByIDResponse |  https://cloud.tencent.com/document/product/1095/42418  | 
| 3 | 通过会议Code查询会议 | MeetingQueryByCodeRequest | MeetingQueryByCodeResponse |  https://cloud.tencent.com/document/product/1095/42420  | 
| 4 | 取消会议 | MeetingCancelRequest | N/A |  https://cloud.tencent.com/document/product/1095/42422  | 
| 5 | 修改会议 | MeetingUpdateRequest | MeetingUpdateResponse |  https://cloud.tencent.com/document/product/1095/42424  | 
| 6 | 获取参会成员列表 | MeetingQueryParticipantsRequest | MeetingQueryParticipantsResponse |  https://cloud.tencent.com/document/product/1095/42701  | 
| 7 | 查询用户的会议列表 | MeetingQueryUserMeetingListRequest | MeetingQueryUserMeetingListResponse |  https://cloud.tencent.com/document/product/1095/42421  | 

### 企业用户管理
| 序号 | 功能 | 请求（Request） | 响应（Response） |  Doc  | 
| :---: | :------: | :----------------- | :-------------------- | :---  | 
| 1 | 创建用户 | UserCreateRequest | N/A |   | 
| 2 | 更新用户 | UserDetailUpdateRequest | N/A |   | 
| 3 | 获取用户详情 | UserDetailQueryRequest | UserDetailQueryResponse |   | 
| 4 | 获取用户列表 | UserListRequest | UserListResponse |   | 
| 5 | 删除用户 | UserDeleteRequest | N/A |   | 

### 录制管理
| 序号 | 功能 | 请求（Request） | 响应（Response） |  Doc  | 
| :---: | :------: | :----------------- | :-------------------- | :---  | 
| 1 | 查询会议录制列表 | MeetingRecordListRequest | MeetingRecordResponse |  https://cloud.tencent.com/document/product/1095/51189  | 
| 2 | 查询会议录制地址 | MeetingRecordAddressRequest | MeetingRecordAddressResponse |  https://cloud.tencent.com/document/product/1095/51174  | 
| 3 | 查询单个录制详情（文件、纪要） | MeetingRecordDetailRequest | MeetingRecordDetailResponse |  https://cloud.tencent.com/document/product/1095/51180  | 

### 注意事项
* 构造请求时注意某些项为可选项，可以参照API文档的指示填写需要的项。另外可以看源码中对应请求的struct，如果某一项的tag中的 json 里有”omitempty“字样说明该项非必填项。

