package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

//Testing etc.
func PokemonCenterDemo() {
	client := PokemonCenterClientSetup()
	//Set-Cookie is not working with this cookie, I think its the format.
	//Also unable to set auth cookie in cookie jar like we do with datadome,
	//I think this is becuase golang is stripping out quotes and breaking the formatting.
	//So I set it directly in the header by passing it to each PokemonCenter task.
	//The addHeader function allows  for direct cookie headers.

	//Must ensure that Datadome cookie (in helpers/setupClient) is up to date
	//Must ensure authCookie above is up to date
	//authId := PokemonCenterGetAuthId(client, []string{})
	authId := PokemonCenterLogin(client, []string{})
	authCookie := []string{"auth={\"access_token\":\"" + authId + "\",\"token_type\":\"bearer\",\"expires_in\":604799,\"scope\":\"pokemon\",\"role\":\"PUBLIC\",\"roles\":[\"PUBLIC\"]}"}

	//atcForm := PokemonCenterConvertSku(client, authCookie)
	//	PokemonCenterStockCheck(client, authCookie, atcForm)
	PokemonCenterAddToCart(client, authCookie)
	//PokemonCenterSubmitAddressDetailsValidate(client, authCookie) //will tell u if there are any issues with address
	//PokemonCenterSubmitAddressDetails(client, authCookie)
	rawKeyId := PokemonCenterGetPaymentKeyId(client, authCookie)

	//Extract KeyID
	var pokemonCenterResponseKeyId PokemonCenterResponseKeyId
	json.Unmarshal([]byte(rawKeyId), &pokemonCenterResponseKeyId)

	paymentKey := CyberSourceV2(pokemonCenterResponseKeyId.KeyId)
	privateKey := PokemonCenterToken(client, []string{}, paymentKey)

	paymentToken := retrievePaymentToken(privateKey)
	//	fmt.Println(pokemonCenterResponseKeyId.KeyId)
	fmt.Println(paymentKey)
	//replace paymentKey and paymentToken with values returned from CyberSourceV2
	//paymentKey = "eyJraWQiOiJ3ZiIsImFsZyI6IlJTMjU2In0.eyJmbHgiOnsicGF0aCI6Ii9mbGV4L3YyL3Rva2VucyIsImRhdGEiOiIzeDdGRjZCUUhrWkt1TnlpQTNMQ2d4QUFFSmk0SGh0MjZMSVg1SnUrc1ZVaGY2UUFXQkFJZ2RnNktCbUJrS3RuUnM4YlBmb1VrUzBXeTYvTVMxMzZNLzlyZ043U1FmZFJtNFo4eXRJYVpOWDByVDdpVjhKcDBuYTFPZHlBYzBaSTk5UWkiLCJvcmlnaW4iOiJodHRwczovL2ZsZXguY3liZXJzb3VyY2UuY29tIiwiandrIjp7Imt0eSI6IlJTQSIsImUiOiJBUUFCIiwidXNlIjoiZW5jIiwibiI6InBVdHdpcGJrT181VnJYeXlOMUI1T21hWVM5UmxkRzJZWHFGYmhlMkJkVHBwdjBfdGtONFhIYkNmNkhKWjl2eVdwRURnYllGUTdsSGlOYnF0UXpWWEUtbDVGdTVhMlBES3N2d0Ryek1kWkQ4R0liU0phaUs0U0RaNERNX3hsVlgwYXBfdm9rVTFyZDJYN0o0MHE0RnUzMUhycURIQ0VCcUVuZFk4MEx4Q3hCX09nMzlhckRmZVpEVmFrNE1tcDJHTFluaG9lM1pOMUFXWm00R0lnNVZFQ0RUZGRPN3VWb05UUWVyVFFpcWxBN3pIQ1lOaXVjYkpqdlE0RUxwSkQyQkdzc0lJSHJjcERfbmVoTmp6NklHVm51VzJ2N25oMGluQUtVTFEyMmRDVU9KUDlkTHhUcC05NG04Z2FXVGd0TWljbHFadWhOX2tnbXRiLURTY21FV3VsdyIsImtpZCI6IjAzelRvU25STGdXN0xNRHZiY2RtUjFZY1ZvV1d4bkFKIn19LCJjdHgiOlt7ImRhdGEiOnsidGFyZ2V0T3JpZ2lucyI6WyJodHRwczovL3d3dy5wb2tlbW9uY2VudGVyLmNvbSIsImh0dHBzOi8vdGVzdC5wb2tlbW9uY2VudGVyLmNvbSIsImh0dHA6Ly9sb2NhbGhvc3Q6MzAwMCJdLCJtZk9yaWdpbiI6Imh0dHBzOi8vZmxleC5jeWJlcnNvdXJjZS5jb20ifSwidHlwZSI6Im1mLTAuMTEuMCJ9XSwiaXNzIjoiRmxleCBBUEkiLCJleHAiOjE2MjU1OTA1MzgsImlhdCI6MTYyNTU4OTYzOCwianRpIjoicDlncU1Wc0xzTGtJOFdIcSJ9.kM6YfQbCK4w4U1SIppPExDQjLmXUoO6JOTFD9JhcKyueNf7F_d5JCYkmeIg_u0BRx5pmkANbSn_ODxYgkfWZPqI6_TR3CTR-BtuZJRhejiciNwYmwVB_NuIjggYTSLEEnumb617otxLalxt7Fyp6LrU0p_BwgPG4Hx9txdPM2AWySaYb4UT57wcITmBE5Yanr_CoXLj0iCnZC43wvnkfdQ6ZHhO7ErRHSghsVT1Yh7D-JsSQ-J7s3oTDpMk-kV5mF4t7klwJcikyOpsr6FTzed9Bn2c-Vnl4F_lkCITDybxlPQcuZnX-MSi_Z4C9HW3fFd2P6q1SraGUSRA8LidH5w"
	//paymentToken := "1C3J1NO3TQJO4TKRDEGZXP42TENC1Z8Q9IH776B50UC08F7I1D1N60E48B1DCBF2"

	//paymentToken := "1C0SIGRUV2U999XPDWYQJBO4FNW2O0DV48AQT8GT0B2GE8O5DSHE60E34D388CC3"

	paymentBytes := PokemonCenterSubmitPaymentDetails(client, authCookie, pokemonCenterResponseKeyId.KeyId, paymentToken)

	//we need to get the return value from payment to pass to checkout
	var pokemonCenterResponseSetPayment = PokemonCenterResponseSetPayment{}
	json.Unmarshal(paymentBytes, &pokemonCenterResponseSetPayment)

	checkoutPayload := strings.Replace(pokemonCenterResponseSetPayment.Self.Uri, "paymentmethods", "purchases", 1) + "/form"

	PokemonCenterCheckout(client, authCookie, checkoutPayload)

	//Todo: Set AUTHID Organically.
}

func PokemonCenterLogin(client http.Client, directCookie []string) string {
	params := url.Values{}
	params.Add("username", "anthonyreeder123@gmail.com")
	params.Add("password", "thekid225")
	params.Add("grant_type", "password")
	params.Add("role", "REGISTERED")
	params.Add("scope", "pokemon")
	fmt.Println(params.Encode())

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/auth?format=zoom.nodatalinks",
		Payload:  bytes.NewReader([]byte(params.Encode())),
	}

	request := PokemonCenterNewRequest(post)
	//request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader([]byte(params.Encode()))})
	respBytes, respString := NewResponse(client, request)
	pokemonCenterLoginResponse := PokemonCenterLoginResponse{}
	json.Unmarshal(respBytes, &pokemonCenterLoginResponse)
	fmt.Println("response Body:", respString)
	return pokemonCenterLoginResponse.Access_token

}

func PokemonCenterStockCheck(client http.Client, directCookie []string, product string) bool {
	payloadBytes, err := json.Marshal(PokemonCenterRequestAddToCart{ProductUri: product, Quantity: 1, Configuration: ""})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?type=product&format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	//	request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	responseBytes, _ := NewResponse(client, request)

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

	payloadBytes, err := json.Marshal(PokemonCenterRequestAddToCart{ProductUri: "/carts/items/pokemon/qgqvhkjxgazs2mbvgqydg=/form", Quantity: 1, Configuration: ""})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?type=product&format=zoom.nodatalinks",
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := PokemonCenterNewRequest(post)
	//request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := NewResponse(client, request)

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
	//	request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := NewResponse(client, request)

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
	//request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := NewResponse(client, request)

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
	//request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	respBytes, respString := NewResponse(client, request)

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
	//	request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := NewResponse(client, request)

	fmt.Println("response Body:", respString)
}

func PokemonCenterGetPaymentKeyId(client http.Client, directCookie []string) string {
	get := GET{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/payment/key?microform=true&locale=en-US",
	}

	request := PokemonCenterNewRequest(get)
	//	request.Header = AddHeaders(Header{cookie: directCookie, content: nil})
	respBytes, _ := NewResponse(client, request)

	//fmt.Println("response Body:", respString)

	return string(respBytes)
}

func PokemonCenterGetAuthId(client http.Client, directCookie []string) string {
	get := GET{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?format=zoom.nodatalinks",
	}

	request := PokemonCenterNewRequest(get)
	//	request.Header = AddHeaders(Header{cookie: directCookie, content: nil})

	_, resp := NewResponse(client, request)

	rawHeader := resp.Header.Get("Set-Cookie")
	re := regexp.MustCompile("({)(.*?)(})")
	match := re.FindStringSubmatch(rawHeader)
	fmt.Println(match[0])

	var auth Auth
	json.Unmarshal([]byte(match[0]), &auth)

	return auth.Access_token
}

func PokemonCenterConvertSku(client http.Client, directCookie []string) string {
	get := GET{
		Endpoint: "https://www.pokemoncenter.com/product/703-05994/", //pass in sku
	}

	request := PokemonCenterNewRequest(get)
	//	request.Header = AddHeaders(Header{cookie: directCookie, content: nil})

	respBytes, _ := NewResponse(client, request)

	responseBody := soup.HTMLParse(string(respBytes))
	nextData := responseBody.Find("script", "id", "__NEXT_DATA__")
	//priceText := priceBlock.Find("span", "class", "visuallyhidden").Text()
	//price, err = strconv.Atoi(reg.ReplaceAllString(priceText, ""))

	nextDataString := nextData.Pointer.FirstChild.Data
	pokemonCenterNextData := PokemonCenterNextData{}
	json.Unmarshal([]byte(nextDataString), &pokemonCenterNextData)

	//rawHeader := resp.Header.Get("Set-Cookie")
	//re := regexp.MustCompile("({)(.*?)(})")
	//match := re.FindStringSubmatch(rawHeader)
	//fmt.Println(match[0])
	//fmt.Println(nextDataString)
	//	//
	//

	return pokemonCenterNextData.Props.InitialState.Product.AddToCartForm
}

func PokemonCenterToken(client http.Client, directCookie []string, payloadBytes string) string {
	//payloadBytes := "eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkEyNTZHQ00iLCJqd2siOnsiZSI6IkFRQUIiLCJraWQiOiIwMzFycHFIMzhrSWxPNmt6dm5hQXFxZU5wU1daMFB1SyIsImt0eSI6IlJTQSIsIm4iOiJoMU9nYk9Za3lia21ERDhqVG9wTUFEQ0dGOHBETG5mOEVXU19PdWh6V3JycVV3M1gyYkZla1Y3TUpJNE1uWnBxVmYxbER2elZYbUlyMkNUSEo3RXN6OEhRTjd1VkNBakdpclFrY2gwWnI4Q0ZSRWRoakxxaFI3QjNHOW5tekgxSnl0QUJJc2lXQjUzeTVUREV2YzAxOWNQRHloQUlHWVpIdFF5ZVdZVTB4SGpUMkxGaFJ6dXhTdVhNZUVkaXNoVEhCWllCWXZMVFFTdnV5Z014bWFSUGFfc2xRZnhEcmtXaEdqWjdmN0hBSnd5UGFjZi0xN2wtX2NydEVNaGFLWDBldGt0cmQ5V1Y3ZFdoS2dUWHJMVXpDWXdtTWJ0M3dPaXFkNkNhTjU4REJ5d2lHZERBbmdMV2JyU2l0XzFMb0x5ZGw5ZDlrTkdpRk1Nb1VyWml3eGdhSVEiLCJ1c2UiOiJlbmMifSwia2lkIjoiMDMxcnBxSDM4a0lsTzZrenZuYUFxcWVOcFNXWjBQdUsifQ.eto8EdVCNwZN6stcuivQOH-oZl6CDx0ljmZVZAofh2cREkqW8r80Sdz6OnHeXlMORz-B1gU0SUEkQrnVv8S3AG94XdF3v_MoE_o_VVkQnmGDiINuJtwksGVZvydntdJRKdR3NoaFCVdngTmsdsXnDGurmmYr2s4kmdYox3DwIAI7F82L6tj0rp0b9mTpl3xKcz9qUMCCxi2Bt5DVPqC588VU1CG5_yjJF-eILP-jMGVZ5Dcgg5xrONryqmfcWCyFo5PLMjiYzaAP16PuDRMhujI3zmJjX_IsU2fb4_F8mNycPzX5LUV1RSNnLiJDK04OjE85SO9Njp76auo9Bi8v1Q.eraN8mno6AHHOuKM.t4x9nmy2Jfsn7bUN-EAPqJ1AOZvh7a-scFIy6R-wmWeFMeK-f27mwHGppApv3c4Wx7TYIof1uDsx22XlFcTAmpVMhy_YvvOXTo66hJi2UN6_DXq7JiU76zH1mynEiZmr5XTY_i2OJK5kEGHYNM-p0jfw5owk3rQBWKZVInGQum0JxGTjtCPlYYa2AAZnymwxCiHdtt2LozFPoVSXh4kVQUS3e99md3WOjqK-CkP6Oy7ln5gPawJONFmX_71I4C_-1wRbwuaGCyXUYXzf_KABWqNvwzZGEXRGSLvZRWrXXGc2K7gpxeFJRvfwgLGDp99usYS8y3HtmDdo8iWj3KQH_QRGRjC4Xpsq_5DuEmbEsjfHnzURh1xqUxO82BS0-htPUoxri5T1dtRzgPacBEIXweTLo9Zbdv3N89KB0h-qhuC8t8FpQSrAaqR4CkqccEPvJjdbdntRBwzL2YupFy7ldGBn8eEjzurrSrOdv8DnP3eifUUHSte2vN3bLeWtAXnB4muR-lkErk3BaVA-vTST-Ev9xDO5Ow1FoulPd-tNlc2LE_Y0Kdc22C7TTvlDNHVsQYDe7QrBkjufztHzGPZzq44O0csEaKxVe3-pa7M-zNq4ZYj0ZbOD7lyZOB8ooqs9EedAfMulafK4xkt06dJtm3dZ6y10xulKUNsdwpq5NxpuOGtI5Fd_kpRA53wNqThgnz6I1nkDcEBc46mZdMhCPDAjnKXbfNRwWVuwvn_hzoxXaAHNdv8eGoeS5Xj2jXT1dyzmqyiPs4-FgyjtBuOAtn0LK0TAz9J41HQC6EzuJd1RIoUY5ZxmClUKJIR7s1ubJ7DPZB2KcGM2Tg1lGAoL7taeAe1nr-IsJf5fcGrgycMcOtwfeCSu6-gIkEQxUDfAM7qQGDEoH5QbIiLxkPbQMjX-BUy8dTcowJ_88DqYb6J41KADN5863RC5HAEXIdWWAP3IQgl5A4pdg4MVcN_aipXXi0HTSPJt4sQ1ERuXreHdOyDvnx3Vs7WrdT51n94PXsUVSUeJYHY42Nf1r5cIm9dvj_Ad5jTZmzu6dqijbAh-73tepKTjPYJuykx6FN2vrMzpMaYtcfwktF5hYqH6GQBr1SRSZOUl_-Np8EGjLpaDPVbcE7fLtxyLu3IWCbWhwrlQgsBW79NLJw29rU2RJMFGkBBU6UMrIjBw6OJuQKYv3NseXTB7mRxDzeOd-V2sHs4mwRSHWVQsEi0d1k_4pJhde2oq3PDXhe5eE_cPGg76NJFtO6ckY8UyL-wq_HymmcXY7Bou47M4Ktubl7R5Ur9frBDiFKyvDYOL-MKNPNAmXGd0WOhchrmkbq2fT1JGmREFmCjqA4LK48MVgy4zBxU1ckzoFWH6abTdk9qLreslBnnV3-2CcZrYPNDOiSjrDojCjK97mA7CqtLiuMouAt9Epfi2ITpECuu_9ASHVQs-94IuRiv2D3o9DKIbJIqjmXePT2doDeLkQawTuLD2nXConjjK0HRJlmdPKBh9gUEmimBwioOB7qlcWOTkwG2eZHtnn3J56C1mvzHCpClXqUcsrwcIlsm940xSBdnniWUy0QkBiI1c1iLLwCWTE6en7LEmOV-qCMcTcvhos5RCQBfiGYxmnp3waDc9_cBIebIf47Ni_emdH3TklZDGssh8QWsWdOG3FGxlOPNR9tX-K5wqkQ6UXJ1mgclWTmiyBRlPe3O4yEyyhW1wBD_VMCtyXei6aV52FloiXRfkbC4lwcbI4zYQFcbMcBXSnfASNRMDXIOHgdknlTygcOt2kZfmgx85oPk-XCUe3Aple8iUtZVEkDSN6Jpx3p1rRP-JFc21H9WiSnpGXtsV6WbJoAGx-wnTGk8PNWv3MssY437LtgTkQjSPvQ7K4U06SGtwGTrGyKL6CVCrjQWthZ7cDw70iBaBuWIZaqptsluLKvGezxofjpUpGahZdeHtTAKPjgxGM6BIvzWI2KoEtOARIsWi55jmKIRYFrOhb_CuP1Jc6pb61BANbKydRedfLKsFAC-Wi8bNFD_7RDU6TUxTyxfEcxMGoEr0JpTNLW4sBIPSm1FKF6kXj5BQwhb3vMPokDxwEfK8GZ-W5G_lJvJ1iCxYN9n8vQxC4sj7FWSnjppjl7mTWAqFDtVBshyk5UjQa9_l6xOam-QnVoCQ_-Y1DrviFwextX7mi1Y67Nz-obUQRy_CT8nVIwYjPmF47EGf4npaeEN-LudzzyRYydLx8bZixxByLEvLK_TujSx8va_v9DBviKxap0cb9oPU8YgMEvhi4d8B3NDD9FYR-9gaOcs4OVXxNPB9InVZAA9DT6ixRrCc65lFAtMtdPoVdwiVx1oVrwfdEWieVSDunDMhSvQLcmMReB1S4Q.S8NT-jS5LJFsLIJ06rCjAg"

	post := POST{
		Endpoint: "https://flex.cybersource.com/flex/v2/tokens",
		Payload:  bytes.NewReader([]byte(payloadBytes)),
	}

	request := PokemonCenterNewRequest(post)
	//	request.Header = AddHeaders(Header{cookie: directCookie, content: bytes.NewReader([]byte(payloadBytes))})
	respbytes, _ := NewResponse(client, request)

	//fmt.Println("response Body:", respString)
	return string(respbytes)
}

type Auth struct {
	Access_token string
}
