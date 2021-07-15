package jmessage

import (
	"encoding/json"
	"fmt"
	"strings"
	_ "time"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Gender   int    `json:"gender"` //性别 0 - 未知， 1 - 男 ，2 - 女
	// Birthday 	string `json:"birthday"`
	// Appkey 	string `json:"app_key"`
	// Signature string `json:"signature"` //	用户签名 	Byte(0~250)
	// Region 	string `json:"region"` //用户所属地区 	Byte(0~250)
	// Address 	string `json:"address"` //用户详细地址 	Byte(0~250)
	// Ctime 	 time.Time `json:"ctime"` //用户创建时间
	// Mtime 	time.Time `json:"mtime"` //用户最后修改时间
	// Extras 	string `json:"extras"`//用户自定义json对象 	Byte(0~512)
}

// 批量用户注册
func Register(userInfo []UserInfo) (rest string) {
	data, _ := json.Marshal(userInfo)
	res := post("/v1/users/", strings.NewReader(string(data)))
	str := make([]map[string]string, 0)
	json.Unmarshal([]byte(res.String()), &str)
	fmt.Println(res.String())
	return res.String()
}

// 更新用户信息
func UpdateuserInfo(userInfo UserInfo) (rest string) {
	data, _ := json.Marshal(userInfo)
	res := put("/v1/users/"+userInfo.Username, strings.NewReader(string(data)))
	str := make([]map[string]string, 0)
	json.Unmarshal([]byte(res.String()), &str)
	fmt.Println(res.String())
	return res.String()
}

// 获取用户信息
func GetUserInfo(username string) UserInfo { //
	res := get("/v1/users/" + username)
	userInfo := UserInfo{}
	json.Unmarshal([]byte(res.String()), &userInfo)
	return userInfo
}

// 获取用户在线状态
func GetUserstat(username string) struct {
	Login  bool `json:"login"`
	Online bool `json:"online"`
} { // (bool,bool)
	res := get("/v1/users/" + username + "/userstat")
	rest := struct {
		Login  bool `json:"login"`
		Online bool `json:"online"`
	}{}
	json.Unmarshal([]byte(res.String()), &rest)

	return rest
}

// 更新用户密码
func UpdateUserPassword(username string, password string) bool {
	var data = struct {
		NewPassword string `json:"new_password"`
	}{NewPassword: password}
	sendData, _ := json.Marshal(data)
	res := put("/v1/users/"+username+"/password", strings.NewReader(string(sendData)))
	rest := RError{}
	json.Unmarshal([]byte(res.String()), &rest)
	if rest.Error.Code > 0 {
		return false
	}
	return true
}

// 删除用户
func DeleteUser(username string) bool {
	res := delete("/v1/users/"+username, nil)
	rest := RError{}
	json.Unmarshal([]byte(res.String()), &rest)
	if rest.Error.Code > 0 {
		return false
	}
	return true
}

// 批量删除用户
func DeleteUsers(usernames []string) bool {
	sendData, _ := json.Marshal(usernames)
	res := delete("/v1/users/", strings.NewReader(string(sendData)))
	rest := RError{}
	json.Unmarshal([]byte(res.String()), &rest)
	if rest.Error.Code > 0 {
		return false
	}
	return true
}
