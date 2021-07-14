package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//Testing etc.
func Demo() {
	client := PokemonCenterClientSetup()
	//Set-Cookie is not working with this cookie, I think its the format.
	//Also unable to set auth cookie in cookie jar like we do with datadome,
	//I think this is becuase golang is stripping out quotes and breaking the formatting.
	//So I set it directly in the header by passing it to each PokemonCenter task.
	//The addHeader function allows  for direct cookie headers.
	authCookie := []string{"auth={\"access_token\":\"59400864-da63-4fd3-b875-98ffb9a0cd3d\",\"token_type\":\"bearer\",\"expires_in\":604799,\"scope\":\"pokemon\",\"role\":\"PUBLIC\",\"roles\":[\"PUBLIC\"]}"}

	//Must ensure that Datadome cookie (in helpers/setupClient) is up to date
	//Must ensure authCookie above is up to date
	PokemonCenterGetAuthId(client, []string{})
	//get auth cookie
	u, _ := url.Parse("http://www.pokemoncenter.com")
	test := client.Jar.Cookies(u)
	println(test[len(test)-1])
	PokemonCenterStockCheck(client, authCookie, "qgqvhkjxgazs2mbvgqyds")
	PokemonCenterAddToCart(client, authCookie)
	PokemonCenterSubmitAddressDetailsValidate(client, authCookie) //will tell u if there are any issues with address
	PokemonCenterSubmitAddressDetails(client, authCookie)
	rawKeyId := PokemonCenterGetPaymentKeyId(client, authCookie)

	//Extract KeyID
	var pokemonCenterResponseKeyId PokemonCenterResponseKeyId
	json.Unmarshal([]byte(rawKeyId), &pokemonCenterResponseKeyId)

	//CyberSourceV2(pokemonCenterResponseKeyId.KeyId)
	//
	//replace paymentKey and paymentToken with values returned from CyberSourceV2
	paymentKey := "eyJraWQiOiJ3ZiIsImFsZyI6IlJTMjU2In0.eyJmbHgiOnsicGF0aCI6Ii9mbGV4L3YyL3Rva2VucyIsImRhdGEiOiIzeDdGRjZCUUhrWkt1TnlpQTNMQ2d4QUFFSmk0SGh0MjZMSVg1SnUrc1ZVaGY2UUFXQkFJZ2RnNktCbUJrS3RuUnM4YlBmb1VrUzBXeTYvTVMxMzZNLzlyZ043U1FmZFJtNFo4eXRJYVpOWDByVDdpVjhKcDBuYTFPZHlBYzBaSTk5UWkiLCJvcmlnaW4iOiJodHRwczovL2ZsZXguY3liZXJzb3VyY2UuY29tIiwiandrIjp7Imt0eSI6IlJTQSIsImUiOiJBUUFCIiwidXNlIjoiZW5jIiwibiI6InBVdHdpcGJrT181VnJYeXlOMUI1T21hWVM5UmxkRzJZWHFGYmhlMkJkVHBwdjBfdGtONFhIYkNmNkhKWjl2eVdwRURnYllGUTdsSGlOYnF0UXpWWEUtbDVGdTVhMlBES3N2d0Ryek1kWkQ4R0liU0phaUs0U0RaNERNX3hsVlgwYXBfdm9rVTFyZDJYN0o0MHE0RnUzMUhycURIQ0VCcUVuZFk4MEx4Q3hCX09nMzlhckRmZVpEVmFrNE1tcDJHTFluaG9lM1pOMUFXWm00R0lnNVZFQ0RUZGRPN3VWb05UUWVyVFFpcWxBN3pIQ1lOaXVjYkpqdlE0RUxwSkQyQkdzc0lJSHJjcERfbmVoTmp6NklHVm51VzJ2N25oMGluQUtVTFEyMmRDVU9KUDlkTHhUcC05NG04Z2FXVGd0TWljbHFadWhOX2tnbXRiLURTY21FV3VsdyIsImtpZCI6IjAzelRvU25STGdXN0xNRHZiY2RtUjFZY1ZvV1d4bkFKIn19LCJjdHgiOlt7ImRhdGEiOnsidGFyZ2V0T3JpZ2lucyI6WyJodHRwczovL3d3dy5wb2tlbW9uY2VudGVyLmNvbSIsImh0dHBzOi8vdGVzdC5wb2tlbW9uY2VudGVyLmNvbSIsImh0dHA6Ly9sb2NhbGhvc3Q6MzAwMCJdLCJtZk9yaWdpbiI6Imh0dHBzOi8vZmxleC5jeWJlcnNvdXJjZS5jb20ifSwidHlwZSI6Im1mLTAuMTEuMCJ9XSwiaXNzIjoiRmxleCBBUEkiLCJleHAiOjE2MjU1OTA1MzgsImlhdCI6MTYyNTU4OTYzOCwianRpIjoicDlncU1Wc0xzTGtJOFdIcSJ9.kM6YfQbCK4w4U1SIppPExDQjLmXUoO6JOTFD9JhcKyueNf7F_d5JCYkmeIg_u0BRx5pmkANbSn_ODxYgkfWZPqI6_TR3CTR-BtuZJRhejiciNwYmwVB_NuIjggYTSLEEnumb617otxLalxt7Fyp6LrU0p_BwgPG4Hx9txdPM2AWySaYb4UT57wcITmBE5Yanr_CoXLj0iCnZC43wvnkfdQ6ZHhO7ErRHSghsVT1Yh7D-JsSQ-J7s3oTDpMk-kV5mF4t7klwJcikyOpsr6FTzed9Bn2c-Vnl4F_lkCITDybxlPQcuZnX-MSi_Z4C9HW3fFd2P6q1SraGUSRA8LidH5w"
	paymentToken := "1C3J1NO3TQJO4TKRDEGZXP42TENC1Z8Q9IH776B50UC08F7I1D1N60E48B1DCBF2"

	//paymentToken := "1C0SIGRUV2U999XPDWYQJBO4FNW2O0DV48AQT8GT0B2GE8O5DSHE60E34D388CC3"

	paymentBytes := PokemonCenterSubmitPaymentDetails(client, authCookie, paymentKey, paymentToken)

	//we need to get the return value from payment to pass to checkout
	var pokemonCenterResponseSetPayment = PokemonCenterResponseSetPayment{}
	json.Unmarshal(paymentBytes, &pokemonCenterResponseSetPayment)

	checkoutPayload := strings.Replace(pokemonCenterResponseSetPayment.Self.Uri, "paymentmethods", "purchases", 1) + "/form"

	PokemonCenterCheckout(client, authCookie, checkoutPayload)

	//Todo: Set AUTHID Organically.
}

func PokemonCenterLogin(client http.Client, directCookie []string) []byte {
	payloadBytes, err := json.Marshal(PokemonCenterLoginRequest{
		Username:   "anthonyreeder123@gmail.com",
		Password:   "thekid225",
		Grant_type: "password",
		Role:       "REGISTERED",
		Scope:      "pokemon",
	},
	)
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/auth?format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	respBytes, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
	return respBytes
}

func PokemonCenterStockCheck(client http.Client, directCookie []string, product string) bool {
	payloadBytes, err := json.Marshal(PokemonCenterRequestAddToCart{ProductUri: "/carts/items/pokemon/" + product + "=/form", Quantity: 1, Configuration: ""})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?type=product&format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	responseBytes, _ := PokemonCenterNewResponse(client, request)

	pokemonCenterStockCheckResponse := PokemonCenterStockCheckResponse{}
	json.Unmarshal(responseBytes, &pokemonCenterStockCheckResponse)

	if strings.Contains(pokemonCenterStockCheckResponse.Self.Type, "error") {
		if strings.Contains(pokemonCenterStockCheckResponse.Self.Id, "item.not.available") {
			//out of stock
			//in production we can return oos or 'other error' depending on this 'ID' value.
			return false
		} else {
			//Not out of stock but there was an error adding the item to cart.
			return false
		}
	} else if strings.Contains(pokemonCenterStockCheckResponse.Self.Type, "carts.line-item") {
		return true
		//in stock
	} else {
		//unknown type
		return false
	}
}

//Add to cart
func PokemonCenterAddToCart(client http.Client, directCookie []string) {

	payloadBytes, err := json.Marshal(PokemonCenterRequestAddToCart{ProductUri: "/carts/items/pokemon/qgqvhkjxga3c2mrzga2dq=/form", Quantity: 1, Configuration: ""})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?type=product&format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	//Please enable JS and disable any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)
}

//must be called before submit address
func PokemonCenterSubmitAddressDetailsValidate(client http.Client, directCookie []string) {
	payloadBytes, err := json.Marshal(PokemonCenterRequestSubmitAddressDetails{
		Billing: Address{
			FamilyName:      "Reeder",
			GivenName:       "Ant",
			StreetAddress:   "1301 Reisling Ct",
			ExtendedAddress: "",
			Locality:        "Las Vegas",
			Region:          "NV",
			PostalCode:      "89144",
			CountryName:     "US",
			PhoneNumber:     "+1 (342) 342-3423",
		},
		Shipping: Address{
			FamilyName:      "Reeder",
			GivenName:       "Ant",
			StreetAddress:   "1301 Reisling Ct",
			ExtendedAddress: "",
			Locality:        "Las Vegas",
			Region:          "NV",
			PostalCode:      "89144",
			CountryName:     "US",
			PhoneNumber:     "+1 (342) 342-3423",
		},
	})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/address/validate",
		Payload:  bytes.NewReader(payloadBytes),
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
			FamilyName:      "Reeder",
			GivenName:       "Ant",
			StreetAddress:   "1301 Reisling Ct",
			ExtendedAddress: "",
			Locality:        "Las Vegas",
			Region:          "NV",
			PostalCode:      "89144",
			CountryName:     "US",
			PhoneNumber:     "+1 (342) 342-3423",
		},
		Shipping: Address{
			FamilyName:      "Reeder",
			GivenName:       "Ant",
			StreetAddress:   "1301 Reisling Ct",
			ExtendedAddress: "",
			Locality:        "Las Vegas",
			Region:          "NV",
			PostalCode:      "89144",
			CountryName:     "US",
			PhoneNumber:     "+1 (342) 342-3423",
		},
	})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/address?format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
}

//Submit payment info
func PokemonCenterSubmitPaymentDetails(client http.Client, directCookie []string, paymentKey string, paymentToken string) []byte {
	payloadBytes, err := json.Marshal(PokemonCenterRequestPaymentDetails{
		PaymentDisplay: "Visa 02/2026",
		PaymentKey:     paymentKey,
		PaymentToken:   paymentToken,
	},
	)
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/payment?microform=true&format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	respBytes, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
	return respBytes
}

//Checkout
func PokemonCenterCheckout(client http.Client, directCookie []string, payloadValue string) {
	payloadBytes, err := json.Marshal(PokemonCenterRequestCheckoutDetails{
		PurchaseFrom: payloadValue,
	},
	)
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/order?format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
}

func PokemonCenterGetPaymentKeyId(client http.Client, directCookie []string) string {
	get := GET{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/payment/key?microform=true&locale=en-US",
	}

	request := PokemonCenterNewRequest(get)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: nil})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)

	return respString
}
func PokemonCenterGetAuthId(client http.Client, directCookie []string) string {
	get := GET{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?format=zoom.nodatalinks",
	}

	request := PokemonCenterNewRequest(get)
	request.Header = PokemonCenterAddHeaders(Header{cookie: directCookie, content: nil})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)

	return respString
}
