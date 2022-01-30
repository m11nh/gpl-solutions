// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var (
	minHeight float64
	maxHeight float64
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func init() {
	minV, maxV := math.Inf(1), math.Inf(-1) 
 	for i := 0; i < cells; i++ {
 		for j := 0; j < cells; j++ {
 			x := xyrange * (float64(i)/cells - 0.5)
 			y := xyrange * (float64(j)/cells - 0.5)
 			z := f(x, y)
			min := math.Min(minV, z)
			if !math.IsNaN(min) {
				minV = min 
			}
			max := math.Max(maxV, z)
			if !math.IsNaN(max) {
				maxV = max 
			}
 		}
 	}
 	minHeight = minV
 	maxHeight = maxV
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			c := pointToColor(i, j)
			fmt.Printf(`<polygon points='%g,%g %g,%g %g,%g %g,%g' style="fill:rgb(%d, %d, %d);stroke:grey;stroke-width:1"/>\n`,
				ax, ay, bx, by, cx, cy, dx, dy, c.R, c.G, c.B)
		}
	}
	fmt.Println("</svg>")
}

func pointToColor(i, j int) color.RGBA {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// get the height
	z := f(x, y)

	// convert height to color scale
	f := uint8(((z - minHeight) / (maxHeight - minHeight)) * 255)
	return color.RGBA { f, 0, 255 - f, 0, }
}

func getColor(i, j int) (color.RGBA) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	fmt.Printf("%f\n", z)
	return color.RGBA{uint8(z), 0, 255 - uint8(z), 0}
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
