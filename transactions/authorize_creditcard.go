package transactions

import (
	"bytes"
	"cielo/models"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func AuthorizeCreditCard(OrderId string, Amount int16, CardNumber string, Holder string, ExpirationDate string, SecurityCode string, Brand string) string {
	sale := models.Sale{
		MerchantOrderId: OrderId,
		Payment: models.Payment{
			Type:           "CreditCard",
			Amount:         Amount,
			Installments:   1,
			SoftDescriptor: "123456789ABCD",
			CreditCard: models.CreditCard{
				CardNumber:     CardNumber,
				Holder:         Holder,
				ExpirationDate: ExpirationDate,
				SecurityCode:   SecurityCode,
				Brand:          Brand,
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
