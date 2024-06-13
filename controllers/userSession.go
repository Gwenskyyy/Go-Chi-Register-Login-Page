package controllers

import (
	"fmt"            // Package fmt digunakan untuk formatting string
	"formApp/models" // Import package models untuk mengakses model-model aplikasi
	"formApp/utils"  // Import package utils untuk akses fungsi-fungsi utilitas
	"net/http"       // Package net/http digunakan untuk operasi HTTP
)

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		utils.BadRequestResponse(w, r) // Menangani permintaan yang tidak sesuai dengan metode GET
		return nil
	}

	session, _ := controller.store.Get(r, "session-name") // Mendapatkan sesi pengguna
	userID, ok := session.Values["user_id"].(uint)        // Mendapatkan ID pengguna dari sesi
	if !ok {
		utils.WriteJsonResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "user not authenticated",
		}) // Menangani kasus ketika pengguna belum terautentikasi
		return nil
	}
	var user models.User
	err := controller.db.First(&user, userID).Error // Mengambil informasi pengguna dari database berdasarkan ID
	if err != nil {
		fmt.Println("User not found:", err)
		utils.WriteJsonResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "user not found",
		}) // Menangani kasus ketika pengguna tidak ditemukan di database
		return nil
	}

	utils.WriteJsonResponse(w, http.StatusOK, user) // Mengirimkan informasi pengguna dalam respons JSON
	return nil

}