package main

import (
	"bytes"
	"flag"
	"image"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/elazarl/goproxy"
	goproxy_image "github.com/elazarl/goproxy/ext/image"
	"github.com/nfnt/resize"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	flag.BoolVar(&proxy.Verbose, "verbose", false, "enable verbose logging")

	var imagePath, listen string
	flag.StringVar(&imagePath, "image", "", "set to a local path or an url to use that image instead of the default one")
	flag.StringVar(&listen, "listen", ":8080", "ip/port to listen on")
	flag.Parse()

	var replaceImage image.Image
	var err error

	imagePathLower := strings.ToLower(imagePath)
	switch {
	case imagePath == "":
		replaceImage, _, err = image.Decode(bytes.NewReader(YouShouldBeWritingPNG))
	case strings.HasPrefix(imagePathLower, "http://") || strings.HasPrefix(imagePathLower, "https://"):
		res, err := http.Get(imagePath)
		if err != nil {
			log.Fatalf("Could not load image from %q\n", imagePath)
		}
		replaceImage, _, err = image.Decode(res.Body)
		res.Body.Close()
	default:
		_, err = os.Stat(imagePath)
		if err == nil {
			f, err := os.Open(imagePath)
			if err != nil {
				log.Fatalf("Could not load image from %q\n", imagePath)
			}
			replaceImage, _, err = image.Decode(f)
			f.Close()
		}
	}

	if err != nil {
		log.Fatalf("Error using image %q: %v\n", imagePath, err)
	}

	proxy.OnResponse().Do(goproxy_image.HandleImage(func(img image.Image, ctx *goproxy.ProxyCtx) image.Image {
		dx, dy := img.Bounds().Dx(), img.Bounds().Dy()
		return resize.Resize(uint(dx), uint(dy), replaceImage, resize.NearestNeighbor)
	}))

	log.Printf("Listening on %q\n", listen)
	log.Fatal(http.ListenAndServe(listen, proxy))
}
