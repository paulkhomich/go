package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width   = 640
	height  = 320
	steps   = 30.0
	cells   = 100
	angle   = math.Pi / 6
	xyscale = width / 2 / steps
	zscale  = height * .1
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	file, err := os.Create("surface.svg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err)
	}

	var data string
	data = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	file.WriteString(data)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := convPoints(i+1, j)
			bx, by := convPoints(i, j)
			cx, cy := convPoints(i, j+1)
			dx, dy := convPoints(i+1, j+1)

			data = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			file.WriteString(data)
		}
	}

	data = fmt.Sprintf("</svg>")
	file.WriteString(data)
	file.Close()
}

func convPoints(i int, j int) (float64, float64) {
	x := steps * (float64(i)/cells - 0.5)
	y := steps * (float64(j)/cells - 0.5)
	z := getFuncValue(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func getFuncValue(x float64, y float64) float64 {
	return math.Sin(math.Sqrt(x*x + y*y))
}
