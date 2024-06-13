package config

import (
    "github.com/gorilla/sessions" // Mengimpor paket sessions dari Gorilla untuk manajemen sesi pengguna
)

// Store adalah variabel global yang digunakan untuk menyimpan instance dari sessions.CookieStore.
// sessions.CookieStore adalah tipe penyimpanan sesi yang menggunakan cookie untuk menyimpan data sesi pada sisi klien.
// Byte slice yang dilewatkan ke sessions.NewCookieStore digunakan sebagai seed untuk mengenkripsi dan menandatangani cookie sesi.
var Store = sessions.NewCookieStore([]byte("Form-App")) 
// Dalam hal ini, "Form-App" digunakan sebagai seed, bisa diganti sesuai kebutuhan aplikasi dengan nilai yang lebih aman.

// Menggunakan sessions.NewCookieStore() akan membuat sebuah instance dari CookieStore yang akan digunakan
// untuk menyimpan dan mengelola sesi pengguna dalam aplikasi web Go.
// Instance CookieStore ini dapat digunakan untuk membuat, mengakses, dan menghapus sesi pengguna,
// serta untuk menentukan pengaturan seperti durasi cookie, jalur cookie, dan domain cookie.
