package middleware

import (
	"net/http" // Package net/http digunakan untuk operasi HTTP
	"github.com/gorilla/sessions" // Package sessions dari Gorilla digunakan untuk manajemen sesi pengguna
)

// RequireLogin adalah middleware yang digunakan untuk memastikan pengguna telah login sebelum mengakses suatu halaman
func RequireLogin(store *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "session-name") // Mendapatkan sesi pengguna dari cookie store
			_, ok := session.Values["user_id"] // Memeriksa apakah terdapat nilai user_id di dalam sesi
			if !ok {
				http.Redirect(w, r, "/login", http.StatusFound) // Mengalihkan pengguna ke halaman login jika belum login
				return
			}
			next.ServeHTTP(w, r) // Membiarkan permintaan melanjutkan jika pengguna sudah login
		})
	}
}
