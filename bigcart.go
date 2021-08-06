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

func BigCartelDemo() {
	client := BigCartelClientSetup()

	storeId, cartToken := BigCartelAtc(client, []string{}) //needed to add to cart
	if storeId == "" && cartToken == "" {
		storeId, cartToken = BigCartelGetPage(client, []string{}, "https://shop.thefeebles.com/checkout")
	}

	if storeId != "" && cartToken != "" {
		BigCartelSubmitNameAndEmail(client, []string{}, storeId, cartToken)
		BigCartelSubmitAddress(client, []string{}, storeId, cartToken)
		paymentId := BigCartelSubmtPaymentInfo(client, []string{})
		BigCartelCheckout(client, []string{}, storeId, paymentId, cartToken)
		BigCartelSubmtPaymentDetails(client, []string{}, storeId, paymentId, cartToken)

	}
}

//Add to cart
func BigCartelGetPage(client http.Client, directCookie []string, page string) (string, string) {
	get := GET{
		Endpoint: page,
	}

	request := NewRequest(get)
	request.Header = BigCartelAddHeaders(Header{cookie: directCookie})
	_, respString := NewResponse(client, request)
	s := []string{}
	if strings.Contains(respString.Request.URL.String(), "checkout.bigcartel.com/") {
		s = strings.Split(string(respString.Request.URL.String()), "/")
		return s[3], s[4]
	}

	//Please enable JS and disa	ble any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)

	return "", ""
}

func BigCartelAtc(client http.Client, directCookie []string) (string, string) {

	payload := url.Values{
		"cart[add][id]": {"268318059"},
		"submit":        {""},
	}

	post := POSTUrlEncoded{
		Endpoint:       "https://shop.thefeebles.com/cart",
		EncodedPayload: payload.Encode(),
	}

	request := NewRequest(post)
	request.Header = AddHeaders3(Header{cookie: directCookie, content: bytes.NewReader([]byte(payload.Encode()))})
	_, respString := NewResponse(client, request)
	s := []string{}
	if strings.Contains(respString.Request.URL.String(), "checkout.bigcartel.com/") {
		s = strings.Split(string(respString.Request.URL.String()), "/")
		return s[3], s[4]
	}

	//Please enable JS and disa	ble any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)

	return "", ""
}

func BigCartelSubmitNameAndEmail(client http.Client, directCookie []string, code1 string, code2 string) {

	payloadBytes, _ := json.Marshal(BigCartelRequestSubmitNameAndEmail{
		Buyer_email:                 "anthonyreeder123@gmail.com",
		Buyer_first_name:            "Anthony",
		Buyer_last_name:             "Reeder",
		Buyer_opted_in_to_marketing: false,
		Buyer_phone_number:          "+1 (231) 231-2312",
	})

	post := POST{
		Endpoint: "https://api.bigcartel.com/store/" + code1 + "/carts/" + code2,
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := NewRequest(post)
	request.Header = AddHeaders3(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := NewResponse(client, request)

	//Please enable JS and disable any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)
}

func BigCartelSubmitAddress(client http.Client, directCookie []string, code1 string, code2 string) {

	payloadBytes, _ := json.Marshal(BigCartelRequestSubmitAddress{
		Shipping_address_1:             "49 Thackeray Close",
		Shipping_address_2:             "",
		Shipping_city:                  "Royston",
		Shipping_country_autofill_name: "",
		Shipping_country_id:            "43",
		Shipping_state:                 "Hawaii",
		Shipping_zip:                   "4353453453",
	})

	post := POST{
		Endpoint: "https://api.bigcartel.com/store/" + code1 + "/carts/" + code2,
		Payload:  bytes.NewReader(payloadBytes),
	}

	request := NewRequest(post)
	request.Header = AddHeaders3(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})
	_, respString := NewResponse(client, request)

	//Please enable JS and disable any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)
}

func BigCartelSubmtPaymentInfo(client http.Client, directCookie []string) string {
	payload := url.Values{
		"type":                                  {"card"},
		"billing_details[name]":                 {"Anthony Reeder"},
		"billing_details[address][line1]":       {"49 Thackeray Close"},
		"billing_details[address][line2]":       {""},
		"billing_details[address][city]":        {"Royston"},
		"billing_details[address][state]":       {"Hawaii"},
		"billing_details[address][postal_code]": {"4353453453"},
		"billing_details[address][country]":     {"US"},
		"card[number]":                          {"4767718212263745"},
		"card[cvc]":                             {"260"},
		"card[exp_month]":                       {"02"},
		"card[exp_year]":                        {"26"},
		"guid":                                  {"13bedf9f-0e16-4243-9024-2eecbc91113ee4bb60"},
		"muid":                                  {"f5023d3c-907a-42db-8dc4-d4ef7f2413436b0869"},
		"sid":                                   {"39c026e7-5d19-4c2f-93f5-058749b4de5fa54e58"},
		"pasted_fields":                         {"number"},
		"payment_user_agent":                    {"stripe.js/76aee18e6; stripe-js-v3/76aee18e6"},
		"time_on_page":                          {"13709"},
		"referrer":                              {"https://checkout.bigcartel.com/"},
		"key":                                   {"pk_live_HAopYDMYyyhaXP505VRbXQtT"}, //i think this is speicfic to each site.
	}

	post := POSTUrlEncoded{
		Endpoint:       "https://api.stripe.com/v1/payment_methods",
		EncodedPayload: payload.Encode(),
	}

	request := NewRequest(post)
	request.Header = AddHeaders3(Header{cookie: directCookie, content: bytes.NewReader([]byte(payload.Encode()))})

	respBytes, respString := NewResponse(client, request)
	bigCartelRequestSubmitPaymentMethodResponse := BigCartelRequestSubmitPaymentMethodResponse{}
	json.Unmarshal(respBytes, &bigCartelRequestSubmitPaymentMethodResponse)
	//Please enable JS and disable any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)
	return bigCartelRequestSubmitPaymentMethodResponse.Id

}

func BigCartelSubmtPaymentDetails(client http.Client, directCookie []string, storeId, paymentId, cartToken string) {

	test := "{\"cart_token\":\"" + cartToken + "\",\"stripe_payment_method_id\":\"" + paymentId + "\",\"stripe_payment_intent_id\":null}"
	test2 := []byte(test)

	params := url.Values{}
	params.Add("", "%00")

	post := POST{
		Endpoint: "https://api.bigcartel.com/store/" + storeId + "/checkouts",
		Payload:  bytes.NewReader(test2),
	}
	request := NewRequest(post)
	request.Header = AddHeaders4(Header{cookie: directCookie, content: bytes.NewReader(test2)})

	respBytes, respString := NewResponse(client, request)
	bigCartelRequestSubmitPaymentMethodResponse := BigCartelRequestSubmitPaymentMethodResponse{}
	json.Unmarshal(respBytes, &bigCartelRequestSubmitPaymentMethodResponse)
	//Please enable JS and disable any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)

}

func BigCartelCheckout(client http.Client, directCookie []string, storeId, paymentId, cartToken string) {

	payloadBytes, err := json.Marshal(Payment2{
		Stripe_payment_method_id: paymentId,
	})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://api.bigcartel.com/store/" + storeId + "/carts/" + cartToken,
		Payload:  bytes.NewReader(payloadBytes),
	}
	request := NewRequest(post)
	request.Header = AddHeaders4(Header{cookie: directCookie, content: bytes.NewReader(payloadBytes)})

	respBytes, respString := NewResponse(client, request)
	bigCartelRequestSubmitPaymentMethodResponse := BigCartelRequestSubmitPaymentMethodResponse{}
	json.Unmarshal(respBytes, &bigCartelRequestSubmitPaymentMethodResponse)
	//Please enable JS and disable any ad blocker = captcha, New cookie needed.

	fmt.Println("response Body:", respString)

}
