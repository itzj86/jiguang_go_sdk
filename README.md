# jiguang_go_sdk
用golang实现的极光im的sdk
+ 批量用户注册
```go
 users := make([]jmessage.UserInfo,0,0)
 userList := append(users,jmessage.UserInfo{Username:"jostin",Password:"123456",Nickname:"张三"})
 jmessage.Register(userList)
```
+ 更新用户信息
```go
// 更新用户信息
 user := jmessage.UserInfo{
    Username:"jostin",
	  Nickname:"张三",
	  Gender:1,
 }
 jmessage.UpdateuserInfo(user)
```
+ 获取用户信息
```go
 userInfo := jmessage.GetUserInfo("jostin")
```
+ 获取在线状态
```go
res := jmessage.GetUserstat("jostin")
fmt.Printf("%v",res.Login)
```
+ 更新用户密码
```go
  jmessage.UpdateUserPassword("jostin","12345678")
```
