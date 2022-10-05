// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
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

type CornerType int

const (
    middle CornerType = 0 // not the peak or valley of surface
    peak   CornerType = 1 // the peak of surface corner
    valley CornerType = 2 // the valley of surface corner
)


var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay,cta := corner(i+1, j)
			bx, by,ctb := corner(i, j)
			cx, cy,ctc := corner(i, j+1)
			dx, dy,ctd := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}

			var color string
            if cta == peak || ctb == peak || ctc == peak || ctd == peak {
                color = "#f00"
            } else if cta == valley || ctb == valley || ctc == valley || ctd == valley {
                color = "#00f"
            } else {
                // same as default
                color = "grey"
            }

            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64,CornerType) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z,ct := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ct
}

// edited function f to the egg box equation.

func f(x, y float64) (float64,CornerType) {
    d := math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12 
    ct := middle

    if math.Abs(d-math.Tan(d)) < 3 {
        ct = peak
        if 2*(math.Sin(d)-d*math.Cos(d))-d*d*math.Sin(d) > 0 {
            ct = valley
        }
    }
    return d, ct
}