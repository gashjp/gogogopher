package main

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/gashjp/gopher/encoder"
)

func main() {
	file, _ := os.Open("./img/gopher.png")
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
