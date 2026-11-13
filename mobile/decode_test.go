package mobile

import (
	"encoding/json"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestDecodeGray(t *testing.T) {
	f, err := os.Open("testdata/preview.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	src, err := png.Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	w, h := src.Bounds().Dx(), src.Bounds().Dy()
	stride := w + 13
	pixels := make([]byte, stride*(h-1)+w)
	for y := range h {
		for x := range w {
			pixels[y*stride+x] = color.GrayModel.Convert(src.At(x, y)).(color.Gray).Y
		}
	}
	got, err := DecodeGray(pixels, w, h, stride)
	if err != nil {
		t.Fatal(err)
	}
	var results []struct{ Text string }
	if err := json.Unmarshal([]byte(got), &results); err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 || results[0].Text != "qcore-mobile-preview" {
		t.Fatalf("unexpected results: %s", got)
	}
	if _, err := DecodeGray(pixels[:len(pixels)-1], w, h, stride); err == nil {
		t.Fatal("truncated frame accepted")
	}
	if got, err := DecodeGray(make([]byte, 64), 8, 8, 8); err != nil || got != "[]" {
		t.Fatalf("blank frame: %q, %v", got, err)
	}
}
