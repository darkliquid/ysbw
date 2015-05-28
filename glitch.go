package main

import (
	"image"

	glitchify "github.com/darkliquid/glitch"
	"github.com/elazarl/goproxy"
)

func glitchImage() func(image.Image, *goproxy.ProxyCtx) image.Image {
	if *glitchiness < 0 || *glitchiness > 100 {
		*glitchiness = 5.0
	}

	if *brightness < 0 || *brightness > 100 {
		*brightness = 5.0
	}

	return func(img image.Image, ctx *goproxy.ProxyCtx) image.Image {
		return glitchify.Glitchify(img, *glitchiness, *brightness, *scanlines)
	}
}
