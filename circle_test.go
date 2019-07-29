package gorcle

import (
	"fmt"
	"image/color"
	"testing"
)

func TestRadius(t *testing.T) {

	blue := color.RGBA{0x00, 0x00, 0xff, 0xff}
	circle := NewCircle(blue, 4)

	fmt.Println(circle.String())
}
