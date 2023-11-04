package checkin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	StatusUrl  = "https://glados.rocks/api/user/status"
	CheckinUrl = "https://glados.rocks/api/user/checkin"
)

type StatusJson struct {
	Code int `json:"code"`
	Data struct {
		Code          string      `json:"code"`
		Domain        string      `json:"domain"`
		Phone         string      `json:"phone"`
		Hashed        string      `json:"hashed"`
		Password      string      `json:"password"`
		Port          int         `json:"port"`
		Traffic       int64       `json:"traffic"`
		Site          string      `json:"site"`
		ConfigureID   int         `json:"configureId"`
		UserID        int         `json:"userId"`
		TelegramID    interface{} `json:"telegram_id"`
		Vip           int         `json:"vip"`
		Email         string      `json:"email"`
		Days          int         `json:"days"`
		UsdtAddress   string      `json:"usdt_address"`
		ClientVersion string      `json:"client_version"`
		SystemTime    int64       `json:"system_time"`
		SystemDate    time.Time   `json:"system_date"`
		CreatedAt     int64       `json:"created_at"`
		ReferID       int         `json:"refer_id"`
		LeftDays      string      `json:"leftDays"`
	} `json:"data"`
}

type CheckinJson struct {
	Code    int    `json:"code"`
	Points  int    `json:"points"`
	Message string `json:"message"`
	List    []struct {
		ID       int    `json:"id"`
		UserID   int    `json:"user_id"`
		Time     int64  `json:"time"`
		Asset    string `json:"asset"`
		Business string `json:"business"`
		Change   string `json:"change"`
		Balance  string `json:"balance"`
		Detail   string `json:"detail"`
	} `json:"list"`
}

func GetStatusResp(url, cookie string) (StatusJson, error) {
	client := &http.Client{}
	data := StatusJson{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("referer", "https://glados.rocks/console/checkin")
	req.Header.Set("origin", "https://glados.rocks")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		return data, fmt.Errorf("请求失败: %w", err)
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	if err != nil {
		return data, fmt.Errorf("读取请求结果失败: %w", err)
	}

	// fmt.Println(data)
	if data.Code != 0 {
		return data, errors.New("Cookie 已过期")
	}
	return data, nil
}

func GetCheckinResp(url, cookie string) (CheckinJson, error) {
	client := &http.Client{}
	data := CheckinJson{}
	var body struct {
		Token string `json:"token"`
	}
	body.Token = "glados.one"
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return data, fmt.Errorf("序列化错误: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return data, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("referer", "https://glados.rocks/console/checkin")
	req.Header.Set("origin", "https://glados.rocks")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		return data, fmt.Errorf("请求失败: %w", err)
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	if err != nil {
		return data, fmt.Errorf("读取请求结果失败: %w", err)
	}

	// fmt.Println(data)
	return data, nil
}
