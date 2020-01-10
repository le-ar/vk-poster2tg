package main

import (
	"fmt"
	"net/url"
	"vk-poster2tg/cores"
	"vk-poster2tg/data/datasource"
	"vk-poster2tg/data/model"
	"vk-poster2tg/data/repository"
	"vk-poster2tg/domain/usecase"

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
	container.Provide(func(tgRepositoryImpl *repository.TgRepositoryImpl) *usecase.SendPostToTgChannel {
		return &usecase.SendPostToTgChannel{
			Repository: tgRepositoryImpl,
		}
	})

	return container
}

func main() {
	container := BuildContainer()

	// fmt.Println(container.Invoke(func(getAllPostsFromVkPoster *usecase.GetAllPostsFromVkPoster) {
	// 	fmt.Println(3)
	// 	fmt.Println(getAllPostsFromVkPoster.Execute())
	// 	fmt.Println(2)
	// }))

	fmt.Println(container.Invoke(func(sendPostToTgChannel *usecase.SendPostToTgChannel) {
		fmt.Println(3)

		urlPhoto, _ := url.Parse("https://icatcare.org/app/uploads/2018/07/Thinking-of-getting-a-cat.png")

		sendPostToTgChannel.Execute(&model.VkPostModel{
			Text: "Hello",
		})
		sendPostToTgChannel.Execute(&model.VkPostModel{
			Text: "Hello",
			Images: []*url.URL{
				urlPhoto,
			},
		})
		sendPostToTgChannel.Execute(&model.VkPostModel{
			Text: "Hello",
			Images: []*url.URL{
				urlPhoto,
				urlPhoto,
			},
		})
		sendPostToTgChannel.Execute(&model.VkPostModel{
			Images: []*url.URL{
				urlPhoto,
			},
		})
		sendPostToTgChannel.Execute(&model.VkPostModel{
			Images: []*url.URL{
				urlPhoto,
				urlPhoto,
			},
		})
		fmt.Println(2)
	}))
}
