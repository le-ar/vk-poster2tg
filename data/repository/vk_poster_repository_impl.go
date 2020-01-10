package repository

import (
	"vk-poster2tg/data/datasource"
	"vk-poster2tg/data/model"
)

type VkPosterRepositoryImpl struct {
	VkPosterDatasource datasource.VkPosterDatasource
}

func (vkPosterRepositoryImpl *VkPosterRepositoryImpl) GetPosts() []*model.VkPostModel {
	return vkPosterRepositoryImpl.VkPosterDatasource.GetPosts()
}
