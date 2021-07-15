package jmessage

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	_ "time"
)

var baseApi = "https://api.im.jpush.cn"
var basic string

func init() {
	basic = base64.StdEncoding.EncodeToString([]byte("6605b7ef502f4a832e0540e3:4545d7aa275d461fd7ce03e4"))
}

func (u UserInfo) Register() (rest string) { //(resp string,err error)
	users := make([]UserInfo, 0)
	user := append(users, u)
	data, _ := json.Marshal(user)
	res := post("/v1/users/", strings.NewReader(string(data)))
	str := make([]map[string]string, 0)
	json.Unmarshal([]byte(res.String()), &str)
	fmt.Println(res.String())
	return res.String()
}
