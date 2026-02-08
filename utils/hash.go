package utils

import "golang.org/x/crypto/bcrypt"

// パスワードをbcryptでハッシュ化させる
func GeneratedHash(passwd string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// DBに登録してあるハッシュと入力されたパスワードを照合する
func VeryifyPassword(hashedPassword, passwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwd)) == nil
}


