package gorcle

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

//this code is shamelessly taken and modified
// from https://github.com/ae6rt/golang-examples/blob/master/goeg/src/shaper2/shapes/shapes.go

// A Circle has a radius and a color.
// The zero value is invalid! Use NewCircle() to create a valid Circle.
type Circle struct {
	color  color.Color
	radius int
}

// NewCircle create a new Circle
// By calling newShape() we pass on any checking to newShape() without
// having to know what if any is required. This satisfies the Filler,
// Radiuser, Drawer, and Stringer interfaces.
func NewCircle(fill color.Color, radius int) *Circle {
	return &Circle{color: fill, radius: radius}
}

// Draw draws the given circle onto the image with the params x and y center coordinates within the image
func (circle *Circle) Draw(img draw.Image, x, y int) error {
	// Algorithm taken from
	// http://en.wikipedia.org/wiki/Midpoint_circle_algorithm
	// No need to check the radius is in bounds because you can only
	// create circles using NewCircle() which guarantees it is within
	// bounds. But the x, y might be outside the image so we check.
	if err := checkBounds(img, x, y); err != nil {
		return err
	}
	fill, radius := circle.color, circle.radius
	x0, y0 := x, y
	f := 1 - radius
	ddFx, ddFy := 1, -2*radius
	x, y = 0, radius

	img.Set(x0, y0+radius, fill)
	img.Set(x0, y0-radius, fill)
	img.Set(x0+radius, y0, fill)
	img.Set(x0-radius, y0, fill)

	for x < y {
		if f >= 0 {
			y--
			ddFy += 2
			f += ddFy
		}
		x++
		ddFx += 2
		f += ddFx
		img.Set(x0+x, y0+y, fill)
		img.Set(x0-x, y0+y, fill)
		img.Set(x0+x, y0-y, fill)
		img.Set(x0-x, y0-y, fill)
		img.Set(x0+y, y0+x, fill)
		img.Set(x0-y, y0+x, fill)
		img.Set(x0+y, y0-x, fill)
		img.Set(x0-y, y0-x, fill)
	}
	return nil
}

func (circle *Circle) String() string {
	return fmt.Sprintf("circle(fill=%v, radius=%d)", circle.color,
		circle.radius)
}

func checkBounds(img image.Image, x, y int) error {
	if !image.Rect(x, y, x, y).In(img.Bounds()) {
		return fmt.Errorf("point (%d, %d) is outside the image", x, y)
	}
	return nil
}

//End of taken and modified code

// SavePNG saves the circle as a png to the file fileName or to image.png
func (circle *Circle) SavePNG(fileName string, img draw.Image) error {
	currentFileName := "image.png"
	if len(fileName) > 0 {
		currentFileName = fileName
	}
	f, _ := os.Create(currentFileName)
	err := png.Encode(f, img)
	return err
}
