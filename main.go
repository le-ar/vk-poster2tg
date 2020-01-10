package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"vk-poser2tg/data/datasource"
)

func main() {
	client := http.Client{}
	vkPosterDatasource := datasource.VkPosterDatasourceImpl{
		Client: &client,
	}
	fmt.Println(vkPosterDatasource)
	resp, err := client.PostForm("http://vk-poster.ru/core/login.php", url.Values{"email": {"malmuk2013@gmail.com"}, "password": {"Leoon2000"}})

	fmt.Println(err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes), err)
}
