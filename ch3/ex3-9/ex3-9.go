// Fractal emits a PNG image of a specified fractal (default: Mandelbrot).
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	width, height = 1024, 1024
)

var xmin, ymin, xmax, ymax float64

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	xmax, ymax = +2, +2
	if r.FormValue("x") != "" {
		xmax, _ = strconv.ParseFloat(r.FormValue("x"), 64)
	}
	if r.FormValue("y") != "" {
		ymax, _ = strconv.ParseFloat(r.FormValue("y"), 64)
	}
	if r.FormValue("zoom") != "" {
		zoom, _ := strconv.ParseFloat(r.FormValue("zoom"), 64)
		xmax /= zoom
		ymax /= zoom
	}
	xmin, ymin = -xmax, -ymax
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, fractal(r.FormValue("f"), z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func fractal(f string, z complex128) color.Color {
	switch f {
	case "acos":
		return acos(z)
	case "sqrt":
		return sqrt(z)
	case "newton":
		return newton(z)
	default:
		return mandelbrot(z)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
