package gorcle

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestRadius(t *testing.T) {

	blue := color.RGBA{0x00, 0x00, 0xff, 0xff}
	circle := NewCircle(blue, 4)

	fmt.Println(circle.String())
}

func TestPNGEncode(t *testing.T) {
	green := color.RGBA{0x00, 0xff, 0x00, 0xff}
	cirle := NewCircle(green, 50)
	theImage := image.NewNRGBA(image.Rect(0, 0, 400, 400))
	cirle.Draw(theImage, 200, 200)
	err := cirle.SavePNG("foo.png", theImage)
	if err != nil {
		fmt.Printf("an error happended during writing of png file. Error is: %s", err)
	}

}
