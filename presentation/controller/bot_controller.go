package controller

import (
	"math/rand"
	"sort"
	"time"
	"vk-poster2tg/cores"
	"vk-poster2tg/domain/usecase"
)

type BotController struct {
	removePostFromVkPoster  *usecase.RemovePostFromVkPoster
	getAllPostsFromVkPoster *usecase.GetAllPostsFromVkPoster
	sendPostToTgChannel     *usecase.SendPostToTgChannel
	sortBy                  string // "IA" or "Percent" Default - Percent
}

func InitBot(getAllPostsFromVkPoster *usecase.GetAllPostsFromVkPoster, sendPostToTgChannel *usecase.SendPostToTgChannel, removePostFromVkPoster *usecase.RemovePostFromVkPoster, appConfig *cores.AppSettings) *BotController {
	return &BotController{
		removePostFromVkPoster:  removePostFromVkPoster,
		getAllPostsFromVkPoster: getAllPostsFromVkPoster,
		sendPostToTgChannel:     sendPostToTgChannel,
		sortBy:                  appConfig.SortPostsBy,
	}
}

func (botController *BotController) StartBot() {
	rand.Seed(time.Now().UnixNano())
	lastUpdate := int64(0)
	for {
		if time.Now().Unix()-lastUpdate < 2000 {
			time.Sleep(time.Second * time.Duration(time.Now().Unix()-lastUpdate))
		}

		allPosts := botController.getAllPostsFromVkPoster.Execute()

		lastUpdate = time.Now().Unix()

		if botController.sortBy == "IA" {
			sort.Slice(allPosts[:], func(i, j int) bool {
				return allPosts[i].IA > allPosts[j].IA
			})
		} else {
			sort.Slice(allPosts[:], func(i, j int) bool {
				return allPosts[i].Percents > allPosts[j].Percents
			})
		}

		priorityPosts := allPosts[:5]

		for _, post := range priorityPosts {
			botController.sendPostToTgChannel.Execute(post)

			botController.removePostFromVkPoster.Execute(post)

			time.Sleep(time.Minute * (30 + time.Duration(rand.Int63n(30))))
		}
	}
}
