package sms1688

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type MsgBag struct {
	Url       string   `json:"-"`
	Username  string   `json:"userName"`
	Password  string   `json:"password"`
	Content   string   `json:"content"`
	PhoneList []string `json:"phoneList"`
	CallData  string   `json:"callData"`
}

type MsgResult struct {
	Code  int    `json:"code"`            // 0: success
	Msg   string `json:"message"`         // tips message
	MsgId int64  `json:"msgId,omitempty"` // send msg id
}

// MobileCode send mobile auth code
func MobileCode(msg *MsgBag) (result *MsgResult, err error) {
	if msg == nil {
		err = errors.New("msg parameter is nil")
		return
	}
	if msg.Url == "" || msg.Username == "" || msg.Password == "" || msg.Content == "" || len(msg.PhoneList) == 0 {
		err = errors.New("invalid parameter attribute, empty value exists")
		return
	}
	var bts []byte
	bts, err = json.Marshal(msg)
	if err != nil {
		return
	}
	client := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest("POST", msg.Url, bytes.NewReader(bts))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var bodies []byte
	bodies, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result = &MsgResult{}
	err = json.Unmarshal(bodies, result)
	return
}
