package collage

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

// WriteCollagePNG load images into memory, build a collage, write to disk.
func WriteCollagePNG(filename string, groundSize int, images []Imgres) error {

	collage, err := buildCollage(groundSize, images)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	if err := png.Encode(file, collage); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

// Imgloc stores the location of an image in a collage.
type Imgloc struct {
	Imgres
	X, Y int
}

func buildCollage(groundSize int, images []Imgres) (image.Image, error) {

	var imgLocs []Imgloc
	var barRow []Bar

	// lay out ground row
	x := 0
	for i := 0; i < groundSize; i++ {
		loc := Imgloc{Imgres: images[i], X: x, Y: 0}
		imgLocs = append(imgLocs, loc)
		barRow = append(barRow, Bar{W: images[i].W, H: images[i].H})
		x += images[i].W
	}

	bars := NewBarGraph(1)
	bars.StackRow(0, barRow)

	// stack remaining images
	for i := groundSize; i < len(images); i++ {

		gap := bars.LowIndex()

		x := 0
		for j := 0; j < gap; j++ {
			x += bars[j].W
		}

		loc := Imgloc{Imgres: images[i], X: x, Y: bars[gap].H}
		imgLocs = append(imgLocs, loc)

		bars.Stack(gap, Bar{W: images[i].W, H: images[i].H})
	}

	result := image.NewRGBA(image.Rectangle{
		Min: image.Pt(0, 0),
		Max: image.Pt(bars[0].W, bars[0].H),
	})

	for _, loc := range imgLocs {

		img, err := readImageFile(loc.Filename)
		if err != nil {
			return nil, err
		}

		r := image.Rect(loc.X, loc.Y, loc.X+loc.W, loc.Y+loc.H)
		draw.Draw(result, r, img, image.Pt(0, 0), draw.Over)
	}

	return result, nil
}

func readImageFile(filename string) (image.Image, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
