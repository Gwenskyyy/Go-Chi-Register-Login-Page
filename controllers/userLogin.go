package controllers

import (
	"fmt"            // Package fmt digunakan untuk formatting string
	"formApp/models" // Import package models untuk mengakses model-model aplikasi
	"formApp/utils"  // Import package utils untuk akses fungsi-fungsi utilitas
	"net/http"       // Package net/http digunakan untuk operasi HTTP
)

func (controller *UserController) UserLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		utils.BadRequestResponse(w, r) // Menangani permintaan yang tidak sesuai dengan metode POST
		return nil
	}

	if err := r.ParseForm(); err != nil {
		utils.BadRequestResponse(w, r) // Menangani kesalahan parsing formulir
		return err
	}

	username := r.FormValue("username") // Mendapatkan nilai username dari formulir
	password := r.FormValue("password") // Mendapatkan nilai password dari formulir

	var user models.User
	err := controller.db.Where("username = ?", username).First(&user).Error // Mencari pengguna berdasarkan username di database
	if err != nil {
		fmt.Println("User not found:", err)
		utils.WriteJsonResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "username / password is not match",
		}) // Menangani kasus ketika pengguna tidak ditemukan
		return nil
	}

	formPasswordHash := utils.HashPassword(password) // Mengenkripsi password dari formulir

	if !utils.ComparePassword(user.Password, formPasswordHash) {
		fmt.Println("Password comparison failed")
		utils.WriteJsonResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "username / password is not match",
		}) // Menangani kasus ketika password tidak cocok
		return nil
	}

	response := UserLoginResponse{ // Membuat respons login sukses
		Message: "Anda telah berhasil masuk",
	}
	session, _ := controller.store.Get(r, "session-name") // Mendapatkan atau membuat sesi baru
	session.Values["user_id"] = user.ID                   // Menyimpan ID pengguna ke dalam sesi
	session.Values["username"] = user.Username            // Menyimpan username pengguna ke dalam sesi
	session.Values["fullname"] = user.Fullname            // Menyimpan fullname pengguna ke dalam sesi
	session.Save(r, w)                                    // Menyimpan perubahan pada sesi

	PrintJsonPayload(response) // Mencetak respons JSON ke konsol

	http.Redirect(w, r, "/home", http.StatusFound) // Mengalihkan pengguna ke halaman home setelah login berhasil
	return nil
}