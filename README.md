# You Should Be Writing (YSBW)

You Should Be Writing instead of browsing imgur, let me help you...

## What?

YSBW is a simple image-replacement proxy server that changes all image
requests it can into ones that serve one you specify (or, by default,
one that says "You should be writing").

## Why?

My wife was lamenting the fact she was browsing imgur instead of actually
writing, so I knocked this up in a few minutes.

## How?

    ./ysbw --help
    usage: ysbw [<flags>]

    You Should Be Writing - an image proxy

    Flags:
      --help           Show help (also see --help-long and --help-man).
      -v, --verbose    Enable verbose messages
      -l, --listen=":8080"  
                       IP/port to listen on
      -c, --cache      Passthrough normal cache headers
      -r, --replace    replace images with a different one
      -p, --replace-file=REPLACE-FILE  
                       path to local image to use for replacements
      -u, --replace-url=REPLACE-URL  
                       replace all images with the provided one
      -f, --flip       flip images upside down
      -g, --glitch     glitch/corrupt images
      -a, --glitchiness=5.0  
                       amount to glitch
      -b, --brightness=5.0  
                       glitch brightness adjustment
      -s, --scanlines  glitch scanline filter
      --version        Show application version.

Simply run it and set your browsers proxy settings to refer to right place.
Chrome under Linux doesn't have a UI for proxy settings, so I run chrome
via `google-chrome --proxy-server="127.0.0.1:8080"` for example.
