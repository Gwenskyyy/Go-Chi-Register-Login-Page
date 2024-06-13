package utils

import (
	"encoding/json" // Package encoding/json digunakan untuk encoding dan decoding data JSON
	"net/http" // Package net/http digunakan untuk operasi HTTP

	"github.com/go-chi/render" // Package render dari go-chi digunakan untuk rendering respons HTTP
)

// ErrorResponse adalah struktur untuk menyimpan informasi tentang respons error
type ErrorResponse struct {
	HTTPStatusCode int    `json:"-"` // Kode status HTTP untuk respons error
	StatusText     string `json:"status"` // Pesan status HTTP untuk respons error
	ErrorText      string `json:"error,omitempty"` // Pesan kesalahan spesifik untuk respons error
}

// Render adalah metode untuk menyiapkan respons error untuk dirender
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode) // Mengatur status HTTP pada respons
	return nil // Mengembalikan nil karena tidak ada kesalahan
}

// BadRequestResponse adalah fungsi untuk merespons permintaan yang tidak valid
func BadRequestResponse(w http.ResponseWriter, r *http.Request) {
	err := &ErrorResponse{ // Membuat objek ErrorResponse untuk Bad Request
		HTTPStatusCode: http.StatusBadRequest, // Mengatur kode status HTTP
		StatusText:     http.StatusText(http.StatusBadRequest), // Mendapatkan pesan status HTTP
		ErrorText:      "Bad request", // Pesan kesalahan spesifik
	}
	render.Render(w, r, err) // Merender respons error
}

// NotFoundResponse adalah fungsi untuk merespons permintaan untuk sumber daya yang tidak ditemukan
func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	err := &ErrorResponse{ // Membuat objek ErrorResponse untuk Not Found
		HTTPStatusCode: http.StatusNotFound, // Mengatur kode status HTTP
		StatusText:     http.StatusText(http.StatusNotFound), // Mendapatkan pesan status HTTP
		ErrorText:      "Not found", // Pesan kesalahan spesifik
	}
	render.Render(w, r, err) // Merender respons error
}

// InternalServerErrorResponse adalah fungsi untuk merespons kesalahan internal server
func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request) {
	err := &ErrorResponse{ // Membuat objek ErrorResponse untuk Internal Server Error
		HTTPStatusCode: http.StatusInternalServerError, // Mengatur kode status HTTP
		StatusText:     http.StatusText(http.StatusInternalServerError), // Mendapatkan pesan status HTTP
		ErrorText:      "Internal server error", // Pesan kesalahan spesifik
	}
	render.Render(w, r, err) // Merender respons error
}

// WriteJsonResponse adalah fungsi untuk menulis respons JSON ke ResponseWriter
func WriteJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json") // Mengatur tipe konten sebagai JSON
	w.WriteHeader(status) // Mengatur kode status HTTP pada respons
	json.NewEncoder(w).Encode(payload) // Mengkodekan payload menjadi JSON dan menulisnya ke ResponseWriter
}
