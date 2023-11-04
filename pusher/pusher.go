package pusher

import (
	"log"

	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
)

func WxPusher(appToken, uId string, message string) error {
	msg := model.NewMessage(appToken).SetContentType(2).SetContent(message).AddUId(uId)
	if err := msg.Check(); err != nil {
		log.Println(err)
	}
	msgArr, err := wxpusher.SendMessage(msg)
	log.Println(msgArr)
	return err
}
