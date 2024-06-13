package views

import (
	"formApp/utils" // Import package utils untuk mengakses utilitas yang membantu menangani kesalahan dan respons HTTP
	"html/template" // Import package html/template untuk memanipulasi template HTML
	"net/http" // Import package net/http untuk operasi HTTP
	"path" // Import package path untuk manipulasi path file
)

// IndexPage adalah handler untuk halaman indeks (index page) aplikasi
func IndexPage(w http.ResponseWriter, r *http.Request) {
	// Mendefinisikan path file template index.html
	var filepath = path.Join("pages/index.html")
	// Membaca template HTML dari file
	tmpl, err := template.ParseFiles(filepath)

	// Memeriksa apakah terjadi kesalahan saat parsing template
	if utils.HandleError(w, err) {
		return
	}

	// Menangani template dengan fungsi utilitas HandleTemplate
	utils.HandleTemplate(w, tmpl)
}
