package home

import (
	"log/slog"
	"net/http"

	"github.com/raffy-io/taskify/internal/image"
)

type Handler struct {
	Queries *Queries
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := Index().Render(ctx, w); err != nil {
		http.Error(w, "Failed to render the template..", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	form := Form{}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := Edit(form).Render(ctx, w); err != nil {
		http.Error(w, "Failed to render the template..", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. Enforce body size limit and parse multipart form (10MB max)
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "The uploaded file is too large or invalid (Max 10MB).", http.StatusBadRequest)
		return
	}
	defer r.MultipartForm.RemoveAll()

	// 2. Bind form fields
	form := Form{
		HeroImageAlt: r.FormValue("hero_image_alt"),
		HeroTitle:    r.FormValue("hero_title"),
		HeroText:     r.FormValue("hero_text"),
	}

	// 3. Validate text fields
	isValid := form.ValidateText()

	// 4. Handle file upload validation
	 _, fileHeader, err := r.FormFile("hero_image") // just get the header for validation skip the image file
	 if err == http.ErrMissingFile {
		form.AddError("hero_image", "Hero image is required.")
		isValid = false
	} else if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else if fileHeader != nil {
		imgErr := image.ValidateHeroImage(fileHeader)
		if imgErr != "" {
			form.AddError("hero_image", imgErr)
			isValid = false
		}
	}

	// 5. Render form with errors if invalid
	if !isValid {
		w.WriteHeader(http.StatusOK) // Must be 200 so HTMX performs the outerHTML swap
		EditForm(form).Render(ctx, w)
		return
	}

	// Grab both the file stream and header (ignore the header if you already validated size/type)
	file, _, err := r.FormFile("hero_image")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return  
	}
	defer file.Close()

	// Pass the file stream directly into your compressor
	webpImage, err := image.CompressImage(file)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return 
	}

	// Here will be the logic to save in database

	slog.Info("form submitted successfully %v",slog.Int("size_bytes", len(webpImage)))
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := Test().Render(ctx, w); err != nil {
		http.Error(w, "Failed to render the template..", http.StatusInternalServerError)
		return
	}
}
