package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	app     = kingpin.New("ysbw", "You Should Be Writing - an image proxy")
	verbose = app.Flag("verbose", "Enable verbose messages").Short('v').Bool()
	listen  = app.Flag("listen", "IP/port to listen on").Default(":8080").Short('l').String()
	cache   = app.Flag("cache", "Passthrough normal cache headers").Default("true").Short('c').Bool()

	replace     = app.Flag("replace", "replace images with a different one").Short('r').Bool()
	replaceFile = app.Flag("replace-file", "path to local image to use for replacements").Short('p').File()
	replaceURL  = app.Flag("replace-url", "replace all images with the provided one").Short('u').URL()

	flip = app.Flag("flip", "flip images upside down").Short('f').Bool()

	glitch      = app.Flag("glitch", "glitch/corrupt images").Short('g').Bool()
	glitchiness = app.Flag("glitchiness", "amount to glitch").Short('a').Default("5.0").Float()
	brightness  = app.Flag("brightness", "glitch brightness adjustment").Short('b').Default("5.0").Float()
	scanlines   = app.Flag("scanlines", "glitch scanline filter").Short('s').Bool()
)
