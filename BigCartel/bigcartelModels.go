package bigcartel

type BigCartelRequestAddToCart struct {
	Configuration string `json:"configuration"`
	ProductUri    string `json:"productURI"`
	Quantity      int    `json:"quantity"`
}
