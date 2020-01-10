package model

import (
	_ "image/jpeg"
	"net/url"
	"strconv"
)

type VkPostModel struct {
	Text   string
	Images []*url.URL
	IA     float64
}

func VkPostModelFromInterface(parsed map[string]interface{}) (*VkPostModel, error) {
	var photos []*url.URL

	for i := 1; i < 11; i++ {
		if photoURL, ok := parsed["image"+strconv.Itoa(i)]; ok && len(photoURL.(string)) > 0 {
			if urlPhoto, err := url.Parse(photoURL.(string)); err == nil {
				photos = append(photos, urlPhoto)
			}
		}
	}

	ia, _ := strconv.ParseFloat(parsed["ia"].(string), 64)

	return &VkPostModel{
		Text:   parsed["message"].(string),
		Images: photos,
		IA:     ia,
	}, nil
}
