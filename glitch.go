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
)

func main() {
	var seed = flag.Int("seed", 42, "seed")
	var interval = flag.Int("interval", 10, "interval")
	flag.Parse()

	rand.Seed(int64(*seed))

	outFilename := "blank.png"
	inFilename := "test.jpg"
	//width, height := 500, 500

	log.Print("reading file: ", inFilename)
	file, err := os.Open(inFilename)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)

	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(500, 500, img, resize.Lanczos3)

	//canvas := Newcanvas(image.Rect(0,0, 500, 500))
        i := image.NewRGBA(image.Rect(0, 0, 500, 500))
        gc := draw2d.NewGraphicContext(i)
        gc.Save()

	float_interval := float64(*interval)

	// move to a partition
	for i := 0; i < 500; i += *interval {
		for j := 0; j < 500 ; j += *interval {

			x := i
			y := j

			c := m.At(x + rand.Intn(*interval), y + rand.Intn(*interval))
			gc.Save()


			gc.Translate(float64(x), float64(y))
			gc.BeginPath()

			gc.MoveTo(0,0)
			gc.RLineTo(float_interval, 0)
			gc.RLineTo(0, float_interval)
			gc.RLineTo(-1 * float_interval, 0)

			gc.SetLineWidth(float_interval)
		        gc.SetStrokeColor(c)
			//gc.SetFillColor(c)

			//gc.Fill()
			//gc.FillStroke()
			gc.Stroke()
			//gc.ClosePath()

			//r := image.Rect(x, y, x + interval, y + interval)
			gc.Restore()
		}
	}


	out, err := os.Create(outFilename)

	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	log.Print("Saving  image to ", outFilename)
	png.Encode(out, i)

}
// import
