package bcrypt

import libbcrypt "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytes, err := libbcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func Compare(password, hash string) bool {
	return libbcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
