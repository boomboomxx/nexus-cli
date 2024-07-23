package registry

import (
	"log"
	"testing"
)

func Test_Image_List(t *testing.T) {
	r, err := NewRegistry()
	if err != nil {
		log.Print(err)
	}
	images, err := r.ListImages()
	if err != nil {
		log.Print(err)
	}
	log.Print(images)
}
