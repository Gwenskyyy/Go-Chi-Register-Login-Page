package views

import (
    "net/http"
    "formApp/config" // Import package config untuk mengakses session store
)

// LogoutPage adalah handler untuk proses logout
func LogoutPage(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan sesi pengguna
    session, _ := config.Store.Get(r, "session-name")
    // Menghapus semua nilai dalam sesi
    session.Values = make(map[interface{}]interface{})
    // Menyimpan perubahan sesi
    session.Save(r, w)
    // Redirect ke halaman login setelah logout
    http.Redirect(w, r, "/login", http.StatusFound)
}