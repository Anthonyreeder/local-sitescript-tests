package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Testing etc.
func Demo() {
	client := PokemonCenterClientSetup()
	//Set-Cookie is not working, I think its the format.
	//Also unable to set auth cookie in cookie jar like we do with datadome,
	//I think this is becuase golang is stripping out quotes and breaking the formatting.
	//So I set it directly in the header by passing it to each PokemonCenter task.
	//The addHeader function allows for direct cookie headers.
	authCookie := []string{"auth={\"access_token\":\"aa4a6061-b0a4-44b0-bae8-4bf81fea7489\",\"token_type\":\"bearer\",\"expires_in\":604799,\"scope\":\"pokemon\",\"role\":\"PUBLIC\",\"roles\":[\"PUBLIC\"]}"}

	//Must ensure that Datadome cookie (in helpers/setupClient) is up to date
	//Must ensure authCookie above is up to date

	PokemonCenterAddToCart(client, authCookie)            //tested and working, Currently hard coded to a product. Will be passed from monitor
	PokemonCenterSubmitAddressDetails(client, authCookie) //tested and working
	PokemonCenterSubmitPaymentDetails(client, authCookie)
	PokemonCenterCheckout(client, authCookie)
	//TODO
	//ATC Referal link
	//ATC get order ID from Product ID
	//Encrypt payment info
	//Encrypt:: Parse paymentKey from payment/key? <------- PaymentKey
	//Encrypt:: Use paymentKey as part of CyberSourceV2
	//Encrypt:: Pass encrpyted data too flex.cybersource to get the jweResponse
	//Encrypt:: Return the JTI string, this is the payment token <---------PaymentToken
	//Encrypt:: paymentDisplay = Visa 02/2026 <-----------PaymentDisplay
	//Submit payment info Load PaymentKey, PaymentToken and PaymentDisplay into json and post to payment API
	//Checkout order, Use response of above to get the 'URI'. Remove the junk, post to Order api
}

//Add to cart
func PokemonCenterAddToCart(client http.Client, directCookie []string) {

	payloadBytes, err := json.Marshal(PokemonCenterRequestAddToCart{ProductUri: "/carts/items/pokemon/qgqvhkjxga3c2mrzga2dq=/form", Quantity: 1, Configuration: ""})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?type=product&format=zoom.nodatalinks",
		Payload:  bytes.NewReader([]byte(payloadBytes)),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
}

//Submit billing and shipping info
func PokemonCenterSubmitAddressDetails(client http.Client, directCookie []string) {
	payloadBytes, err := json.Marshal(PokemonCenterRequestSubmitAddressDetails{
		Billing: Address{
			FamilyName:      "familyName",
			GivenName:       "givenName",
			StreetAddress:   "streetAddress",
			ExtendedAddress: "extendedAddress",
			Locality:        "locality",
			Region:          "region",
			PostalCode:      "12312312",
			CountryName:     "ga",
			PhoneNumber:     "13312423423",
		},
		Shipping: Address{
			FamilyName:      "familyName",
			GivenName:       "givenName",
			StreetAddress:   "streetAddress",
			ExtendedAddress: "extendedAddress",
			Locality:        "locality",
			Region:          "region",
			PostalCode:      "12312312",
			CountryName:     "ga",
			PhoneNumber:     "13312423423",
		},
	})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/address?format=zoom.nodatalinks",
		Payload:  bytes.NewReader([]byte(payloadBytes)),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
}

//Submit payment info
func PokemonCenterSubmitPaymentDetails(client http.Client, directCookie []string) {
	payloadBytes, err := json.Marshal(PokemonCenterRequestPaymentDetails{
		PaymentDisplay: "Visa 02/2026",
		PaymentKey:     "eyJraWQiOiJ3ZiIsImFsZyI6IlJTMjU2In0.eyJmbHgiOnsicGF0aCI6Ii9mbGV4L3YyL3Rva2VucyIsImRhdGEiOiJBL0wwUFVMTDIyeDg2U3pWb01neE54QUFFTFpPYzdkUjliQ0pqNWlYdDFFYjR4dFdvSEIrdHJYTE84OG0vUThYWThabUtDUnVMbUl6Q0cyQWw1c0FxeitoSWRVZFEzSVNkazBSdURsK2txZXlZMTRXWUxFMDBVZEwyTkdSb3F3cEdYUTMiLCJvcmlnaW4iOiJodHRwczovL2ZsZXguY3liZXJzb3VyY2UuY29tIiwiandrIjp7Imt0eSI6IlJTQSIsImUiOiJBUUFCIiwidXNlIjoiZW5jIiwibiI6Im1ONFM1cEhhbS1NczRreEstTnJ5blloS2V3YzJSWTB5VHkyX3NxcFdOQmRmZmJwWVBvcU9xREpQLXJwbWJDSkd4RXA1VWZlT1N6VllBcndLOVVZbkZyUWc1TzJzLUxHSmVQajdaXzZ0UTF2Z1JWczFWclNzdTFJR1ZtT2ZGRF9jN3pldDJRUno3TDljVDBmbkk2ZXBsS0pKNS1ZS2g3eEJYSl83ZmR3akJ0QXZNVDMzbm1Bb1B4WE1iQ3MwejlsX1VBMUtoRjU0Z1VwcHpmQkRWX3dYMW9GRjRVS0xTdTUzQmpaVjJUQkQ1VFNZdjBLZmhveVhSUXd5dGNMaVdhTjFuRndHd2g3QlJUODRGUDF2dnNLQkIyckExd3JTN3dmTjNQLTJrLWFGQmNlY3ZiZVRPUkdSdUNXVWU1bTRFeTd4OHJqbGtza1puUmZ0QjQ4YU1reGdNdyIsImtpZCI6IjAzTjlXQXBjQUoxeFFHQUN0ME5NWG5FVnVjanJXYUFtIn19LCJjdHgiOlt7ImRhdGEiOnsidGFyZ2V0T3JpZ2lucyI6WyJodHRwczovL3d3dy5wb2tlbW9uY2VudGVyLmNvbSIsImh0dHBzOi8vdGVzdC5wb2tlbW9uY2VudGVyLmNvbSIsImh0dHA6Ly9sb2NhbGhvc3Q6MzAwMCJdLCJtZk9yaWdpbiI6Imh0dHBzOi8vZmxleC5jeWJlcnNvdXJjZS5jb20ifSwidHlwZSI6Im1mLTAuMTEuMCJ9XSwiaXNzIjoiRmxleCBBUEkiLCJleHAiOjE2MjU1MDkxNTMsImlhdCI6MTYyNTUwODI1MywianRpIjoicFZxeExyam5tZ24wZGhFNSJ9.G5U8Dmb5TOfwWYdh5SIFKixzCrF41kf_B9Kx-Kkfs9x4KRwCN-7-zc01onRlsURL4mBRk9L_Uz3lreMgJoGPB6sxdW_iJv0Wg_zYEYFRwlMg6NuVewUyfyGwlDqmij62o88Z16MBhtJmJRhUCHzgRnIeVVv3BKXq_6b8xsVyZwkq-TpgiaCcRDw2NFOcaSfdWciIfWkNBz0Y60KcYU8E2ykQ-I5Wdu6svrMJmvftpPyyKnD9r_KXuBjmOtvXNUy3knIRxlyynbmspF1Q8DRgMM9nRCPogJlysgD7F8bJUnM0dHLZsqJr5epp_N73-yQ9PRrnuAalqtwNjssG3SVv_w",
		PaymentToken:   "1C0SIGRUV2U999XPDWYQJBO4FNW2O0DV48AQT8GT0B2GE8O5DSHE60E34D388CC3",
	},
	)
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/payment?microform=true&format=zoom.nodatalinks",
		Payload:  bytes.NewReader([]byte(payloadBytes)),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
}

//Checkout
func PokemonCenterCheckout(client http.Client, directCookie []string) {
	payloadBytes, err := json.Marshal(PokemonCenterRequestCheckoutDetails{
		PurchaseFrom: "/purchases/orders/pokemon/miywkmjsgq4dgljugi2wcljummygellcha2tmljxmiygeyrtgjsgiy3dgy=/form",
	},
	)
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/order?format=zoom.nodatalinks",
		Payload:  bytes.NewReader([]byte(payloadBytes)),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
}
