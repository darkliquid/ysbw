package main

import (
	"image"

	"github.com/elazarl/goproxy"
)

func flipImage() func(image.Image, *goproxy.ProxyCtx) image.Image {
	return func(img image.Image, ctx *goproxy.ProxyCtx) image.Image {
		dx, dy := img.Bounds().Dx(), img.Bounds().Dy()

		nimg := image.NewRGBA(img.Bounds())
		for i := 0; i < dx; i++ {
			for j := 0; j <= dy; j++ {
				nimg.Set(i, j, img.At(i, dy-j-1))
			}
		}
		return nimg
	}
}
