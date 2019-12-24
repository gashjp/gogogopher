package main

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/gashjp/gogogopher/encoder"
	_ "github.com/gashjp/gogogopher/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	// file, _ := os.Open("img/gopher.png")
	fs, err := fs.New()
	if err != nil {
		fmt.Printf(err.Error() + " a\n")
		return
	}

	file, err := fs.Open("/gopher.png")
	if err != nil {
		fmt.Printf(err.Error() + " b\n")
		return
	}

	src, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		return
	}

	var buf bytes.Buffer
	err = encoder.PNGtoASCII(&buf, src)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	os.Stdout.Write(buf.Bytes())
	os.Stdout.Sync()
}
