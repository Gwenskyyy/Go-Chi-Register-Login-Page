package views

import (
	"formApp/config" // Import package config untuk mengakses koneksi database dan session store
	"formApp/controllers" // Import package controllers untuk mengakses fungsi yang mengatur logika aplikasi
	"formApp/utils" // Import package utils untuk mengakses utilitas yang membantu menangani respons HTTP dan kesalahan
	"html/template" // Import package html/template untuk memanipulasi template HTML
	"net/http" // Import package net/http untuk operasi HTTP
	"path" // Import package path untuk manipulasi path file
)

// RegisterPage adalah handler untuk halaman pendaftaran (register page) aplikasi
func RegisterPage(w http.ResponseWriter, r *http.Request) {
	// Memeriksa apakah metode HTTP adalah POST
	if r.Method == http.MethodPost {
		// Memanggil fungsi CreateUser dari UserController untuk membuat pengguna baru
		if err := controllers.NewUserController(config.ConnectDb(), config.Store).CreateUser(w, r); err != nil {
			// Jika terjadi kesalahan, tangani dengan menampilkan pesan kesalahan internal server
			utils.InternalServerErrorResponse(w, r)
			return
		}
		return // Mengembalikan respons jika berhasil
	}

	// Mendefinisikan path file template register.html
	var filepath = path.Join("pages/register.html")
	// Membaca template HTML dari file
	tmpl, err := template.ParseFiles(filepath)

	// Memeriksa apakah terjadi kesalahan saat parsing template
	if utils.HandleError(w, err) {
		return
	}

	// Menangani template dengan fungsi utilitas HandleTemplate
	utils.HandleTemplate(w, tmpl)
}
