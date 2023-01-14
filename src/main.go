package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func load(filePath string) *image.NRGBA {
	imgFile, err := os.Open(filePath)
	check(err)
	defer imgFile.Close()

	// detect file format
	buffer := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	_, err = imgFile.Read(buffer)
	check(err)
	format := http.DetectContentType(buffer)
	// rewind reader to start of file
	imgFile.Seek(0, io.SeekStart)
	// format := "image/png"

	fmt.Printf("format: %s\n", format)
	switch format {
	case "image/png":
		img, err := png.Decode(imgFile)
		check(err)
		return img.(*image.NRGBA)
	case "image/jpeg":
		img, err := jpeg.Decode(imgFile)
		check(err)
		bounds := img.Bounds()
		return_image := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(return_image, return_image.Bounds(), img, bounds.Min, draw.Src)
		return return_image
	}
	return nil
}

func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	check(err)
	defer imgFile.Close()
	png.Encode(imgFile, img.SubImage(img.Rect))
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No arguments found. Please pass the path of the input image and optionally the output path of the image.")
		return
	}
	log.Printf("loading %s", os.Args[1])
	image := load(os.Args[1])
	max := image.Bounds().Max
	width := max.X
	height := max.Y
	for i := 0; i < width; i++ {
		process_column(image, i, height)
	}
	save("out.png", image)
}
