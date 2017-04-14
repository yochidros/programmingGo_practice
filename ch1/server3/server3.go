package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xf0, 0x00, 0x10},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xf0}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
	blueIndex  = 3
)

var cycles int = 0

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Form {
			if k == "cycles" {
				var s = v[0]
				n, err := strconv.Atoi(s)
				cycles = n
				if err != nil {
					return
				}
				// fmt.Fprinf(w, "Form[%q] = %q\n", k, v)
			}
		}
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddress = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {

			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 1000
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(4)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
