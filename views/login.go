package views

import (
	"formApp/config" // Import package config untuk mengakses koneksi database dan session store
	"formApp/controllers" // Import package controllers untuk mengakses fungsi yang mengatur logika aplikasi
	"formApp/utils" // Import package utils untuk mengakses utilitas yang membantu menangani respons HTTP dan kesalahan
	"html/template" // Import package html/template untuk memanipulasi template HTML
	"net/http" // Import package net/http untuk operasi HTTP
	"path" // Import package path untuk manipulasi path file
)

// LoginPage adalah handler untuk halaman login aplikasi
func LoginPage(w http.ResponseWriter, r *http.Request) {
	// Memeriksa apakah metode HTTP adalah POST
	if r.Method == http.MethodPost {
		// Membuat instance UserController dengan koneksi database dan session store
		userController := controllers.NewUserController(config.ConnectDb(), config.Store)
		// Memanggil fungsi UserLogin dari UserController untuk melakukan proses login pengguna
		if err := userController.UserLogin(w, r); err != nil {
			// Jika terjadi kesalahan, tangani dengan menampilkan pesan kesalahan internal server
			utils.InternalServerErrorResponse(w, r)
			return
		}
		return // Mengembalikan respons jika berhasil
	}

	// Mendefinisikan path file template login.html
	var filepath = path.Join("pages/login.html")
	// Membaca template HTML dari file
	tmpl, err := template.ParseFiles(filepath)

	// Memeriksa apakah terjadi kesalahan saat parsing template
	if utils.HandleError(w, err) {
		return
	}

	// Menangani template dengan fungsi utilitas HandleTemplate
	utils.HandleTemplate(w, tmpl)
}
