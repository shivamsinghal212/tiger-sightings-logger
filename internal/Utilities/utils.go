package Utilities

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"time"
)

func ResizeImageUtil(file_path string) string {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	m := resize.Resize(250, 200, img, resize.NearestNeighbor)
	fileName := fmt.Sprintf("%d_resized.jpg", time.Now().Unix())
	out, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	return fileName
}
