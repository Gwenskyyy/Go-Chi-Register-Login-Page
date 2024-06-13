package utils

import (
	"crypto/md5" // Package md5 digunakan untuk melakukan hashing menggunakan algoritma MD5, SHA256?
	"encoding/hex" // Package hex digunakan untuk melakukan encoding hexadecimal
)

// HashPasswordMD5 melakukan hashing password menggunakan algoritma MD5.
func HashPassword(password string) string {
	hasher := md5.New() // Membuat objek hasher untuk algoritma MD5
	hasher.Write([]byte(password)) // Menulis password ke dalam hasher
	return hex.EncodeToString(hasher.Sum(nil)) // Mengembalikan hasil hashing dalam bentuk string hexadecimal
}

// ComparePasswordMD5 membandingkan password yang di-hash dengan password teks biasa.
func ComparePassword(hashedPassword, plainPassword string) bool {
	return hashedPassword == HashPassword(plainPassword) // Membandingkan password yang di-hash dengan password teks biasa
}
