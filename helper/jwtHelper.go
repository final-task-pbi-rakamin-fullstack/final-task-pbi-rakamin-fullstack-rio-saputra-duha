package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"goAPI/models"
	"os"
	"time"
)

type JWTClaims struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.Users) (string, error) {
	claims := JWTClaims{
		Uuid: user.Uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	//Get the secretKey
	secretKey := os.Getenv("SECRET_KEY")
	//Slice the secretKey from the salt
	slicedKey := secretKey[25 : len(secretKey)-25]
	//Generate Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(slicedKey))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ValidateToken(tokenString string) (*string, error) {

	//Get the secretKey
	secretKey := os.Getenv("SECRET_KEY")
	//Slice the secretKey from the salt
	slicedKey := secretKey[25 : len(secretKey)-25]
	//Check the token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(slicedKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("Invalid token signature")
		}
		return nil, errors.New("Invalid token")
	}
	//fmt.Println(token)

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	//fmt.Println(claims.Uuid)

	return &claims.Uuid, nil
}

//func GenerateSalt(length int) (string, error) {
//	// Create a byte slice with the specified length
//	b := make([]byte, length)
//
//	// Read random bytes into the byte slice
//	_, err := rand.Read(b)
//	if err != nil {
//		return "", err
//	}
//
//	// Encode the byte slice to a base64 string and return
//	return base64.URLEncoding.EncodeToString(b)[:length], nil
//}

func ExtractTokenString(saltedTokenString, salt string, secretKeyLength int) string {
	saltLength := len(salt)
	secretKey := saltedTokenString[saltLength : saltLength+secretKeyLength]
	return secretKey
}
