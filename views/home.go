package views

import (
	"formApp/config"      // Import package config untuk mengakses koneksi database dan session store
	"formApp/utils"       // Import package utils untuk mengakses utilitas yang membantu menangani respons HTTP dan kesalahan
	"html/template"       // Import package html/template untuk memanipulasi template HTML
	"net/http"            // Import package net/http untuk operasi HTTP
	"path"                // Import package path untuk manipulasi path file
)

// HomePage adalah handler untuk halaman utama aplikasi
func HomePage(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan sesi pengguna
	session, _ := config.Store.Get(r, "session-name")
	// Mendapatkan fullname dari sesi
	fullname, ok := session.Values["fullname"].(string)
	if !ok {
		// Jika fullname tidak ditemukan dalam sesi, maka pengguna belum login
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// Mendefinisikan path file template home.html
	var filepath = path.Join("pages/home.html")
	// Membaca template HTML dari file
	tmpl, err := template.ParseFiles(filepath)

	if err != nil {
		// Jika terjadi kesalahan saat parsing template, tangani dengan menampilkan pesan kesalahan internal server
		utils.InternalServerErrorResponse(w, r)
		return
	}

	// Menampilkan template HTML dengan data pengguna
	err = tmpl.Execute(w, struct {
		Fullname string
	}{
		Fullname: fullname,
	})
	if err != nil {
		// Jika terjadi kesalahan saat mengeksekusi template, tangani dengan menampilkan pesan kesalahan internal server
		utils.InternalServerErrorResponse(w, r)
		return
	}
}
