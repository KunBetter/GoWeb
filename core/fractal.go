package GoWeb

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func repeat(i, j int) color.RGBA {
	x := float64(i - 350) / 200
	y := float64(j - 250) / 200
	a := 0.0
	b := 0.0
	for t := 0; t < 256; t++ {
		m := a * a
		n := b * b
		o := a * b
		a = m - n + x
		b = o + o + y
		if m + n > 4 {
			return color.RGBA{uint8(t), uint8(t), uint8(t), 255}
		}
	}
	return color.RGBA{255, 255, 255, 255}
}

func Fractal() {
	file, _ := os.Create("mdb.png")
	defer file.Close()
	img := image.NewRGBA(image.Rect(0, 0, 5000, 5000))
	defer png.Encode(file, img)

	for i := 0; i < 5000; i++ {
		for j := 0; j < 5000; j++ {
			c := repeat(i, j)
			img.Set(i, j, c)
		}
	}
}