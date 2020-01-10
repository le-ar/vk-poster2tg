package repository

import "vk-poster2tg/data/model"

type TgRepository interface {
	SendPost(*model.VkPostModel)
}
