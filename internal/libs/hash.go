package libs

import "golang.org/x/crypto/bcrypt"

func HashPass(originalPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePass(originalPassword, databaseHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(databaseHash), []byte(originalPassword))
}
