package repository

import "vk-poster2tg/data/model"

type VkPosterRepository interface {
	GetPosts() []*model.VkPostModel
	RemovePost(post *model.VkPostModel)
}
