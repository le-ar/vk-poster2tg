package main

import (
	"fmt"
	"vk-poster2tg/cores"
	"vk-poster2tg/data/datasource"
	"vk-poster2tg/data/repository"
	"vk-poster2tg/domain/usecase"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(cores.ReadOrInit)
	container.Provide(datasource.AuthVkPoster)
	container.Provide(func(authVkPoster *datasource.VkPosterDatasourceImpl) *repository.VkPosterRepositoryImpl {
		return &repository.VkPosterRepositoryImpl{
			VkPosterDatasource: authVkPoster,
		}
	})
	container.Provide(func(vkPosterRepositoryImpl *repository.VkPosterRepositoryImpl) *usecase.GetAllPostsFromVkPoster {
		return &usecase.GetAllPostsFromVkPoster{
			Repository: vkPosterRepositoryImpl,
		}
	})

	return container
}

func main() {
	container := BuildContainer()

	fmt.Println(container.Invoke(func(getAllPostsFromVkPoster *usecase.GetAllPostsFromVkPoster) {
		fmt.Println(3)
		fmt.Println(getAllPostsFromVkPoster.Execute())
		fmt.Println(2)
	}))
}
