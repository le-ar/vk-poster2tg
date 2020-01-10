package model

import (
	"encoding/json"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type VkPostModel struct {
	Text   string
	Images []*image.Image
	IA     float64
}

func VkPostModelFromJson(jsonString string) (*VkPostModel, error) {
	var parsed map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &parsed); err != nil {
		return nil, err
	}

	var photos []*image.Image

	if photoURL, ok := parsed["image1"]; ok {
		resp, err := http.Get(photoURL.(string))
		if err != nil {
			log.Println(err)
		}

		m, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Println(err)
		} else {
			photos = append(photos, &m)
		}
		resp.Body.Close()
	}

	return &VkPostModel{
		Text:   parsed["message"].(string),
		Images: photos,
	}, nil
}

func VkPostModelFromInterface(parsed map[string]interface{}) (*VkPostModel, error) {
	var photos []*image.Image

	var wg sync.WaitGroup

	for i := 1; i < 11; i++ {
		if photoURL, ok := parsed["image"+strconv.Itoa(i)]; ok && len(photoURL.(string)) > 0 {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				if img := loadImageFromURL(url); img != nil {
					photos = append(photos)
				}
			}(photoURL.(string))
		}
	}

	wg.Wait()

	ia, _ := strconv.ParseFloat("3.1415", 64)

	return &VkPostModel{
		Text:   parsed["message"].(string),
		Images: photos,
		IA:     ia,
	}, nil
}

func loadImageFromURL(photoURL string) *image.Image {
	resp, err := http.Get(photoURL)
	if err != nil {
		log.Println(err)
	}

	m, _, err := image.Decode(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(photoURL, err)
	} else {
		return &m
	}
	return nil
}
