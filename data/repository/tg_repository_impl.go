package repository

import (
	"vk-poster2tg/data/datasource"
	"vk-poster2tg/data/model"
)

type TgRepositoryImpl struct {
	TgDatasource datasource.TgDatasource
}

func (tgRepositoryImpl *TgRepositoryImpl) SendPost(post *model.VkPostModel) {
	tgRepositoryImpl.TgDatasource.SendPost(post)
}
