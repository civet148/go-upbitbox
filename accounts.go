package upbit

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Accounts struct {
	accessKey string
	secretKey string
}

func (acc Accounts) Sign(queryString string) string {
	payload := acc.createClaim(queryString)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))
	signedToken, err := token.SignedString([]byte(acc.secretKey))
	if err != nil {
		panic(errors.New("Signed fail: " + err.Error()))
	}

	return signedToken
}

func (acc Accounts) createClaim(queryString string) map[string]interface{} {
	payload := make(map[string]interface{})
	payload["access_key"] = acc.accessKey
	payload["nonce"] = uuid.New().String()
	payload["query"] = queryString

	queryHash := sha512.New()
	queryHash.Reset()
	queryHash.Write([]byte(queryString))

	payload["query_hash"] = hex.EncodeToString(queryHash.Sum(nil))
	payload["query_hash_alg"] = "SHA512"

	return payload
}
