package usecase

import (
	"vk-poster2tg/data/model"
	"vk-poster2tg/domain/repository"
)

type SendPostToTgChannel struct {
	Repository repository.TgRepository
}

func (usecase *SendPostToTgChannel) Execute(post *model.VkPostModel) {
	usecase.Repository.SendPost(post)
}
