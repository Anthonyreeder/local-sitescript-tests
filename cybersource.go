package main

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/lestrrat-go/jwx/jwk"
)

type RSA struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N   string `json:"n"`
}

type Header__ struct {
	Kid string `json:"kid"`
	Jwk RSA    `json:"jwk"`
}

type EncryptedObject struct {
	Context string `json:"context"`
	Index   int    `json:"index"`
	Data    Card   `json:"data"`
}

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}
func Base64ToInt(s string) (*big.Int, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetBytes(data)
	return i, nil
}
func CyberSourceV2(keyId string) (returnVal string) {
	key := strings.Split(keyId, ".")[1]

	decodedKeyBytes, _ := base64.StdEncoding.DecodeString(key)
	decodedKeyString := string(decodedKeyBytes)

	var encrypt Encrypt
	json.Unmarshal([]byte(decodedKeyString), &encrypt)

	kid := string(encrypt.Flx.Jwk.Kid)

	e := string(encrypt.Flx.Jwk.E)
	n := string(encrypt.Flx.Jwk.N)
	kty := string(encrypt.Flx.Jwk.Kty)
	use := string(encrypt.Flx.Jwk.Use)

	rsa_ := new(RSA)
	rsa_.Kid = kid
	rsa_.E = e
	rsa_.Kty = kty
	rsa_.N = n
	rsa_.Use = use

	header_ := new(Header__)
	header_.Kid = kid
	header_.Jwk = *rsa_

	card_ := new(Card)
	card_.SecurityCode = "260"
	card_.Number = "4767718212263745"
	card_.Type = "001"
	card_.ExpMonth = "02"
	card_.ExpYear = "2026"

	encryptedObject_ := new(EncryptedObject)
	encryptedObject_.Context = keyId
	encryptedObject_.Index = 0
	encryptedObject_.Data = *card_

	jwkJSON := `{
		"keys": [ 
		  {
			"kty": "` + kty + `",
			"n": "` + n + `",
			"use": "` + use + `",
			"alg": "RSA-OAEP",
			"e": "` + e + `",
			"kid": "` + kid + `"
		  }
		]
	  }
	  `

	set, err := jwk.Parse([]byte(jwkJSON))
	if err != nil {
		panic(err)
	}

	for it := set.Iterate(context.Background()); it.Next(context.Background()); {
		pair := it.Pair()
		key := pair.Value.(jwk.Key)

		var rawkey interface{}
		if err := key.Raw(&rawkey); err != nil {
			log.Printf("failed to create public key: %s", err)
			return
		}

		rsa___, ok := rawkey.(*rsa.PublicKey)

		if !ok {
			panic(fmt.Sprintf("expected ras key, got %T", rawkey))
		}

		payload := `{
			"context": "` + keyId + `",
			"index": 0,
			"data":{
				"securityCode":"260",
				"number":"4767718212263745",
				"type":"001",
				"expirationMonth":"02",
				"expirationYear":"2026"
			}
		}`

		h_map := `{
			"kid":"` + kid + `",
			"jwk":{
				"kty":"` + kty + `",
				"e":"` + e + `",
				"use":"` + use + `",
				"kid":"` + kid + `",
				"n":"` + n + `"
			}
		}`

		headerMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(h_map), &headerMap)
		if err != nil {
			panic(err)
		}
		dumpMap("", headerMap)

		token__, err__ := jose.Encrypt(payload, jose.RSA_OAEP, jose.A256GCM, rsa___, jose.Headers(headerMap))
		if err__ != nil {
			fmt.Println(err__)
		}
		returnVal = token__

	}
	return returnVal
}
func retrievePaymentToken(keyId string) (jti string) {
	key := strings.Split(keyId, ".")[1]
	decodedKeyBytes, _ := base64.StdEncoding.DecodeString(key)
	decodedKeyString_ := string(decodedKeyBytes) + "}"
	fmt.Println(decodedKeyString_)
	var encrypt PaymentToken
	if err := json.Unmarshal([]byte(decodedKeyString_), &encrypt); err != nil {
		fmt.Println(err)
	}
	return encrypt.Jti
}

type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
}
type PaymentToken struct {
	Jti string `json:"jti"`
}

type Card struct {
	SecurityCode string `json:"securityCode"`
	Number       string `json:"number"`
	Type         string `json:"type"`
	ExpMonth     string `json:"expMonth"`
	ExpYear      string `json:"expYear"`
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
