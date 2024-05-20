/*
The "pxl" package provides functionality for image manipulations, like

finding the average brightness of a set of pixels (converted from an image
file) and similar functionality.
*/
package pxl

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// A pixel defines a set of RBGA values in range [0; 255].
type Pixel struct {
	R, G, B, A uint8
}

// Defines a slice of [pxl.Pixel].
//
// An image can be converted to [pxl.Pixels] by the
// [FromReader()] function.
type Pixels [][]Pixel

// Gets the pixels of an input [io.Reader], and returns it as a [Pixels] struct.
//
// Typically this function would be used to convert an image [*os.File] to a
// set of pixels.
//
// For using the manipulations methods of the [Pixels] struct, resize the
// reader (image) to a lower resolution for faster run times.
func FromReader(r io.Reader) (Pixels, error) {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("could not decode input file: %v", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel
	for y := bounds.Min.Y; y < height; y++ {
		var row []Pixel
		for x := bounds.Min.X; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

// Converts the input RGBA fields to a single [pxl.Pixel].
func rgbaToPixel(r, g, b, a uint32) Pixel {
	return Pixel{uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)}
}
