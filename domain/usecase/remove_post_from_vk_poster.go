package usecase

import (
	"vk-poster2tg/data/model"
	"vk-poster2tg/domain/repository"
)

type RemovePostFromVkPoster struct {
	Repository repository.VkPosterRepository
}

func (usecase *RemovePostFromVkPoster) Execute(post *model.VkPostModel) {
	usecase.Repository.RemovePost(post)
}
