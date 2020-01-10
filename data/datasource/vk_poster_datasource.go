package datasource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"vk-poster2tg/cores"
	"vk-poster2tg/data/model"
)

type VkPosterDatasource interface {
	GetPosts() []*model.VkPostModel
	RemovePost(post *model.VkPostModel)
}

type VkPosterDatasourceImpl struct {
	Client   *http.Client
	email    string
	password string
}

func (vkPosterDataSource *VkPosterDatasourceImpl) auth() {
	cookieJar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: cookieJar,
	}
	resp, err := client.PostForm("http://vk-poster.ru/core/login.php", url.Values{"email": {vkPosterDataSource.email}, "password": {vkPosterDataSource.password}})
	fmt.Println(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes), err)
	resp.Body.Close()
	vkPosterDataSource.Client = &client
}

func AuthVkPoster(appConfig *cores.AppSettings) *VkPosterDatasourceImpl {
	vkPosterDataSource := &VkPosterDatasourceImpl{
		email:    appConfig.Email,
		password: appConfig.Password,
	}

	vkPosterDataSource.auth()

	return vkPosterDataSource
}

func (vkPosterDatasource *VkPosterDatasourceImpl) GetPosts() []*model.VkPostModel {
	postsCount := 0
	resultPosts := []*model.VkPostModel{}
	for {
		resp, err := vkPosterDatasource.Client.PostForm("http://vk-poster.ru/core/feed/posts_download.php", url.Values{
			"startFrom":     {strconv.Itoa(postsCount)},
			"order_by":      {"0"},
			"order_dir":     {"0"},
			"show_common":   {"1"},
			"show_postpone": {"0"},
			"tabWall":       {"grabberwall"},
			"tabAudio":      {"grabberwallaudio"},
			"tabVideo":      {"grabberwallvideo"},
			"tabDoc":        {"grabberwalldoc"},
			"set":           {"2"},
		})
		if err != nil {
			fmt.Println(err)
		}

		postsCount += 5

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		resp.Body.Close()

		if strings.Contains(string(bodyBytes), "\"error_code\":1") {
			vkPosterDatasource.auth()
			return vkPosterDatasource.GetPosts()
		}

		var posts []map[string]interface{}

		if err = json.Unmarshal(bodyBytes, &posts); err == nil {
			for _, post := range posts {
				if vkPostModel, err := model.VkPostModelFromInterface(post); err == nil {
					resultPosts = append(resultPosts, vkPostModel)
				}
			}
		}

		if len(posts) < 5 {
			return resultPosts
		}
	}
}

func (vkPosterDatasource *VkPosterDatasourceImpl) RemovePost(post *model.VkPostModel) {
	resp, err := vkPosterDatasource.Client.PostForm("http://vk-poster.ru/core/feed/post_delete.php", url.Values{
		"msg":     {post.ID},
		"tabWall": {"grabberwall"},
		"set":     {"2"},
	})

	if err != nil {
		fmt.Println(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()

	if strings.Contains(string(bodyBytes), "\"error_code\":1") {
		vkPosterDatasource.auth()
		vkPosterDatasource.RemovePost(post)
	}
}
