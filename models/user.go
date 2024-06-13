package models

import (
	"github.com/asaskevich/govalidator" // Package govalidator digunakan untuk validasi data
	"gorm.io/gorm" // Package gorm digunakan untuk ORM (Object-Relational Mapping) pada Go
	"formApp/utils" // Import package utils untuk akses fungsi-fungsi utilitas
)

// User adalah struktur yang merepresentasikan entitas pengguna dalam aplikasi
type User struct {
	gorm.Model // Menggunakan gorm.Model untuk mendapatkan kolom-kolom standar seperti ID, CreatedAt, UpdatedAt, dan DeletedAt
	Username string `gorm:"not null;uniqueIndex" json:"username,omitempty" form:"username" valid:"required~Your username is required"` // Kolom untuk menyimpan username pengguna
	Password string `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"` // Kolom untuk menyimpan kata sandi pengguna
	Fullname string `gorm:"not null" json:"fullname,omitempty" form:"fullname" valid:"required~Your name is required"` // Kolom untuk menyimpan nama lengkap pengguna
}

// BeforeCreate adalah hook yang dipanggil sebelum membuat model baru di database
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
_, errCreate := govalidator.ValidateStruct(user) // Melakukan validasi struktur pengguna sebelum dibuat 
	if errCreate != nil {
		err = errCreate // Jika terdapat kesalahan validasi, mengembalikan kesalahan
		return err	
	} 

	hash := utils.HashPassword(user.Password) // Mengenkripsi kata sandi pengguna sebelum disimpan di database
	user.Password = hash // Menyimpan kata sandi terenkripsi ke dalam kolom Password
	return // Mengembalikan nil karena tidak terdapat kesalahan
}

