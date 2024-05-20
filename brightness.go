package pxl

import (
	"math"
)

// Returns the perceived brightness of a set of [img.Pixels].
//
// The returned value is a floating point number in the range [0.0, 255.0].
func (pxls Pixels) Brightness() (float64, error) {
	avgR, avgG, avgB, _ := pxls.Average()
	brightness := math.Sqrt(
		0.241*math.Pow(float64(avgR), 2) +
			0.691*math.Pow(float64(avgG), 2) +
			0.068*math.Pow(float64(avgB), 2),
	)

	return brightness, nil
}

// Finds the average color of each RGBA color component on a set of
// [img.Pixels].
func (pxls Pixels) Average() (uint8, uint8, uint8, uint8) {
	width := len(pxls)
	height := len(pxls[0])
	var pxCount = width * height

	var avgR, avgG, avgB, avgA int
	for y := range height - 1 {
		for x := range width - 1 {
			p := pxls[x][y]
			avgR += int(p.R)
			avgG += int(p.G)
			avgB += int(p.B)
			avgA += int(p.A)
		}
	}

	r := uint32(avgR) / uint32(pxCount)
	g := uint32(avgG) / uint32(pxCount)
	b := uint32(avgB) / uint32(pxCount)
	a := uint32(avgA) / uint32(pxCount)

	return uint8(r), uint8(g), uint8(b), uint8(a)

}
