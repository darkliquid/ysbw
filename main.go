package main

import (
	"image"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	goproxy_image "github.com/elazarl/goproxy/ext/image"
)

func main() {
	app.Author("Andrew Montgomery-Hurrell")
	app.Version("0.0.1")
	app.Parse(os.Args[1:])

	proxy := goproxy.NewProxyHttpServer()
	if *verbose {
		proxy.Verbose = *verbose
	}

	var processors []func(image.Image, *goproxy.ProxyCtx) image.Image

	// add replacement processor
	if replace != nil && *replace {
		processors = append(processors, replaceImage())
	}

	// add flip processor
	if flip != nil && *flip {
		processors = append(processors, flipImage())
	}

	// add glitch processor
	if glitch != nil && *glitch {
		processors = append(processors, glitchImage())
	}

	proxy.OnResponse().Do(goproxy_image.HandleImage(func(img image.Image, ctx *goproxy.ProxyCtx) image.Image {
		if !*cache {
			// Disable caching for images
			ctx.Resp.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
			ctx.Resp.Header.Set("Pragma", "no-cache")
			ctx.Resp.Header.Set("Expires", "0")
		}

		for _, f := range processors {
			img = f(img, ctx)
		}

		return img
	}))

	log.Printf("Listening on %q\n", *listen)
	log.Fatal(http.ListenAndServe(*listen, proxy))
}
