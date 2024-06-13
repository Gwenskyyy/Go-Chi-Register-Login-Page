package routes

import (
	"formApp/config"      // Import package config untuk mendapatkan koneksi database dan cookie store
	"formApp/controllers" // Import package controllers untuk mengatur logika aplikasi
	"formApp/middleware"  // Import package middleware untuk menambahkan middleware ke router
	"formApp/views"       // Import package views untuk menangani tampilan aplikasi
	"net/http"            // Package net/http digunakan untuk operasi HTTP

	"github.com/go-chi/chi/v5" // Package chi digunakan untuk routing HTTP
)

// Routes adalah fungsi untuk mendefinisikan routing dalam aplikasi
func Routes() *chi.Mux {
	r := chi.NewRouter() // Membuat router baru menggunakan chi

	r.Get("/", views.IndexPage) // Menambahkan handler untuk halaman utama

	r.Route("/login", func(r chi.Router) { // Menentukan routing untuk halaman login
		r.Get("/", views.LoginPage)                                // Handler untuk tampilan halaman login
		r.Post("/", func(w http.ResponseWriter, r *http.Request) { // Handler untuk proses login
			controllers.NewUserController(config.ConnectDb(), config.Store).UserLogin(w, r) // Memanggil controller untuk proses login
		})
	})

	r.Route("/register", func(r chi.Router) { // Menentukan routing untuk halaman registrasi
		r.Get("/", views.RegisterPage)                             // Handler untuk tampilan halaman registrasi
		r.Post("/", func(w http.ResponseWriter, r *http.Request) { // Handler untuk proses registrasi
			controllers.NewUserController(config.ConnectDb(), config.Store).CreateUser(w, r) // Memanggil controller untuk proses registrasi
		})
	})

	r.Route("/home", func(r chi.Router) { // Menentukan routing untuk halaman beranda
		r.Use(middleware.RequireLogin(config.Store)) // Menambahkan middleware untuk memastikan pengguna telah login sebelum mengakses halaman beranda
		r.Get("/", views.HomePage)                   // Handler untuk tampilan halaman beranda
	})

	r.Route("/logout", func(r chi.Router) {
		r.Get("/", views.LogoutPage)
	})

	return r
}
