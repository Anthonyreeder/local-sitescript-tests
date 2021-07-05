package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

func CyberSourceV2(keyId string) {
	key := strings.Split(keyId, ".")[1]
	decodedKeyBytes, _ := base64.StdEncoding.DecodeString(key)
	decodedKeyString := string(decodedKeyBytes)

	var encrypt Encrypt
	json.Unmarshal([]byte(decodedKeyString), &encrypt)

	eDecoded, _ := base64.StdEncoding.DecodeString(encrypt.Flx.Jwk.E)
	nDecoded, _ := base64.StdEncoding.DecodeString(encrypt.Flx.Jwk.N)

	var publicKey = PublicKey{N: eDecoded}
	fmt.Println(decodedKeyString)
}

type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
}

type Encrypt struct {
	Flx struct {
		Jwk struct {
			Kty string
			E   string
			Use string
			N   string
			Kid string
		}
	}
}
