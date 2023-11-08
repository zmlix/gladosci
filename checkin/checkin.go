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
		return fmt.Sprintf("<span>签到结果: %v</span><br/>", msg)
	}

	date := strings.Split(s.List[0].Business, ":")[2]
	points := strings.Split(s.List[0].Balance, ".")[0]

	return fmt.Sprintf("<span>签到结果: %v</span><br/><span>签到时间: %v</span><br/><span>总点数: %v</span><br/>", msg, date, points)
}

func CheckinOneUser(cookie string) string {
	checkin, err := GetCheckinResp(CheckinUrl, cookie)
	if err != nil {
		log.Println("GetCheckinResp", err)
		return fmt.Sprintf("<span>GetCheckinResp %v</span>", err.Error())
	}
	checkinStr := GetCheckinInfo(checkin)

	status, err := GetStatusResp(StatusUrl, cookie)
	if err != nil {
		log.Println("GetStatusResp", err)
		return checkinStr + fmt.Sprintf("<span>GetStatusResp %v</span>", err.Error())
	}
	statusStr := GetStatusInfo(status)
	return checkinStr + statusStr
}
