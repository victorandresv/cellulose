package models

type CreditCard struct {
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	SecurityCode   string `json:"SecurityCode"`
	Brand          string `json:"Brand"`
}

type Payment struct {
	Type           string     `json:"Type"`
	Amount         int16      `json:"Amount"`
	Installments   int        `json:"Installments"`
	SoftDescriptor string     `json:"SoftDescriptor"`
	CreditCard     CreditCard `json:"CreditCard"`
}

type Sale struct {
	MerchantOrderId string  `json:"MerchantOrderId"`
	Payment         Payment `json:"Payment"`
}
