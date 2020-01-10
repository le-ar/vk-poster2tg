package main

import (
	"fmt"
	"vk-poster2tg/cores"
	"vk-poster2tg/data/datasource"
	"vk-poster2tg/data/repository"
	"vk-poster2tg/domain/usecase"
	"vk-poster2tg/presentation/controller"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(cores.ReadOrInit)
	container.Provide(datasource.AuthVkPoster)
	container.Provide(datasource.InitBot)
	container.Provide(func(authVkPoster *datasource.VkPosterDatasourceImpl) *repository.VkPosterRepositoryImpl {
		return &repository.VkPosterRepositoryImpl{
			VkPosterDatasource: authVkPoster,
		}
	})
	container.Provide(func(tgDatasource *datasource.TgDatasourceImpl) *repository.TgRepositoryImpl {
		return &repository.TgRepositoryImpl{
			TgDatasource: tgDatasource,
		}
	})
	container.Provide(func(vkPosterRepositoryImpl *repository.VkPosterRepositoryImpl) *usecase.GetAllPostsFromVkPoster {
		return &usecase.GetAllPostsFromVkPoster{
			Repository: vkPosterRepositoryImpl,
		}
	})
	container.Provide(func(vkPosterRepositoryImpl *repository.VkPosterRepositoryImpl) *usecase.RemovePostFromVkPoster {
		return &usecase.RemovePostFromVkPoster{
			Repository: vkPosterRepositoryImpl,
		}
	})
	container.Provide(func(tgRepositoryImpl *repository.TgRepositoryImpl) *usecase.SendPostToTgChannel {
		return &usecase.SendPostToTgChannel{
			Repository: tgRepositoryImpl,
		}
	})
	container.Provide(controller.InitBot)

	return container
}

func main() {
	container := BuildContainer()

	// fmt.Println(container.Invoke(func(getAllPostsFromVkPoster *usecase.GetAllPostsFromVkPoster) {
	// 	fmt.Println(3)
	// 	fmt.Println(getAllPostsFromVkPoster.Execute())
	// 	fmt.Println(2)
	// }))

	fmt.Println(container.Invoke(func(botController *controller.BotController) {
		fmt.Println(3)

		botController.StartBot()

		fmt.Println(2)
	}))
}
