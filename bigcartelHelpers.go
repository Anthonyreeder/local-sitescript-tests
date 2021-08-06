package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

func BigCartelClientSetup() http.Client {
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
		Name:   "_storefront_session",
		Value:  "VldwVTJQa1JaMHVWZk1PbXVZWHJwelNtU3ljTlNEcjBQNjg4YS9WV3MvZlZpNEdjMTNmS1k1dWovZklPR0NSc2F6MVNjRzJtMHhHZEhMZnBub3VMSUhVcjhnWlJqaTFVVFI1Z1lQTm5GVjRFWnFjTGZRZXBVTUFZRFRDSHB0ckl1UmxIRUlhT1pCL1J0Q2ttdkQwREFCL2tQVE1jYXp5SzFlTnZudURRMEYwNEpVLy9jVzNmT0RFS1E0VEllaFNNNE92VU5ES05JVVJXREtTNk9MSFBCYzM1MStKeENCUEpEVU00UkxIQUpoeWZHT0doQktFWHowL2N4TzF3cEFSMldzaXRFdlAyYnF5anhOLzBjRjUwZ0ZxYzlJVjVsZ2lsUHk4WHNkM0cydnFycG1EUi9Ba083N1hjbjZiL045NzJJbVZBSmx2Z1Y1SG5qV3NUYW1LK1k3N2xtbitESzY4dkZnT1NxbkMrUFYwPS0tclptaElDUlFVWFE4YUVZRlh2TFJlZz09--f111bb834dbb7fc504cb213cd8ed3c21bef0bb38",
		Path:   "/",
		Domain: ".jimgoldenstudio.bigcartel.com",
	}

	//set datadome cookie to the cookie array
	cookies = append(cookies, cookie)
	//url/domain for cookie
	//u, _ := url.Parse("https://jimgoldenstudio.bigcartel.com")

	//Set client cookieJar with cookiejar array
	//cookieJar.SetCookies(u, cookies)

	return client
}

func BigCartelAddHeaders(header Header) http.Header {
	var x = http.Header{
		"Host": {"shop.thefeebles.com"},
		//Connection: keep-alive
		"sec-ch-ua":                 {"\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\""},
		"sec-ch-ua-mobile":          {"?0"},
		"Upgrade-Insecure-Requests": {"1"},
		"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36"},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
		"Sec-Fetch-Site":            {"none"},
		"Sec-Fetch-Mode":            {"navigate"},
		"Sec-Fetch-User":            {"?1"},
		"Sec-Fetch-Dest":            {"document"},
		"Accept-Language":           {"en-GB,en-US;q=0.9,en;q=0.8"},
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
