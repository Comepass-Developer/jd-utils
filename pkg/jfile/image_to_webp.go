package jfile

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/chai2010/webp"
)

func ConvertImageToWebP(r io.Reader, ext string) ([]byte, error) {
	fmt.Println("Convert image to webp")
	var img image.Image
	var err error

	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(r)
	case ".png":
		img, err = png.Decode(r)
	case ".webp":
		return io.ReadAll(r) // skip convert
	default:
		return nil, fmt.Errorf("unsupported image format: %s", ext)
	}
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 80})
	return buf.Bytes(), err
}
