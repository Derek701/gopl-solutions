// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		ys := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			xs := float64(px+1)/width*(xmax-xmin) + xmin
			z := [4]complex128{ // Supersample four points.
				complex(x, y), complex(x, ys), complex(xs, y), complex(xs, ys)}
			// Image point (px, py) represents complex value z.
			var ss int
			for i := range z {
				ss += int(mandelbrot(z[i]))
			}
			ss /= 4
			img.Set(px, py, color.Gray{uint8(ss)})
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) uint8 {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - contrast*n
		}
	}
	return 0 // Black
}
