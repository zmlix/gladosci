package main

import (
	"flag"
	"gladosci/admin"
	"gladosci/checkin"
	"gladosci/pusher"
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

func Checkin() {
	conf := admin.GetConfig()
	appToken := conf.AppToken
	users := conf.Users
	for _, user := range users {
		msg := checkin.CheckinOneUser(user.Cookie)
		err := pusher.WxPusher(appToken, user.UId, msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func CronCheckin(cronStr string) {
	c := cron.New(cron.WithSeconds())
	c.AddFunc(cronStr, Checkin)
	c.Run()
}

func main() {
	ciLog, err := os.OpenFile("checkinLog.log", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer ciLog.Close()
	log.SetOutput(ciLog)

	var cronStr string
	var addUser bool
	flag.StringVar(&cronStr, "cron", "0 0 8 * * *", "默认开启定时任务")
	flag.BoolVar(&addUser, "add", false, "添加签到用户")
	flag.Parse()

	if addUser {
		admin.AddUserTUI()
	} else if len(cronStr) > 0 {
		CronCheckin(cronStr)
	}

}
