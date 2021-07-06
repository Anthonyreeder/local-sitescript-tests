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
	"unicode/utf8"

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
func CyberSourceV2(keyId string) {
	key := strings.Split(keyId, ".")[1]

	decodedKeyBytes, _ := base64.StdEncoding.DecodeString(key)
	decodedKeyString := string(decodedKeyBytes)

	fmt.Println(utf8.RuneCountInString(decodedKeyString))
	fmt.Println("Key: " + decodedKeyString)
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

	b, err := json.Marshal(header_)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println((string(b)))

	card_ := new(Card)
	card_.SecurityCode = "260"
	card_.Number = "4767718212263745"
	card_.Type = "002"
	card_.ExpMonth = "02"
	card_.ExpYear = "2026"

	encryptedObject_ := new(EncryptedObject)
	encryptedObject_.Context = keyId
	encryptedObject_.Index = 0
	encryptedObject_.Data = *card_

	b_, err_ := json.Marshal(encryptedObject_)

	//eDecoded, _ := base64.StdEncoding.DecodeString(encrypt.Flx.Jwk.E)
	//nDecoded, _ := base64.StdEncoding.DecodeString(encrypt.Flx.Jwk.N)
	//var publicKey = PublicKey{N: eDecoded}

	if err_ != nil {
		fmt.Println(err_)
	}
	fmt.Println((string(b_)))

	//Get RSA
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
	//fmt.Println(set)
	for it := set.Iterate(context.Background()); it.Next(context.Background()); {
		pair := it.Pair()
		key := pair.Value.(jwk.Key)

		var rawkey interface{} // This is the raw key, like *rsa.PrivateKey or *ecdsa.PrivateKey
		if err := key.Raw(&rawkey); err != nil {
			log.Printf("failed to create public key: %s", err)
			return
		}

		// We know this is an RSA Key so...
		rsa___, ok := rawkey.(*rsa.PublicKey)
		fmt.Println(rsa___)
		if !ok {
			panic(fmt.Sprintf("expected ras key, got %T", rawkey))
		}
		// As this is a demo just dump the key to the console
		payload := `{
			"context": "` + keyId + `",
			"index": 0,
			"data":{
				"securityCode":"260",
				"number":"4767718212263745",
				"type":"002",
				"expMonth":"02",
				"expYear":"2026"
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
		fmt.Println(h_map)
		headerMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(h_map), &headerMap)
		if err != nil {
			panic(err)
		}
		dumpMap("", headerMap)
		//fmt.Println(rsa.Size())
		/*z := new(big.Int)
		z.SetBytes(eDecoded)

		byteToInt, _ := strconv.Atoi(string(nDecoded))

		p_key := rsa.PublicKey{
			N: z,
			E: byteToInt,
		}*/
		fmt.Println(utf8.RuneCountInString(payload))
		token__, err__ := jose.Encrypt(payload, jose.RSA_OAEP, jose.A256GCM, rsa___, jose.Headers(headerMap))
		if err__ != nil {
			fmt.Println(err__)
		}
		//fmt.Println(rsa___.N.BitLen())
		fmt.Println(utf8.RuneCountInString(token__))
		fmt.Println(token__)
	}
	// Header, Rsa, encryptedObject --> found

	//Let's encrypt

}

type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
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
