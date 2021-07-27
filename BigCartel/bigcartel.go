package bigcartel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BigCartelDemo() {
	//PokemonCenterAddToCart()
}

//Add to cart
func BigCartelAddToCart(client http.Client, directCookie []string) {

	payloadBytes, err := json.Marshal(BigCartelRequestAddToCart{ProductUri: "", Quantity: 1, Configuration: ""})
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
