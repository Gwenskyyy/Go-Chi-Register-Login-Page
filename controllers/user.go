package controllers

import (
	"encoding/json"  // Package encoding/json digunakan untuk marshalling dan unmarshalling JSON
	"fmt"            // Package fmt digunakan untuk formatting string

	"github.com/gorilla/sessions" // Package sessions dari Gorilla digunakan untuk manajemen sesi pengguna
	"gorm.io/gorm"                // Package gorm digunakan untuk ORM (Object-Relational Mapping) pada Go
)

// UserController adalah controller untuk operasi-operasi yang berkaitan dengan pengguna (user)
type UserController struct { 
	db    *gorm.DB              // Variabel untuk menyimpan instance koneksi database
	store *sessions.CookieStore // Variabel untuk menyimpan instance sessions.CookieStore
}

// UserCreateResponse adalah respons JSON untuk operasi pembuatan pengguna
type UserCreateResponse struct {
	Username string `json:"username"` // Username pengguna yang dibuat
	Fullname string `json:"fullname"` // Nama lengkap pengguna yang dibuat
}

// UserLoginRequest adalah struktur untuk permintaan login pengguna
type UserLoginRequest struct {
	Username string `json:"username"` // Username pengguna yang melakukan login
	Password string `json:"password"` // Kata sandi pengguna yang melakukan login
}

// UserLoginResponse adalah respons JSON untuk operasi login pengguna
type UserLoginResponse struct {
	Message string `json:"message"` // Pesan sukses login
}

// NewUserController adalah fungsi untuk membuat instance UserController baru
func NewUserController(db *gorm.DB, store *sessions.CookieStore) *UserController {
	return &UserController{
		db:    db,    // Menginisialisasi db dengan instance koneksi database yang diberikan
		store: store, // Menginisialisasi store dengan instance sessions.CookieStore yang diberikan
	}
} 

// PrintJsonPayload adalah fungsi untuk mencetak payload JSON ke konsol
func PrintJsonPayload(payload interface{}) {
	jsonData, err := json.Marshal(payload) // Melakukan marshalling payload ke bentuk JSON
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData)) // Mencetak payload JSON ke konsol
}