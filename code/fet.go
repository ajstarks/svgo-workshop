package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

var (
	canvas = svg.New(os.Stdout)
)

func ims(fname string) (int, int) {
	f, ferr := os.Open(fname)
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", ferr)
		return -1, -1
	}
	
	defer f.Close()
	switch strings.ToLower(path.Ext(fname)) {
	case ".jpg", ".jpeg":
		j, jerr := jpeg.DecodeConfig(f)
		if jerr == nil {
			return j.Width, j.Height
		} else {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fname, jerr)
			return -1, -1
		}
	case ".png":
		p, perr := png.DecodeConfig(f)
		if perr == nil {
			return p.Width, p.Height
		} else {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fname, perr)
			return -1, -1
		}
	}
	return -1, -1
}

func main() {
	imagefile := "scene.jpg"
	iw, ih := ims(imagefile)
	if iw < 0 || ih < 0 {
		return
	}
	
	filters := []string{
		"reference",	"f1",	"f2",	"f3",	"f4",	"f5",
		"sat3",	"sat2",	"sat0",	"rot0",	"rot1",	"rot2",
		"blur1", "blur2", "blur3", "sepia", "sepia2", "bright2",
	}


	gutter := 20
	nc := 6
	pw := (iw*nc) + gutter*(nc+1)
	ph := (ih*3) + gutter*4
	canvas.Start(pw, ph)
	canvas.Rect(0,0,pw,ph,canvas.RGB(0,0,0))
	canvas.Def()

	canvas.Filter("reference")
	canvas.Saturate(1.0)
	canvas.Fend()

	canvas.Filter("f1")
	canvas.FeComponentTransfer()
	canvas.FeFuncLinear("blue", 3, 0.0)
	canvas.FeFuncLinear("green", 1.5, 0.2)
	canvas.FeFuncLinear("red", 1.5, 0.2)
	canvas.FeCompEnd()
	canvas.Fend()

	canvas.Filter("f2")
	canvas.FeComponentTransfer()
	canvas.FeFuncGamma("b", 1, 0.2, 0)
	canvas.FeFuncGamma("R", 1, 0.707, 0)
	canvas.FeFuncGamma("g", 1, 0.707, 0)
	canvas.FeCompEnd()
	canvas.Fend()

	canvas.Filter("f3")
	canvas.FeComponentTransfer()
	canvas.FeFuncTable("G", []float64{0, 0.5, 0.6, 0.85, 1.0})
	canvas.FeCompEnd()
	canvas.Fend()

	canvas.Filter("f4")
	canvas.FeComponentTransfer()
	canvas.FeFuncTable("Red", []float64{1, 0})
	canvas.FeFuncTable("green", []float64{1, 0})
	canvas.FeFuncTable("b", []float64{1, 0})
	canvas.FeCompEnd()
	canvas.Fend()

	canvas.Filter("f5")
	canvas.FeComponentTransfer()
	canvas.FeFuncDiscrete("G", []float64{0.125, 0.375, 0.625, 0.875})
	canvas.FeCompEnd()
	canvas.Fend()
	
	
	canvas.Filter("sepia")
	canvas.Sepia()
	canvas.Fend()
	
	canvas.Filter("sepia2")
	p := 1.0
	s2 := [20]float64{
		1, p, p, 0,0, 
		p, 1, p, 0,0,
		p, p, 1, 0,0,
		0, 0, 0, 1, 0}
	canvas.FeColorMatrix(svg.Filterspec{}, s2)
	canvas.Fend()

	for i, sat := 0, 0.0; sat <= 0.75; sat += .25 {
		canvas.Filter(fmt.Sprintf("sat%d", i))
		canvas.Saturate(sat)
		canvas.Fend()
		i++
	}

	for i, r := 0, 90.0; r <= 360; r += 45 {
		canvas.Filter(fmt.Sprintf("rot%d", i))
		canvas.HueRotate(r)
		canvas.Fend()
		i++
	}

	for i, b := 0, 0.0; b < 20.0; b += 2.0 {
		canvas.Filter(fmt.Sprintf("blur%d", i))
		canvas.Blur(b)
		canvas.Fend()
		i++
	}

	for i, b := 0, 10.0; b < 100; b += 10 {
		canvas.Filter(fmt.Sprintf("bright%d", i))
		canvas.Brightness(b)
		canvas.Fend()
		i++
	}

	canvas.DefEnd()

	x := gutter
	y := gutter
	
	canvas.Gstyle("font-family:sans-serif;text-anchor:middle;font-size:16pt")
	for i, f := range filters {
		if i != 0 && i%nc == 0 {
			x = gutter
			y += ih + gutter
		}
		canvas.Image(x, y, iw, ih, imagefile, "filter:url(#"+f+")")
		canvas.Text(x+iw/2, y+ih/2, f)
		x += iw + gutter
	}
	canvas.Gend()
	canvas.End()
}
