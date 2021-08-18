```go
package main

import (
	"fmt"
	"github.com/xooooooox/sms1688"
)

func main() {
	result, err := sms1688.MobileCode(&sms1688.MsgBag{
		Url:       "http://xxxxxxxx/xxx",
		Username:  "xxxxxxxx",
		Password:  "xxxxxxxx",
		Content:   fmt.Sprintf("【xxx】您的验证码是 %s, 有效时间5分钟, 请妥善保存!", "123123"),
		PhoneList: []string{"13012345678"},
		CallData:  "",
	})
	if err != nil {
		fmt.Printf("mobile code error: %s\n", err.Error())
		return
	}
	fmt.Printf("%#v\n", result)
}
```