package controllers

import (
	"formApp/models" // Import package models untuk mengakses model-model aplikasi
	"formApp/utils"  // Import package utils untuk akses fungsi-fungsi utilitas
	"net/http"       // Package net/http digunakan untuk operasi HTTP
)

func (controller UserController) CreateUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		utils.BadRequestResponse(w, r) // Menangani permintaan yang tidak sesuai dengan metode POST
		return nil
	}

	if err := r.ParseForm(); err != nil {
		utils.InternalServerErrorResponse(w, r) // Menangani kesalahan parsing formulir
		return err
	}

	username := r.FormValue("username") // Mendapatkan nilai username dari formulir
	password := r.FormValue("password") // Mendapatkan nilai password dari formulir
	fullname := r.FormValue("fullname") // Mendapatkan nilai fullname dari formulir
	hashedPassword := utils.HashPassword(password) // Mengenkripsi password

	user := models.User{ // Membuat instance user baru
		Username: username,
		Password: hashedPassword,
		Fullname: fullname,
	} 

	err := controller.db.Create(&user).Error // Menyimpan user ke database 
	if err != nil {
		utils.InternalServerErrorResponse(w, r) // Menangani kesalahan saat membuat user
		return err
	}

	response := UserCreateResponse{ // Membuat respons JSON
		Username: user.Username,
		Fullname: user.Fullname,
	}

	PrintJsonPayload(response) // Mencetak respons JSON ke konsol

	http.Redirect(w, r, "/login", http.StatusFound) // Mengalihkan pengguna ke halaman login setelah membuat user

	return nil
}