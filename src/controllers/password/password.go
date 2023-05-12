package password

import "golang.org/x/crypto/bcrypt"

// Fungsi untuk menghash password menggunakan algoritma BCrypt
func Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(bytes), err
}

// Fungsi untuk membandingkan password yang diinputkan dengan password yang sudah dihash
func Verify(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
