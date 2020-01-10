package datasource

import (
	"net/http"
)

type VkPosterDatasource interface {
	authVkPoster(login, pasword string)
	getPosts(groupID int)
}

type VkPosterDatasourceImpl struct {
	Client *http.Client
}

func (vkPosterDatasource VkPosterDatasourceImpl) authVkPoster(login, pasword string) {

}

func (vkPosterDatasource VkPosterDatasourceImpl) getPosts(groupID int) {

}