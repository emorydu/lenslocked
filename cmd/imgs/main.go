package main

import (
	"fmt"

	"github.com/emorydu/lenslocked/models"
)

func main() {
	gs := models.GalleryService{}
	fmt.Println(gs.Images(2))
}
