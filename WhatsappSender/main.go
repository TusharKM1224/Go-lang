package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	whatsappAPIURL = "https://graph.facebook.com/v22.0/<PhoneNumberID>/messages"
	accessToken    = "<access token>"
)

func sendWhatsAppOTP(toPhoneNumber, otp string) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"messaging_product": "whatsapp",
			"to":                toPhoneNumber,
			"type":              "template",
			"template": map[string]interface{}{
				"name": "<Approved_template_name>",
				"language": map[string]string{
					"code": "en_US",
				},
				"components": []map[string]interface{}{
					{
						"type": "body",
						"parameters": []map[string]string{
							{"type": "text", "text": otp},
						},
					},
					{
						"type":     "button",
						"sub_type": "url",
						"index":    "0",
						"parameters": []map[string]string{
							{"type": "text", "text": otp},
						},
					},
				},
			},
		}).
		Post(whatsappAPIURL)

	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

	fmt.Println("Response:", resp.String())
}

func main() {
	sendWhatsAppOTP("<recipentPhoneNumber>", "<otp>")
}
