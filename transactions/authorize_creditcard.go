package transactions

import (
	"bytes"
	"cielo/models"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func AuthorizeCreditCard() string {
	sale := models.Sale{
		MerchantOrderId: "123",
		Payment: models.Payment{
			Type:           "CreditCard",
			Amount:         15700,
			Installments:   1,
			SoftDescriptor: "123456789ABCD",
			CreditCard: models.CreditCard{
				CardNumber:     "4551870000000181",
				Holder:         "Teste Holder",
				ExpirationDate: "12/2021",
				SecurityCode:   "123",
				Brand:          "Visa",
			},
		},
	}
	payload, _ := json.Marshal(sale)
	reader := bytes.NewReader(payload)

	request, _ := http.NewRequest("POST", os.Getenv("API_TRANSACTIONS_URL")+"/1/sales", reader)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("MerchantId", os.Getenv("MERCHANT_ID"))
	request.Header.Add("MerchantKey", os.Getenv("MERCHANT_KEY"))
	result, err := http.DefaultClient.Do(request)

	if err != nil {
		return ""
	}
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)

	return string(body)
}
