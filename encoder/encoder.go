package encoder

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/zyxar/image2ascii/ascii"
	"golang.org/x/crypto/ssh/terminal"
)

// JPGtoASCII convert image to ascii
func JPGtoASCII(buf *bytes.Buffer, img image.Image) error {
	width, height, err := getWH(img)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	if err := jpeg.Encode(&buff, img, nil); err != nil {
		fmt.Printf(err.Error())
		return err
	}

	_w, _h := width, height/2
	err = writeASCII(buff, buf, _w, _h)
	if err != nil {
		return err
	}

	return nil
}

// PNGtoASCII convert image to ascii
func PNGtoASCII(buf *bytes.Buffer, img image.Image) error {

	width, height, err := getWH(img)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	if err := png.Encode(&buff, img); err != nil {
		fmt.Printf(err.Error())
		return err
	}

	_w, _h := width, height/2
	err = writeASCII(buff, buf, _w, _h)
	if err != nil {
		return err
	}

	return nil
}

// getWH return (width, height, error)
func getWH(img image.Image) (int, int, error) {
	width, height := img.Bounds().Dx(), img.Bounds().Dy()
	if width == 0 || height == 0 {
		return 0, 0, fmt.Errorf("Error: width or height is 0, (%d,%d)", width, height)
	}
	tw, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, fmt.Errorf("terminal size syutoku error")
	}

	if err == nil && width > tw {
		width = tw
	}
	return width, height, nil
}

// writeASCII decode "from" to ASCII and write "to"
func writeASCII(from bytes.Buffer, to *bytes.Buffer, width, height int) error {
	a, err := ascii.Decode(&from, ascii.Options{Width: width, Height: height})
	if err != nil {
		return err
	}
	if _, err := a.WriteTo(to); err != nil {
		return err
	}
	return nil
}
