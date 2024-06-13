package main

import (
	"fmt"
	"formApp/config"
	"formApp/routes"
	"net/http"
)

func main() {
	// Membuat koneksi ke database saat aplikasi dimulai
	config.ConnectDb()

	// Menetapkan rute aplikasi

	r := routes.Routes()

	// Menampilkan pesan bahwa server telah dimulai
	fmt.Println("Server started at localhost:8080")

	// Start server
	if err := http.ListenAndServe(":8080", r); err != nil {
		// Handle server startup error
		panic(err)
	} // Mendengarkan permintaan HTTP pada port 8080 menggunakan router yang telah ditentukan
}
