package checkin

import (
	"fmt"
	"log"
	"strings"
)

func GetStatusInfo(s StatusJson) string {
	leftDays := strings.Split(s.Data.LeftDays, ".")[0]
	date := s.Data.SystemDate
	return fmt.Sprintf("<span>剩余天数: %v</span><br/><span>查询时间: %v</span>", leftDays, date)
}

func GetCheckinInfo(s CheckinJson) string {

	msg := s.Message

	if len(s.List) == 0 {
		return fmt.Sprintf("签到结果: %v", msg)
	}

	date := strings.Split(s.List[0].Business, ":")[2]
	points := strings.Split(s.List[0].Balance, ".")[0]

	return fmt.Sprintf("<span>签到结果: %v</span><br/><span>签到时间: %v</span><br/><span>总点数: %v</span><br/>", msg, date, points)
}

func CheckinOneUser(cookie string) string {
	status, err := GetStatusResp(StatusUrl, cookie)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("%v", err.Error())
	}
	checkin, err := GetCheckinResp(CheckinUrl, cookie)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("%v", err.Error())
	}
	statusStr := GetStatusInfo(status)
	checkinStr := GetCheckinInfo(checkin)
	return checkinStr + statusStr
}
