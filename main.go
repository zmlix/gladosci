package main

import (
	"flag"
	"fmt"
	"gladosci/admin"
	"gladosci/checkin"
	"gladosci/pusher"
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

func Checkin() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	ciLog, err := os.OpenFile("checkinLog.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer ciLog.Close()
	log.SetOutput(ciLog)

	conf := admin.GetConfig()
	appToken := conf.AppToken
	users := conf.Users
	for _, user := range users {
		log.Println("正在签到用户:", user.UId)
		msg := checkin.CheckinOneUser(user.Cookie)
		err := pusher.WxPusher(appToken, user.UId, msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func CheckinOneUser(uId string) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	ciLog, err := os.OpenFile("checkinLog.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer ciLog.Close()
	log.SetOutput(ciLog)

	conf := admin.GetConfig()
	appToken := conf.AppToken
	users := conf.Users
	for _, user := range users {
		if user.UId == uId {
			fmt.Println("正在签到用户:", user.UId)
			msg := checkin.CheckinOneUser(user.Cookie)
			err := pusher.WxPusher(appToken, user.UId, msg)
			if err != nil {
				log.Println(err)
			}
			return nil
		}
	}
	return fmt.Errorf("未找到用户")
}

func CronCheckin(cronStr string) {
	c := cron.New(cron.WithSeconds())
	c.AddFunc(cronStr, Checkin)
	c.Run()
}

func main() {
	var cronStr string
	var addUser bool
	var checkinStr string
	flag.StringVar(&cronStr, "cron", "0 0 8 * * *", "默认开启定时任务")
	flag.BoolVar(&addUser, "add", false, "添加签到用户")
	flag.StringVar(&checkinStr, "checkin", "", "指定用户进行签到\n'all' 全部签到\nuId 指定uId用户签到")
	flag.Parse()

	if addUser {
		admin.AddUserTUI()
	} else if len(checkinStr) != 0 {
		switch checkinStr {
		case "all":
			Checkin()
		default:
			err := CheckinOneUser(checkinStr)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("签到成功")
			}
		}
	} else {
		CronCheckin(cronStr)
	}

}
