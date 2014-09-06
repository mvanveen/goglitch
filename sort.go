package main

import (
  "code.google.com/p/draw2d/draw2d"
  "github.com/nfnt/resize"
  "image"
  "image/png"
  "image/jpeg"
  "log"
  "os"
)

func brightness (color.RGBA color) int brightness{
	//brightness = ;
}
type ImageRow struct {
    Row     image.Rectangle
    Width   int
}

func (a ImageRow) Len() int      { return len(a.Width) }
func (a ImageRow) Less(i, j int) { return a.Row.At(0, i) < a.Row.At(0, j) }
func (a ImageRow) Swap(i, j int) {
	color := a.Row.At(0, j)

	a.Row.Set(0, j, a.Row.At(0, i))
	a.Row.Set(0, i, color)
}

func main() {
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

	// move to a partition
	for i := 0; i < 500; i++ {

		row_rect := image.Rect(i, 0, i, 500)
		// take whole row
		row = ImageRow({
		  Row: m.SubImage(row_rect),
		  Width: 500
		})

		// sort it
		sort.Sort(row)

		// set entire new row
		for j := 0; j < 500; j++ {
			m.Set(i, j, row.Row.At(0, j))
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
