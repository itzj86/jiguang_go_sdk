package main

import (
	"fmt"
	"jiguang_go_sdk/jmessage"
)

func main() {
	fmt.Println("hello go")

	// user := [1]jmessage.UserInfo{
	// 	{"jostin", "123456"},
	// }

	//  注册 explop
	// users := make([]jmessage.UserInfo,0,0)
	// userList := append(users,jmessage.UserInfo{Username:"jostin",Password:"123456",Nickname:"张三"})
	// jmessage.Register(userList)
	// fmt.Printf("%T",userList)
	// fmt.Println(userList)

	// 更新用户信息
	// user := jmessage.UserInfo{
	// 	Username:"jostin",
	// 	Nickname:"张三",
	// 	Gender:1,
	// }
	// jmessage.UpdateuserInfo(user)

	// 获取用户信息
	// userInfo := jmessage.GetUserInfo("jostin")
	// fmt.Println(userInfo.Username)
	// 获取在线状态
	// res := jmessage.GetUserstat("jostin")
	// fmt.Printf("%v",res.Login)
	jmessage.UpdateUserPassword("jostin","12345678")


}
