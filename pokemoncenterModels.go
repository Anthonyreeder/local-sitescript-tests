package main

import "bytes"

type PokemonCenterLoginResponse struct {
	Access_token string
}

type PokemonCenterNextData struct {
	Props struct {
		IsServer     bool
		InitialState struct {
			Product struct {
				AddToCartForm string
				Availability  string
				ListPrice     struct {
					Amount float32
				}
			}
		}
		InitialProps struct {
		}
	}
}

type PokemonCenterResponseSetPayment struct {
	Self struct {
		Uri string
	}
}

type PokemonCenterResponseKeyId struct {
	KeyId string `json:"keyId"`
}

type PokemonCenterRequestCheckoutDetails struct {
	PurchaseFrom string `json:"purchaseForm"`
}

type PokemonCenterRequestPaymentDetails struct {
	PaymentDisplay string `json:"paymentDisplay"`
	PaymentKey     string `json:"paymentKey"`
	PaymentToken   string `json:"paymentToken"`
}

type PokemonCenterRequestSubmitAddressDetails struct {
	Shipping Address `json:"shipping"`
	Billing  Address `json:"billing"`
}

type Address struct {
	FamilyName      string `json:"familyName"`
	GivenName       string `json:"givenName"`
	StreetAddress   string `json:"streetAddress"`
	ExtendedAddress string `json:"extendedAddress"`
	Locality        string `json:"locality"`
	Region          string `json:"region"`
	PostalCode      string `json:"postalCode"`
	CountryName     string `json:"countryName"`
	PhoneNumber     string `json:"phoneNumber"`
}

type PokemonCenterRequestAddToCart struct {
	Configuration string `json:"configuration"`
	ProductUri    string `json:"productURI"`
	Quantity      int    `json:"quantity"`
}

type PokemonCenterStockCheckResponse struct {
	Self Self
}

type PokemonCenterLoginRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Grant_type string `json:"grant_type"`
	Role       string `json:"role"`
	Scope      string `json:"scope"`
}

type Self struct {
	Type     string
	Id       string
	Quantity int
}

type POST struct {
	Endpoint string
	Payload  *bytes.Reader
}

type GET struct {
	Endpoint string
}

type Header struct {
	cookie  []string
	content *bytes.Reader
}
