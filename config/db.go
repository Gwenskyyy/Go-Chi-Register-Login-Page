package config

import (
	"fmt" // Package fmt digunakan untuk formatting string
	"log" // Package log digunakan untuk logging
	"gorm.io/driver/postgres" // Package postgres dari GORM digunakan untuk koneksi ke PostgreSQL
	"gorm.io/gorm" // Package gorm digunakan untuk ORM (Object-Relational Mapping) pada Go
	"formApp/models" // Import package models untuk mengakses model-model aplikasi
)

// Konstanta untuk menyimpan informasi koneksi ke database PostgreSQL
const (
	PostgresHost     = "localhost" // Host dari database PostgreSQL
	PostgresPort     = 5432 // Port dari database PostgreSQL
	PostgresDb       = "formApp" // Nama database PostgreSQL
	PostgresUser     = "postgres" // Pengguna database PostgreSQL
	PostgresPassword = "12345" // Kata sandi untuk pengguna database PostgreSQL
)

var (
	db  *gorm.DB // Variabel untuk menyimpan instance koneksi database
	err error    // Variabel untuk menyimpan error saat terjadi koneksi atau migrasi model
)

// ConnectDb adalah fungsi untuk membuat koneksi ke database dan melakukan migrasi model
func ConnectDb() *gorm.DB {
	// Format string untuk konfigurasi koneksi ke database PostgreSQL
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDb)
	
	// Membuka koneksi ke database menggunakan konfigurasi yang telah ditentukan
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	// Memeriksa apakah terjadi error saat membuka koneksi
	if err != nil { 
		log.Fatalf("Error connecting to database: %v", err) // Log error jika koneksi gagal
	}

	// Melakukan migrasi model ke dalam database
	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error during migration: %v", err) // Log error jika migrasi model gagal
	}

	// Mengembalikan instance koneksi database
	return db
}
