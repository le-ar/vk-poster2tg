package usecase

import (
	"vk-poster2tg/data/model"
	"vk-poster2tg/domain/repository"
)

type GetAllPostsFromVkPoster struct {
	Repository repository.VkPosterRepository
}

func (usecase *GetAllPostsFromVkPoster) Execute() []*model.VkPostModel {
	return usecase.Repository.GetPosts()
}
