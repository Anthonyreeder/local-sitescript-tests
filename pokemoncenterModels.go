package main

import "bytes"

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
