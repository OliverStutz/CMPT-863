package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	gomp3 "github.com/hajimehoshi/go-mp3"
)

func decodeAll(t *testing.T, b []byte) []int16 {
	t.Helper()
	d, err := gomp3.NewDecoder(bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	raw, err := io.ReadAll(d)
	if err != nil {
		t.Fatal(err)
	}
	out := make([]int16, len(raw)/2)
	_ = binary.Read(bytes.NewReader(raw), binary.LittleEndian, out)
	return out
}
