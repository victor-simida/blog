package main

import (
	"image/png"
	"os"
	"flag"
	"fmt"
	"github.com/nfnt/resize"
)

var fileName = flag.String("f", "test.png", "filename")
var height = flag.Uint("h", 80, "height")
var width = flag.Uint("w", 80, "width")

var charSet string = "$@BB8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

//将png图片转换为字符图片
func main() {
	flag.Parse()
	file, err := os.Open(*fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("File Open Error:%s", err.Error())
		return
	}
	img, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Decode Error:%s", err.Error())
		return
	}

	result := []byte{}
	img = resize.Resize(*height, *width, img, resize.NearestNeighbor)
	for i := 0; uint(i) < *height; i++ {
		for j := 0; uint(j) < *width; j++ {
			c := getChar(img.At(i, j).RGBA())
			result = append(result, c)
		}
		result = append(result, '\n')
	}

	fmt.Printf(string(result))
}


func getChar(r, g, b, a uint32) byte {
	if a == 0 {
		return ' '
	}
	length := len(charSet)
	index := 0.2126 * float64(r) + 0.7152 * float64(g) + 0.0722 * float64(b)
	num := int(index) % length
	return charSet[num]

}

