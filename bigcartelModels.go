package main

type BigCartelRequestAddToCart struct {
	Configuration string `json:"configuration"`
	ProductUri    string `json:"productURI"`
	Quantity      int    `json:"quantity"`
}

type BigCartelRequestSubmitNameAndEmail struct {
	Buyer_email                 string `json:"buyer_email"`
	Buyer_first_name            string `json:"buyer_first_name"`
	Buyer_last_name             string `json:"buyer_last_name"`
	Buyer_opted_in_to_marketing bool   `json:"buyer_opted_in_to_marketing"`
	Buyer_phone_number          string `json:"buyer_phone_number"`
}

type BigCartelRequestSubmitAddress struct {
	Shipping_address_1             string `json:"shipping_address_1"`
	Shipping_address_2             string `json:"shipping_address_2"`
	Shipping_city                  string `json:"shipping_city"`
	Shipping_country_autofill_name string `json:"shipping_country_autofill_name"`
	Shipping_country_id            string `json:"shipping_country_id"`
	Shipping_state                 string `json:"shipping_state"`
	Shipping_zip                   string `json:"shipping_zip"`
}

type BigCartelRequestSubmitPaymentMethodResponse struct {
	Id              string
	Object          string
	Billing_details struct {
		Address struct {
			City        string
			Country     string
			Line1       string
			Line2       string
			Postal_code string
			State       string
		}
		Email string
		Name  string
		Phone string
	}
	Card struct {
		Brand  string
		Checks struct {
			Address_line1_check       string
			Address_postal_code_check string
			Cvc_check                 string
		}
		Country        string
		Exp_month      string
		Exp_year       string
		Funding        string
		Generated_from string
		Last4          string
		Networks       struct {
			Available []string
			Preferred string
		}
		Three_d_secure_usage struct {
			Supported bool
		}
		Wallet string
	}
	Created  string
	Customer string
	Livemode string
	Type     string
}

type Payment struct {
	Cart_token               string `json:"cart_token"`
	Stripe_payment_method_id string `json:"stripe_payment_method_id"`
	Stripe_payment_intent_id string `json:"stripe_payment_intent_id"`
}

type Payment2 struct {
	Stripe_payment_method_id string `json:"stripe_payment_method_id"`
}

type Note struct {
	Note string `json:"note"`
}
