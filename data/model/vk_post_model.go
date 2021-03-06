package model

import (
	_ "image/jpeg"
	"net/url"
	"strconv"
	"strings"
)

type VkPostModel struct {
	ID       string
	Text     string
	Images   []*url.URL
	IA       float64
	Percents int
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
	percents, _ := strconv.Atoi(strings.Split(parsed["ia_view"].(string)[:len(parsed["ia_view"].(string))-2], "(")[1])

	return &VkPostModel{
		ID:       parsed["id"].(string),
		Text:     parsed["message"].(string),
		Images:   photos,
		IA:       ia,
		Percents: percents,
	}, nil
}
