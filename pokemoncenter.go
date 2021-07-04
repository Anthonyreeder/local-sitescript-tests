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
	authCookie := []string{"auth={\"access_token\":\"6fa83426-6181-4413-ad27-9ac601aa3232\",\"token_type\":\"bearer\",\"expires_in\":604799,\"scope\":\"pokemon\",\"role\":\"PUBLIC\",\"roles\":[\"PUBLIC\"]}"}

	//Must ensure that Datadome cookie (in helpers/setupClient) is up to date
	//Must ensure authCookie above is up to date

	PokemonCenterAddToCart(client, authCookie)            //tested and working, Currently hard coded to a product. Will be passed from monitor
	PokemonCenterSubmitAddressDetails(client, authCookie) //tested and working
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
