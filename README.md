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

    ./ysbw -help
    Usage of ./ysbw:
      -image="": set to a local path or an url to use that image instead of the default one
      -listen=":8080": ip/port to listen on
      -verbose=false: enable verbose logging
      
Simply run it and set your browsers proxy settings to refer to right place. 
Chrome under Linux doesn't have a UI for proxy settings, so I run chrome
via `google-chrome --proxy-server="127.0.0.1:8080"` for example.
