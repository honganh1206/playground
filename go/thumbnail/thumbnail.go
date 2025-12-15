package main

import (
	"os"
	"path/filepath"
	"strings"
)

func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // jpg, jpeg, etc.
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}

// Read an image from infile and write a thumbnail-sized version to outfile
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
}
