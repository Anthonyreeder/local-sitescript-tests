package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func FindInString(str, start, end string) ([]byte, error) {
	var match []byte
	index := strings.Index(str, start)

	if index == -1 {
		return match, errors.New("Not found")
	}

	index += len(start)

	for {
		char := str[index]

		if strings.HasPrefix(str[index:index+len(match)], end) {
			break
		}

		match = append(match, char)
		index++
	}

	return match, nil
}

func PokemonCenterClientSetup() http.Client {
	//Create empty cookiejar
	cookieJar, _ := cookiejar.New(nil)
	//Create client
	client := NewClient()
	//Create and set cookiejar in client
	client.Jar = cookieJar

	//Create temp cookie array
	var cookies []*http.Cookie

	//Create Datadome cookie
	cookie := &http.Cookie{
		Name:   "datadome",
		Value:  "GN-PNJYbciby6w3IofWWZGxlJhR_MUPrkwj9U_XPL3fIIcZ4BGCFufFudt-j4wNnwXn96~7O14ZOa-7c6dt_gEe~nODOYMXkcxmqzs0IdZ",
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
		"Connection":       {"keep-alive"},
		"sec-ch-ua":        {"Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91"},
		"content-type":     {"application/jwt; charset=UTF-8"},
		"sec-ch-ua-mobile": {"?0"},
		"User-Agent":       {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"},
		"X-Store-Scope":    {"pokemon"},
		"Accept":           {"*/*"},
		"Origin":           {"https://flex.cybersource.com"},
		"Sec-Fetch-Site":   {"same-origin"},
		"Sec-Fetch-Mode":   {"cors"},
		"Sec-Fetch-Dest":   {"empty"},
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
