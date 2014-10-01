package main

import (
  "code.google.com/p/draw2d/draw2d"
  "github.com/nfnt/resize"
  "flag"
  "image"
  "image/png"
  "image/jpeg"
  "log"
  "math/rand"
  "os"
  "strconv"
)

func renderPng(gc *draw2d.ImageGraphicContext, m image.Image, interval int) {
	float_interval := float64(interval)
	gc.Save()

	// move to a partition
	for i := 0; i < 500; i += interval {
		for j := 0; j < 500 ; j += interval {
			x := i
			y := j

			c := m.At(x + rand.Intn(interval), y + rand.Intn(interval))
			gc.Save()


			gc.Translate(float64(x), float64(y))
			gc.BeginPath()

			gc.MoveTo(0,0)
			gc.RLineTo(float_interval, 0)
			gc.RLineTo(0, float_interval)
			gc.RLineTo(-1 * float_interval, 0)

			gc.SetLineWidth(float_interval)
		  gc.SetStrokeColor(c)

			gc.Stroke()

			gc.Restore()
		}
	}
}

func renderPngToFile(m image.Image, outFile string, seed int, interval int, ch chan<-bool) {
	// make a new image
  out_img := image.NewRGBA(image.Rect(0, 0, 500, 500))
  gc := draw2d.NewGraphicContext(out_img)

	renderPng(gc, m, interval)

	out, err := os.Create(outFile)

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	log.Print("Saving  image to ", outFile)
	png.Encode(out, out_img)

	ch <- true
}


func main() {

	outFilename := "out"
	inFilename := "test.jpg"

	var seed = flag.Int("seed", 42, "seed")
	var interval = flag.Int("interval", 10, "interval")
	flag.Parse()

	log.Print("reading file: ", inFilename)

	file, err := os.Open(inFilename)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("decoding jpg...")
	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.Print("ok")

	log.Print("resizing jpg")
	m := resize.Resize(500, 500, img, resize.Lanczos3)

	rand.Seed(int64(*seed))


	ch := make(chan bool)
	for iteration := 0 ; iteration < 10; iteration++ {
		go renderPngToFile(
			m,
		  outFilename + strconv.Itoa(iteration) + ".png",
		  *seed,
			*interval,
			ch)
	}

	for iteration := 0 ; iteration < 10; iteration++ {
		<-ch   // Wait for sort to finish; discard sent value.
	}
}
// import
