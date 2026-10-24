package image

import (
	"bytes"
	"image"
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"io"
	"os"

	webp "github.com/mayahiro/go-webp"
)

// SaveCompressedImage takes any io.Reader and saves it as a compressed WebP file to disk
func SaveCompressedImage(r io.Reader, outputPath string) error {
	// Decode any standard image format automatically from the stream
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Configure lossy compression and quality directly via Options
	options := &webp.Options{
		Compression: webp.CompressionLossy,
		Quality:     80,
	}

	// Compress and save straight to the destination file
	return webp.Encode(out, img, options)
}

// CompressImage takes any io.Reader (like a multipart.File) and returns WebP bytes
func CompressImage(r io.Reader) ([]byte, error) {
	// Decode any standard image format automatically from the stream
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	options := &webp.Options{
		Compression: webp.CompressionLossy,
		Quality:     80,
	}

	if err := webp.Encode(&buf, img, options); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}