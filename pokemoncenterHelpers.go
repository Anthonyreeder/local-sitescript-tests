package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func PokemonCenterClientSetup() http.Client {
	//Create empty cookiejar
	cookieJar, _ := cookiejar.New(nil)
	//Create client
	client := PokemonCenterNewClient()
	//Create and set cookiejar in client
	client.Jar = cookieJar

	//Create temp cookie array
	var cookies []*http.Cookie

	//Create Datadome cookie
	cookie := &http.Cookie{
		Name:   "datadome",
		Value:  "GTtyIGIRRISWq_7uqUKRpltf4UJB2sXEciUyI0zAELMDIkP~W-XNj~ATFjZZwwA_MBfFF30wumX90U3xSoO_KsEglhk6ELhaQS39hRwDqZ",
		Path:   "/",
		Domain: ".pokemoncenter.com",
	}

	//set datadome cookie to the cookie array
	cookies = append(cookies, cookie)

	//url/domain for cookie
	u, _ := url.Parse("http://www.pokemoncenter.com")

	//Set client cookieJar with cookiejar array
	cookieJar.SetCookies(u, cookies)

	return client
}

func PokemonCenterNewResponse(client http.Client, request *http.Request) ([]byte, string) {
	request.Close = true

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, string(body)
}

func PokemonCenterNewClient() http.Client {
	proxyUrl, err := url.Parse("http://localhost:8866")
	if err != nil {
		log.Fatal("Failed + " + err.Error())
	}

	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
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
		"Referer":          {"https://www.pokemoncenter.com/product/706-29048/pokemon-super-special-flip-book-kalos-region-and-unova-region"},
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
