package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Demo() {
	PokemonCenterGetAuth()
}

type PokemonCenterRequestAddToCart struct {
	Configuration string `json:"configuration"`
	ProductUri    string `json:"productURI"`
	Quantity      int    `json:"quantity"`
}

func PokemonCenterGetAuth() {

	get := GET{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/order?type=status&format=zoom.nodatalinks",
	}

	client := PokemonCenterNewClient()
	request := PokemonCenterNewRequest(get)
	request.Header = PokemonCenterAddHeaders(Header{cookie: []string{"datadome=74zhxkLgdEtsbbguS0Lf-aaSIPd0AOEM4Tc6lozwPDjQfaHKq9maCIX1-qpL2WLxh1qZWmNBsLvbzh.btZ59AKwgS-R~_b7nbWUMEUdkdz"}})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)

	//var authCookie = request.getData("https://www.pokemoncenter.com/tpci-ecommweb-api/order?type=status&format=zoom.nodatalinks", true)
}
func PokemonCenterAddTOCart() {

	payloadBytes, err := json.Marshal(PokemonCenterRequestAddToCart{Configuration: "{}", ProductUri: "\"/carts/items/pokemon/qgqvhljxgqys2mbzgiydolkbfvjq=/form\"", Quantity: 1})
	if err != nil {
		log.Fatal("Marshal payload failed with error " + err.Error())
	}

	post := POST{
		Endpoint: "https://www.pokemoncenter.com/tpci-ecommweb-api/cart?type=product&format=zoom.nodatalinks",
		Payload:  bytes.NewReader([]byte(payloadBytes)),
	}

	client := PokemonCenterNewClient()
	request := PokemonCenterNewRequest(post)
	request.Header = PokemonCenterAddHeaders(Header{cookie: []string{"datadome=74zhxkLgdEtsbbguS0Lf-aaSIPd0AOEM4Tc6lozwPDjQfaHKq9maCIX1-qpL2WLxh1qZWmNBsLvbzh.btZ59AKwgS-R~_b7nbWUMEUdkdz"}, content: bytes.NewReader(payloadBytes)})
	_, respString := PokemonCenterNewResponse(client, request)

	fmt.Println("response Body:", respString)
}

func PokemonCenterNewResponse(client *http.Client, request *http.Request) ([]byte, string) {
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, string(body)
}

func PokemonCenterNewClient() *http.Client {
	proxyUrl, err := url.Parse("http://localhost:8866")
	if err != nil {
		log.Fatal("Failed + " + err.Error())
	}

	return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}

func PokemonCenterNewRequest(requestType interface{}) *http.Request {
	switch v := requestType.(type) {
	case POST:
		req, err := http.NewRequest("POST", v.Endpoint, v.Payload)
		if err != nil {
			log.Fatal("Failed + " + err.Error())
		}
		return req

	case GET:
		req, err := http.NewRequest("GET", v.Endpoint, nil)
		if err != nil {
			log.Fatal("Failed + " + err.Error())
		}
		return req

	default:
		log.Fatal("Request type was invalid")
		return nil
	}
}

func PokemonCenterAddHeaders(header Header) http.Header {
	var x = http.Header{
		"Host":             {"www.pokemoncenter.com"},
		"Connection":       {"keep-alive"},
		"sec-ch-ua":        {"Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91"},
		"content-type":     {"application/json"},
		"sec-ch-ua-mobile": {"?0"},
		"User-Agent":       {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"},
		"X-Store-Scope":    {"pokemon"},
		"Accept":           {"*/*"},
		"Origin":           {"https://www.pokemoncenter.com"},
		"Sec-Fetch-Site":   {"same-origin"},
		"Sec-Fetch-Mode":   {"cors"},
		"Sec-Fetch-Dest":   {"empty"},
		"Referer":          {"https://www.pokemoncenter.com/product/741-09207/pikachu-and-raichu-black-polo-shirt-adult"},
		"Accept-Language":  {"en-GB,en-US;q=0.9,en;q=0.8"},
	}

	if header.content != nil {
		x.Set("Content-Length", fmt.Sprint(header.content.Size()))
	}

	if len(header.cookie) > 0 {
		buildString := ""
		for i := 0; i < len(header.cookie); i++ {
			buildString += header.cookie[i] + "; "
		}
		x.Set("Cookie", buildString+strings.Join(x.Values("Cookie"), "; "))
	}

	return x
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
