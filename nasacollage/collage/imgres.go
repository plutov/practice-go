package collage

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// Imgres image resolution
type Imgres struct {
	Filename string
	W, H     int
}

// ListDir read resolutions of images in a directory.
func ListDir(path string) ([]Imgres, error) {

	var result []Imgres

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	infos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	for _, inf := range infos {
		if inf.IsDir() {
			continue
		}

		filename := filepath.Join(path, inf.Name())
		f, err := os.Open(filename)
		conf, _, err := image.DecodeConfig(f)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", filename, err)
		}

		if err := f.Close(); err != nil {
			return nil, err
		}

		result = append(result, Imgres{
			Filename: filename,
			W:        conf.Width,
			H:        conf.Height})
	}

	return result, nil
}
