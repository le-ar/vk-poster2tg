package controller

import (
	"math/rand"
	"sort"
	"time"
	"vk-poster2tg/cores"
	"vk-poster2tg/domain/usecase"
)

type BotController struct {
	getAllPostsFromVkPoster usecase.GetAllPostsFromVkPoster
	sendPostToTgChannel     usecase.SendPostToTgChannel
}

func InitBot(getAllPostsFromVkPoster usecase.GetAllPostsFromVkPoster, sendPostToTgChannel usecase.SendPostToTgChannel) *BotController {
	return &BotController{
		getAllPostsFromVkPoster: getAllPostsFromVkPoster,
		sendPostToTgChannel:     sendPostToTgChannel,
	}
}

func (botController *BotController) StartBot(appConfig *cores.AppSettings) {
	rand.Seed(time.Now().UnixNano())
	for {
		allPosts := botController.getAllPostsFromVkPoster.Execute()

		sort.Slice(allPosts[:], func(i, j int) bool {
			return allPosts[i].IA > allPosts[j].IA
		})

		priorityPosts := allPosts[:5]

		for _, post := range priorityPosts {
			botController.sendPostToTgChannel.Execute(post)

			time.Sleep(time.Minute * (30 + time.Duration(rand.Int63n(30))))
		}
	}
}
