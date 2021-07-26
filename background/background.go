package background

import (
	"image"
	"net/http"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// contains task data such image to load and widget to show the image
type ImageLoad struct {
	card   *widget.Card
	imgUrl string
}

// create buffered channle of 20 as task queue
var loads = make(chan ImageLoad, 20)

// submit load image task
func LoadImage(url string, card *widget.Card) {
	loads <- ImageLoad{
		card:   card,
		imgUrl: url,
	}

}

// received finished task from channel and display on GUI
func RefreshImage() {
	for load := range loads {
		if loadedImage, success := getImage(load.imgUrl); success {
			canvasImage := canvas.NewImageFromImage(loadedImage)
			canvasImage.FillMode = canvas.ImageFillOriginal
			load.card.SetImage(canvasImage)
		}
	}
}

// get movie image/poster from themoviedb.org
func getImage(url string) (image.Image, bool) {
	res, err := http.Get(url)
	if err != nil {
		return nil, false
	}

	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, false
	}

	return img, true
}
