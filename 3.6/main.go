// minimal changes to - https://github.com/torbiak/gopl/blob/master/ex3.6/main.go
// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		epsX                   = (xmax - xmin) / width
		epsY                   = (ymax - ymin) / height
	)
	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}

	file, err := os.Create("img2.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file\n")
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// supersample
			c := []color.Color{}
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j ++ {
					c = append(c, mandelbrot(complex(x + offX[i], y + offY[j])))
				}
			}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, avg(c))
		}
	}
	png.Encode(file, img) // NOTE: ignoring errors
}

func avg(colors []color.Color) color.Color {
	start := color.RGBA{0, 0, 0 ,0}
	n := uint8(len(colors))
	for _, c:= range colors {
		r, g, b, a := c.RGBA()
		start.R += uint8(r / uint32(n))
		start.G += uint8(g / uint32(n))
		start.B += uint8(b / uint32(n))
		start.A += uint8(a / uint32(n))
	}
	return start
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black}

//!-
