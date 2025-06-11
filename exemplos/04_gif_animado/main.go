package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // primeira cor
	blackIndex = 1 // segunda cor
)

func main() {
	const (
		cycles  = 5     // número de voltas completas
		res     = 0.001 // resolução angular
		size    = 100   // canvas
		nframes = 64    // número de frames
		delay   = 8     // tempo entre os frames (10ms)
	)

	freq := rand.Float64() * 3.0 // frequência relativa
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(os.Stdout, &anim) // saída do GIF para stdout
}
