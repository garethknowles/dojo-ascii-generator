package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	// filename := "./small-image.png"
	// filename := "./bigger-image.png"
	filename := "./smiley-big.jpg"

	fmt.Println("D&G ASCII Image Machine")

	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file)
	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	ascii := pixelsToAscii(pixels)

	printAscii(ascii)
}

func printAscii(ascii [][]string) {

	for _, row := range ascii {
		var printRow string
		for _, value := range row {
			printRow = printRow + value
		}
		fmt.Println(printRow)
	}
}

func pixelsToAscii(pixels [][]Pixel) [][]string {

	var ascii [][]string
	for _, row := range pixels {
		var newRow []string
		for _, value := range row {
			asciiVal := pixelToAscii(value)
			newRow = append(newRow, asciiVal)
		}
		ascii = append(ascii, newRow)
	}
	return ascii
}

func pixelToAscii(pixel Pixel) string {
	if pixel.R < 120 {
		return "#"
	}
	return "."
}

func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

// func averagePixels(pixels [][]Pixel) [][]Pixel {

// 	var averagedPixels [][]Pixel

// 	for y := 0; y < len(pixels); y++ {
// 		row := pixels[y]
// 		var newRow []Pixel
// 		for x := 0; x < len(row); x++ {
// 			val := row[x]
// 			prevVal := row[x-1]
// 			nextVal := row[x+1]
// 			averagedVal := Pixel{
// 				int(val.R + prevVal.R + nextVal.R/3),
// 				int(val.G + prevVal.G + nextVal.G/3),
// 				int(val.B + prevVal.B + nextVal.B/3),
// 				int(val.A + prevVal.A + nextVal.A/3)}

// 			newRow = append(newRow, averagedVal)
// 		}
// 	}
// 	return averagedPixels
// }

// func averageRGB(pixels [][]Pixel) Pixel {
// 	totalR := 0
// 	totalG := 0
// 	totalB := 0

// 	for x := 0; x < len(pixels); x++ {
// 		pixel := pixels[x]
// 		totalR += pixel.R
// 		totalG += pixel.G
// 		totalB += pixel.B

// 	}

// 	return Pixel{
// 		int(val.R + prevVal.R + nextVal.R/pixels),
// 		int(val.G + prevVal.G + nextVal.G/3),
// 		int(val.B + prevVal.B + nextVal.B/3),
// 		int(val.A + prevVal.A + nextVal.A/3)}

// }

// func isDark(pixel Pixel) string {
// 	return ((pixel.R + pixel.G + pixel.B) < (120 * 3))
// }
