package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/image", imageHandler)
	http.ListenAndServe(":8081", nil)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	width, height := 256, 256
	rectImg := image.NewRGBA(image.Rect(0, 0, width, height))

	backgroundColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{backgroundColor}, image.Point{}, draw.Src)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	numRectangles := 4 + rng.Intn(5)
	for i := 0; i < numRectangles; i++ {
		randomColor := color.RGBA{
			R: uint8(rng.Intn(256)),
			G: uint8(rng.Intn(256)),
			B: uint8(rng.Intn(256)),
			A: 255,
		}

		minX, minY := rng.Intn(width-20), rng.Intn(height-20)
		maxX, maxY := minX+40+rng.Intn(40), minY+40+rng.Intn(40)
		randomRect := image.Rectangle{
			Min: image.Point{X: minX, Y: minY},
			Max: image.Point{X: maxX, Y: maxY},
		}

		draw.Draw(rectImg, randomRect, &image.Uniform{randomColor}, image.Point{}, draw.Src)
	}
	w.Header().Set("Content-Type", "image/png")

	png.Encode(w, rectImg)
}
