package entity

import "image"

type VkPost struct {
	text   string
	images []*image.Image
}
