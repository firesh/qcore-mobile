// Package mobile exposes the qore decoder to gomobile bindings.
package mobile

import (
	"encoding/json"
	"errors"
	"image"

	"github.com/firesh/qcore"
)

// DecodeGray decodes an 8-bit grayscale image and returns a JSON array of results.
// Points are in the input pixel coordinates. The caller owns image orientation.
func DecodeGray(pixels []byte, width, height, stride int) (string, error) {
	if width <= 0 || height <= 0 || stride < width || len(pixels) < width ||
		height-1 > (len(pixels)-width)/stride {
		return "", errors.New("invalid grayscale dimensions, stride, or buffer length")
	}
	img := &image.Gray{
		Pix: pixels, Stride: stride, Rect: image.Rect(0, 0, width, height),
	}
	results, err := qore.DecodeAll(img, nil)
	if err != nil {
		return "", err
	}
	if len(results) == 0 {
		return "[]", nil
	}
	output, err := json.Marshal(results)
	return string(output), err
}
