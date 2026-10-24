package image

import (
	"mime/multipart"
	"net/http"
)

func ValidateHeroImage(fileHeader *multipart.FileHeader) string {
	if fileHeader == nil {
		return "Hero image is required."
	}

	// Check file size (e.g., max 5MB)
	if fileHeader.Size > 10<<20 {
		return "Image size must be less than 5MB."
	}

	// Check content type
	file, err := fileHeader.Open()
	if err != nil {
		return "Unable to read uploaded file."
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "Unable to verify file type."
	}

	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/webp" {
		return "Invalid image format. Only JPEG, PNG, and WebP are allowed."
	}

	return "" // No error
}
