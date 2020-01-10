package datasource

import (
	"log"
	"vk-poster2tg/cores"
	"vk-poster2tg/data/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgDatasource interface {
	SendPost(*model.VkPostModel)
}

type TgDatasourceImpl struct {
	ChannelName string
	BotApi      *tgbotapi.BotAPI
}

func InitBot(appConfig *cores.AppSettings) *TgDatasourceImpl {
	bot, err := tgbotapi.NewBotAPI(appConfig.TgApi)
	if err != nil {
		log.Panic(err)
	}

	return &TgDatasourceImpl{
		BotApi:      bot,
		ChannelName: appConfig.TgChannel,
	}
}

func (tgDatasourceImpl *TgDatasourceImpl) SendPost(post *model.VkPostModel) {
	if len(post.Images) > 0 {
		inputMedia := []interface{}{}
		for i, photoURL := range post.Images {
			newPhoto := tgbotapi.InputMediaPhoto{
				Type:  "photo",
				Media: photoURL.String(),
			}

			if i+1 == len(post.Images) {
				newPhoto.Caption = post.Text
			}

			inputMedia = append(inputMedia, newPhoto)
		}

		tgDatasourceImpl.BotApi.Send(tgbotapi.MediaGroupConfig{
			BaseChat: tgbotapi.BaseChat{
				ChannelUsername: tgDatasourceImpl.ChannelName,
			},
			InputMedia: inputMedia,
		})
	} else {
		msg := tgbotapi.NewMessageToChannel(tgDatasourceImpl.ChannelName, post.Text)
		tgDatasourceImpl.BotApi.Send(msg)
	}
}
